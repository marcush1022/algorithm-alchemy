// package golang
// name 204.(哈希)Count_Primes_test

package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	assert.Equal(t, 0, CountPrimes(1))
	assert.Equal(t, 1, CountPrimes(3))
	assert.Equal(t, 2, CountPrimes(5))
}
