package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/munenari/python-prac/calc"
)

func main4() {
	start_time := time.Now()
	ranges := [][2]int{
		{1_000_000, 1_600_000},
		{2_000_000, 2_900_000},
		{3_000_000, 3_300_000},
	}
	wg := &sync.WaitGroup{}
	for _, pair := range ranges {
		wg.Add(1)
		num_start := pair[0]
		num_end := pair[1]
		go func() {
			defer wg.Done()
			calc.Fcalc_primes(num_start, num_end)
		}()
	}
	wg.Wait()
	fmt.Printf("Processing time: %v seconds\n", time.Since(start_time))
}
