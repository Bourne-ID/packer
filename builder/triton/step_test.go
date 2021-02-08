package triton

import (
	"bytes"
	"testing"

	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("config", testConfig(t))
	state.Put("driver", &DriverMock{})
	state.Put("hook", &packersdk.MockHook{})
	state.Put("ui", &packersdk.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
