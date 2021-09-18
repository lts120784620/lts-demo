package unit

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandomInt() int {
	return rand.Intn(99)
}

func RandomString() string {
	return string(r.Intn(26) + 65)
}

func RandomArrays(size int) []int {
	arr := []int{}
	for i := 0; i < size; i++ {
		arr = append(arr, RandomInt())
	}
	return arr
}
