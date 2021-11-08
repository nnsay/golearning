package golearning_test

import "testing"

type Programer interface {
	writeHello() string
}

type GoProgramer struct{}

func (p *GoProgramer) writeHello() string {
	return "fmt.Println(\"hello word\")"
}

func (p *GoProgramer) sayHello() string {
	return "fmt.Println(\"hello word, I'm golang\")"
}

func TestInherite(t *testing.T) {
	var p Programer = new(GoProgramer)
	var g = new(GoProgramer)

	t.Log(p.writeHello())
	t.Log(g.sayHello())
}
