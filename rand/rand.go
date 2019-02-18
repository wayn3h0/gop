package rand

import (
	"crypto/rand"

	"github.com/wayn3h0/gop/errors"
)

const (
	// Alphabet represents charset of alphabet.
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// Numerals represents charset of numerals.
	Numerals = "0123456789"
	// AlphabetAndNumerals represents charset alphabet plus numberals.
	AlphabetAndNumerals = Alphabet + Numerals
)

// Bytes returns random bytes picks N (length) from given charset.
func Bytes(charset string, length int) ([]byte, error) {
	if length < 1 {
		return nil, errors.Newf("rand: length of random bytes `%d` is invalid", length)
	}

	chars := charset
	if len(chars) == 0 {
		chars = AlphabetAndNumerals
	}
	charsLength := len(chars)

	var bytes = make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, errors.Wrap(err, "rand: could not generate random bytes")
	}

	for i, v := range bytes {
		bytes[i] = chars[v%byte(charsLength)]
	}

	return bytes, nil
}

// String returns random string picks N (length) from given charset.
func String(charset string, length int) (string, error) {
	buf, err := Bytes(charset, length)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
