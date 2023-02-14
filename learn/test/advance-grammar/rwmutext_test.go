package advancegrammar_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Secret struct {
	RWM      sync.RWMutex
	password string
}

func TestRWMutext(t *testing.T) {
	var Password *Secret
	var wg sync.WaitGroup
	Password = &Secret{password: "abc"}

	change := func(pass string) {
		defer wg.Done()
		fmt.Println("change() function")
		Password.RWM.Lock()
		fmt.Println("change() locked")
		Password.password = pass
		Password.RWM.Unlock()
		fmt.Println("change() unlocked")
	}

	show := func(name string) {
		defer wg.Done()
		Password.RWM.RLock()
		fmt.Println(name, "show() function locked!")
		time.Sleep(2 * time.Second)
		fmt.Println(name, "show() pass value:", Password.password)
		Password.RWM.RUnlock()
		fmt.Println(name, "show() function unlocked!")
	}
	wg.Add(1)
	go show("1")
	wg.Add(1)
	go show("2")

	wg.Add(1)
	go change("123") // 除非所有的reader锁被释放, 否则write锁会被block

	wg.Wait()
}
