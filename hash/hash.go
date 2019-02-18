package hash

import (
	"bytes"
	"hash"

	"github.com/wayn3h0/gop/errors"
)

// Compute computes the hash value for given data.
func Compute(hash hash.Hash, data []byte, more ...[]byte) ([]byte, error) {
	_, err := hash.Write([]byte(data))
	if err != nil {
		return nil, errors.Wrap(err, "hash: could not compute hash value of main data")
	}

	for i, v := range more {
		_, err := hash.Write([]byte(v))
		if err != nil {
			return nil, errors.Wrapf(err, "hash: could not compute hash value for additional data [index: %d]", i)
		}
	}

	return hash.Sum(nil), nil
}

// Verify reports whether the value equals to the hash value of given data.
func Verify(hash hash.Hash, value []byte, data []byte, more ...[]byte) (bool, error) {
	val, err := Compute(hash, data, more...)
	if err != nil {
		return false, err
	}

	return bytes.Equal(val, value), nil
}
