package uuid

import (
	"github.com/wayn3h0/gop/uuid/internal/dcesecurity"
)

// Domain represents the identifier for a local domain.
type Domain byte

const (
	// DomainUser represents POSIX UID domain.
	DomainUser = Domain(dcesecurity.User)
	// DomainGroup represents POSIX GID domain.
	DomainGroup = Domain(dcesecurity.Group)
)
