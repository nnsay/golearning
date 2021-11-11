package advancegrammar_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func isCancelByContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelContext(t *testing.T) {
	var mu sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		mu.Add(1)
		go func(ii int, c context.Context, wg *sync.WaitGroup) {
			for {
				// fmt.Printf("goroutine %d is running\n", ii)
				if isCancelByContext(c) {
					break
				}
			}
			fmt.Printf("goroutine %d is canceled\n", ii)
			wg.Done()
		}(i, ctx, &mu)
	}
	cancel()
	mu.Wait()

	fmt.Println("all done")
}
