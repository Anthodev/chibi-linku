package helpers

import (
	"math/big"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Base62Encode(w string) string {
	var result string

	base := big.NewInt(62)
	zero := big.NewInt(0)

	wInt := big.NewInt(0).SetBytes([]byte(w))

	remainder := &big.Int{}
	quotient := new(big.Int).Set(wInt)

	for quotient.Cmp(zero) > 0 {
		quotient, remainder = quotient.DivMod(quotient, base, remainder)
		result = string(alphabet[remainder.Int64()]) + result
	}

	return result
}

func Base62Decode(e string) string {
	base := big.NewInt(62)
	result := big.NewInt(0)

	for _, c := range e {
		index := strings.IndexRune(alphabet, c)
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(index)))
	}

	return string(result.Bytes())
}
