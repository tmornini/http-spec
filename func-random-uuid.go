package main

import (
	"crypto/rand"
	"math/big"
)

const idChars = "" +
	"0123456789" +
	"abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func newRandom(context *context) (string, error) {
	var bigBase = big.NewInt(idBase)
	var space = new(big.Int).Exp(
		bigBase, big.NewInt(idLength), nil,
	)

	uuid, err := rand.Int(rand.Reader, space)

	if err != nil {
		return "", err
	}

	base62 := ""

	for i := 0; i < idLength; i++ {
		index := new(big.Int).Set(uuid)

		index.Mod(index, bigBase)

		base62 = base62 + string(idChars[index.Uint64()])

		uuid.Div(uuid, bigBase)
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
