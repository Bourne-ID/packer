package common

import (
	"bytes"
	"strings"

	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func basicStateBag(errorBuffer *strings.Builder) *multistep.BasicStateBag {
	state := new(multistep.BasicStateBag)
	state.Put("ui", &packersdk.BasicUi{
		Reader:      new(bytes.Buffer),
		Writer:      new(bytes.Buffer),
		ErrorWriter: errorBuffer,
	})
	return state
}
