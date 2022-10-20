package main

import (
	"fmt"
	"sync"
)

type Summer struct {
	n  int
	mu sync.Mutex
}

func (s *Summer) add(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.n += i

}

func main() {
	sum := Summer{}
	sumAll := Summer{}

	wg := sync.WaitGroup{}

	fn := func() {
		defer wg.Done()

		for i := 1; i <= 10000; i++ {
			sum.add(i)
		}

		fmt.Println(sum.n)
	}

	fn1 := func() {
		defer wg.Done()

		for i := 10000; i <= 20000; i++ {
			sumAll.add(i)
		}

		fmt.Println(sumAll.n)
	}

	wg.Add(2)
	go fn1()
	go fn()

	wg.Wait()
}
