package advancegrammar_test

import (
	"errors"
	"testing"
	"time"
)

type ReadableObj struct{}

type ResourcePool struct {
	readChan chan *ReadableObj
}

func (rp *ResourcePool) init(poolSize int) {
	rp.readChan = make(chan *ReadableObj, poolSize)
	for i := 0; i < poolSize; i++ {
		rp.readChan <- &ReadableObj{}
	}
}

func (rp *ResourcePool) getResource(timeout time.Duration) (*ReadableObj, error) {
	//慢请求比快失败更严重, 所以需要增加超时
	select {
	case obj := <-rp.readChan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (rp *ResourcePool) pushResource(obj *ReadableObj) error {
	// default就是就是短路处理, 只有在chan不可用时的时候才执行,
	// 如果不关闭那么这里唯一不可用的原因是资源池已满
	select {
	case rp.readChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestPool(t *testing.T) {
	rp := &ResourcePool{}
	rp.init(10)
	// 测试overfow
	// if err := rp.pushResource(&ReadableObj{}); err != nil {
	// 	t.Error(err)
	// 	return
	// }
	for i := 0; i < 11; i++ {
		if v, err := rp.getResource(time.Second * 1); err != nil {
			t.Error(err)
			break
		} else {
			t.Logf("%T", v)
			// 测试timeout
			// if err := rp.pushResource(v); err != nil {
			// 	t.Error(err)
			// 	break
			// }
		}
	}
	t.Log("All Done")
}
