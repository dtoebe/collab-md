package utils

import (
	"math/rand"
	"time"
)

var chars = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRZTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenRandStr(l int) string {
	r := make([]rune, l)
	ll := len(chars)

	for i := range r {
		r[i] = chars[rand.Intn(ll)]
	}

	return string(r)
}
