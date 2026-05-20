package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//var value int
	var value atomic.Uint64

	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				//value++
				value.Add(1)
			}
		})
	}

	wg.Wait()

	//fmt.Println("Final value: ", value)
	fmt.Println("Final value: ", value.Load())
}
