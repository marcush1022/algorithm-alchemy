// package golang
// name 204.(哈希)Count_Primes

package golang

import (
	"math"
)

func CountPrimes(n int) (count int) {
	if n < 2 {
		return 0
	}

	limit := int(math.Sqrt(float64(n)))
	count = 0
	primes := make([]bool, n, n)

	for i := 0; i < len(primes); i++ {
		primes[i] = true
	}

	for i := 2; i < n; i++ {
		if v := primes[i]; v == true {
			count++
		}
		if i > limit {
			continue
		}
		for j := i * i; j < n; j += i {
			primes[j] = false
		}
	}
	return count
}
