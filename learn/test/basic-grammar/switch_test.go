package golearning_test

import (
	"fmt"
	"testing"
)

func price(weight int) {
	switch {
	case weight > 10:
		fmt.Println(">100")
	case weight > 8:
		fmt.Println(">8")
		fallthrough
	case weight > 5:
		fmt.Println(">5")
	default:
		fmt.Println("default")
	}
}
func TestSwitch(t *testing.T) {
	price(10)
}
