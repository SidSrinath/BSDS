package main

import (
	"fmt"
	"sync"
)

func perform(g int, mpp map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mpp[g*1000+i] = i
	}
}

func main() {
	mpp := make(map[int]int)

	var wg sync.WaitGroup

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go perform(g, mpp, &wg)
	}

	wg.Wait()

	fmt.Println("Length of map: ", len(mpp))
}
