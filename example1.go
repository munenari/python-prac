package main

import (
	"fmt"
	"time"

	"github.com/munenari/python-prac/calc"
)

func main1() {
	start_time := time.Now()
	ranges := [][2]int{
		{1_000_000, 1_600_000},
		{2_000_000, 2_900_000},
		{3_000_000, 3_300_000},
	}
	for _, pair := range ranges {
		calc.Fcalc_primes(pair[0], pair[1])
	}
	fmt.Printf("Processing time: %v seconds\n", time.Since(start_time))
}

// import time
// from prime_number import calc_primes

// start_time = time.time()  # 処理時間計算用

// ranges = [  # 素数を計算する開始値と終了値
//     (1_000_000, 1_600_000),
//     (2_000_000, 2_900_000),
//     (3_000_000, 3_300_000),
// ]
// for num_start, num_end in ranges:
//     calc_primes(num_start, num_end)

// end_time = time.time()  # 終了時間
// print(f"Processing time: {end_time - start_time} seconds")
