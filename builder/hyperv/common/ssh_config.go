package common

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/communicator"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/template/interpolate"
)

type SSHConfig struct {
	Comm communicator.Config `mapstructure:",squash"`
}

func (c *SSHConfig) Prepare(ctx *interpolate.Context) []error {
	return c.Comm.Prepare(ctx)
}
