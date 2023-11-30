package calc

import (
	"fmt"
	"math"
)

func is_prime(num int) bool {
	for i := 2; i <= int(math.Pow(float64(num), 0.5))+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Fcalc_primes(num_start, num_end int) {
	prime_count := 0
	for num := num_start; num <= num_end; num++ {
		if is_prime(num) {
			prime_count++
		}
	}
	fmt.Printf("Prime Count [%d-%d]: %d\n", num_start, num_end, prime_count)
}

// def calc_primes(num_start, num_end):
//     prime_count = 0
//     for num in range(num_start, num_end):
//         # 試し割法を使って素数を計算
//         for i in range(2, int(num**0.5) + 1):
//             if num % i == 0:  # 割り切れる数値がある場合は素数ではない
//                 break
//         else:  # ループが最後まで到達したら素数と判定
//             prime_count += 1
//     print(f"Prime Count [{num_start:,}-{num_end:,}]: {prime_count}")
