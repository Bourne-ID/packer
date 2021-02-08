package chroot

import "github.com/Bourne-ID/packer/packer-plugin-sdk/template/interpolate"

type interpolateContextProvider interface {
	GetContext() interpolate.Context
}
