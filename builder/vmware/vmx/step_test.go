package vmx

import (
	"bytes"
	"testing"

	vmwcommon "github.com/Bourne-ID/packer/builder/vmware/common"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(vmwcommon.DriverMock))
	state.Put("ui", &packersdk.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
