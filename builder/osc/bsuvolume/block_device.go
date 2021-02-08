package bsuvolume

import (
	osccommon "github.com/Bourne-ID/packer/builder/osc/common"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/template/interpolate"
)

type BlockDevice struct {
	osccommon.BlockDevice `mapstructure:"-,squash"`
	Tags                  osccommon.TagMap `mapstructure:"tags"`
}

func commonBlockDevices(mappings []BlockDevice, ctx *interpolate.Context) (osccommon.BlockDevices, error) {
	result := make([]osccommon.BlockDevice, len(mappings))

	for i, mapping := range mappings {
		interpolateBlockDev, err := interpolate.RenderInterface(&mapping.BlockDevice, ctx)
		if err != nil {
			return osccommon.BlockDevices{}, err
		}
		result[i] = *interpolateBlockDev.(*osccommon.BlockDevice)
	}

	return osccommon.BlockDevices{
		LaunchBlockDevices: osccommon.LaunchBlockDevices{
			LaunchMappings: result,
		},
	}, nil
}
