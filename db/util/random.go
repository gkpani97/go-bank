package util

import (
	"math/rand"
	"strings"
	"time"
)

var (
	alphabet = "abcdefjhijklmnopqrstuvwxhyz"
	rng *rand.Rand
)

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt(min, max int64) int64{
	return min + rng.Int63n(max - min + 1)
}

func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string{
	return RandomString(6)
}

func RandomBalance() int64{
	return RandInt(0, 1000)
}

func RandomCurrency() string{
	currencies := []string{"EUR", "USD", "INR"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}