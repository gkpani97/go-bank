package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	alphabet = "abcdefjhijklmnopqrstuvwxhyz"
	rng      *rand.Rand
)

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, INR}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}

func RandomMoney(maxAmt int64) int64 {
	return RandomInt(0, maxAmt)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
