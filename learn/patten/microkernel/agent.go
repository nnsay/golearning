package microkernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Event struct {
	Content string
	Source  string
}

type EventReceiver interface {
	OnEvent(e Event)
}

type CollectorError struct {
	CollectorErrors []error
}

// 只有CollectorError实现了Error方法, 才是error接口类型
func (ce *CollectorError) Error() string {
	var errString []string
	for _, err := range ce.CollectorErrors {
		errString = append(errString, err.Error())
	}
	return strings.Join(errString, ";")
}

// 定义个Collector接口, 第三方实现该接口
type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event

	cancel context.CancelFunc
	ctx    context.Context

	state int
}

const (
	Waiting = iota
	Running
)

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}

	return &agt
}

var AgentErrorStatus = errors.New("the agent status is invalid")

func (agt *Agent) EventProcessGroutine() {
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(evtSeg)
	}
}

// 大写注册方法可以暴露
func (agt *Agent) RegistCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return AgentErrorStatus
	}

	agt.collectors[name] = collector
	return collector.Init(agt)
}

// 实现三个主要功能内部方法, 小写不暴露
func (agt *Agent) startCollectors() error {
	var mux sync.Mutex
	var errs CollectorError

	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mux.Unlock()
			}()
			err := collector.Start(ctx)
			mux.Lock()
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors, errors.New(name+":"+err.Error()))
			}

		}(name, collector, agt.ctx)
	}

	if len(errs.CollectorErrors) != 0 {
		return errors.New(errs.Error())
	}
	return nil
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorError
	for name, collector := range agt.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}

	if len(errs.CollectorErrors) != 0 {
		return errors.New(errs.Error())
	}
	return nil
}

func (agt *Agent) destoryCollectors() error {
	var err error
	var errs CollectorError
	for name, collector := range agt.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) != 0 {
		return errors.New(errs.Error())
	}
	return nil
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return AgentErrorStatus
	}
	agt.state = Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	// agent自身监控, 因为阻塞, 所以协程启动
	go agt.EventProcessGroutine()
	// agent调用第三方Collector各自的业务逻辑
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return AgentErrorStatus
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent) Destory() error {
	if agt.state != Waiting {
		return AgentErrorStatus
	}
	return agt.destoryCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}
