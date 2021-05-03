package unit_tool

import (
	"time"
	"math/rand"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandomInt() {

}

func RandomString() string {
	return string(r.Intn(26) + 65)
}