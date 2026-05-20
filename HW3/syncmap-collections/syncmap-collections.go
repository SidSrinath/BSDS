package main

import (
	"fmt"
	"sync"
)

func perform(g int, m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		m.Store(g*1000+i, i)
	}
}

func main() {
	var m sync.Map

	var wg sync.WaitGroup

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go perform(g, &m, &wg)
	}

	wg.Wait()

	i := 0
	m.Range(func(key, value any) bool {
		i++
		return true
	})

	fmt.Println("Length of map: ", i)
}
