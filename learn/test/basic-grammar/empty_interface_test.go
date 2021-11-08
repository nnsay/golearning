package golearning_test

import (
	"fmt"
	"testing"
)

func GetType(p interface{}) {
	// if
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if i, ok := p.(string); ok {
	// 	fmt.Println("String", i)
	// 	return
	// }
	// fmt.Println("Unknow Type")

	// switch
	switch p.(type) {
	case int:
		fmt.Println("Integer", p)
	case string:
		fmt.Println("String", p)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	GetType(10)
	GetType("10")
}
