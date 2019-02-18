package md5

import (
	"crypto/md5"

	"github.com/wayn3h0/gop/errors"
	"github.com/wayn3h0/gop/uuid/internal"
)

// NewUUID returns a new name-based uses SHA-1 hashing uuid.
func NewUUID(namespace, name string) ([]byte, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(namespace))
	if err != nil {
		return nil, errors.Wrapf(err, "uuid: could not compute hash value for namespace %q", namespace)
	}
	_, err = hash.Write([]byte(name))
	if err != nil {
		return nil, errors.Wrapf(err, "uuid: could not compute hash value for name %q", name)
	}

	sum := hash.Sum(nil)

	uuid := make([]byte, 16)
	copy(uuid, sum)

	// set version(v3)
	internal.SetVersion(uuid, internal.VersionNameBasedMD5)
	// set layout(RFC4122)
	internal.SetLayout(uuid, internal.LayoutRFC4122)

	return uuid, nil
}
