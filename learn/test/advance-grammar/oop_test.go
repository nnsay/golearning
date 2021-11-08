package advancegrammar_test

import (
	"fmt"
	"testing"
)

type Code string
type Programer interface {
	WriteHelloWorld() Code
}

type GoProgramer struct{}

func (p *GoProgramer) WriteHelloWorld() Code {
	return "fmt.Println(\"hello world\")"
}

type NodeProgramer struct{}

func (p *NodeProgramer) WriteHelloWorld() Code {
	return "console.log(\"hello world\")"
}

// 调用接口类型的对象
func writeFirstProgram(p Programer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPlayOOP(t *testing.T) {
	goProg := new(GoProgramer)
	nodeProg := new(NodeProgramer)
	writeFirstProgram(goProg)
	writeFirstProgram(nodeProg)
}
