package chroot

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
)

// Cleanup is an interface that some steps implement for early cleanup.
type Cleanup interface {
	CleanupFunc(multistep.StateBag) error
}
