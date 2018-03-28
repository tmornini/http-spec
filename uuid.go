package main

import (
	"crypto/rand"
	"math/big"
)

var chars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var big62 = big.NewInt(62)
var space = (&big.Int{}).Exp(big62, big.NewInt(22), nil)

// newRandom generate a random uuid
func newRandom() (string, error) {
	uuid, err := rand.Int(rand.Reader, space)

	if err != nil {
		return "", err
	}

	base62 := ""

	for i := 0; i < 22; i++ {
		index := (&big.Int{}).Set(uuid)

		index.Mod(index, big62)

		base62 = base62 + chars[index.Uint64()]

		uuid.Div(uuid, big62)
	}

	return reverse(base62), nil
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
