//go:generate mapstructure-to-hcl2 -type Config

package triton

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/common"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/communicator"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	AccessConfig        `mapstructure:",squash"`
	SourceMachineConfig `mapstructure:",squash"`
	TargetImageConfig   `mapstructure:",squash"`

	Comm communicator.Config `mapstructure:",squash"`

	ctx interpolate.Context
}
