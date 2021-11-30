package microkernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	// 这里receiver其实是agent, agent实现了OnEvent方法也是EventReceiver类型, 通过EventReceiver解耦
	evtReceiver EventReceiver
	stopChan    chan struct{}
	name        string
	content     string
}

func (c *DemoCollector) Init(e EventReceiver) error {
	fmt.Println("initialize collector", c.name)
	c.evtReceiver = e
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector", c.name)
	// 模拟每50毫秒发送一次
	for {
		select {
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collector", c.name)
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("failed to stop for timeout")
	}
}

func (c *DemoCollector) Destroy() error {
	fmt.Println(c.name, "released resources.")
	return nil
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	agt.RegistCollector("c1", c1)
	agt.RegistCollector("c2", c2)
	if err := agt.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	fmt.Println(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destory()
}
