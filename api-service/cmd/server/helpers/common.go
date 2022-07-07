package helpers

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
)

func generateFixedLengthRandomNum(digit int) (fixed_length_num int, err error) {
	if digit < 1 {
		return 0, errors.New("less than 1 digits is not allowed")
	}
	low := int(math.Pow10(digit - 1))
	high := int(math.Pow10(digit) - 1)
	random, err := rand.Int(rand.Reader, big.NewInt(int64(high-low)))
	fixed_length_num = int(random.Int64()) + low
	return
}
