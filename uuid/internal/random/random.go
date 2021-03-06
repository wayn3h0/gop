package random

import (
	"crypto/rand"

	"github.com/wayn3h0/gop/errors"
	"github.com/wayn3h0/gop/uuid/internal"
)

// NewUUID returns a new randomly uuid.
func NewUUID() ([]byte, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid[:])
	if err != nil {
		return nil, errors.Wrap(err, "uuid: could not generate random bytes")
	}
	if n != len(uuid) {
		return nil, errors.New("uuid: could not generate random bytes with 16 length")
	}

	// set version(v4)
	internal.SetVersion(uuid, internal.VersionRandom)
	// set layout(RFC4122)
	internal.SetLayout(uuid, internal.LayoutRFC4122)

	return uuid, nil
}
