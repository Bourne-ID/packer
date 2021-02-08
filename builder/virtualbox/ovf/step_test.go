package ovf

import (
	"bytes"
	"testing"

	vboxcommon "github.com/Bourne-ID/packer/builder/virtualbox/common"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(vboxcommon.DriverMock))
	state.Put("ui", &packersdk.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
