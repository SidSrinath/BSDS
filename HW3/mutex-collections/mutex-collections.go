package main

import (
	"fmt"
	"sync"
)

type Data struct {
	mpp map[int]int
	mu  sync.RWMutex
}

func perform(g int, data *Data, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		data.mu.Lock()
		data.mpp[g*1000+i] = i
		data.mu.Unlock()
	}
}

func main() {
	data := Data{
		mpp: make(map[int]int),
	}

	var wg sync.WaitGroup

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go perform(g, &data, &wg)
	}

	wg.Wait()

	fmt.Println("Length of map: ", len(data.mpp))
}
