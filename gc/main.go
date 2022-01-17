package main

import (
	"fmt"
	"runtime"
)

const VAT = 24

type Item struct {
	Description string
	Value       float64
}

func Value(price float64) float64 {
	total := price + price*VAT/100
	return total
}

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC, "\n")
}

func main() {
	// go run -gcflags '-m' gc/main.go
	t := Item{Description: "keyboard", Value: 100}
	t.Value = Value(t.Value)
	fmt.Println(t)

	tP := &Item{}
	*&tP.Description = "Mouse"
	*&tP.Value = 100
	fmt.Println(tP)

	// GODEBUG=gctrace=1 go run gc/main.go
	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i < 10; i++ {
		// Allocating 50,000,000 bytes
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)
}
