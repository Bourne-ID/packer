package docker

import (
	"testing"

	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func TestBuilder_implBuilder(t *testing.T) {
	var _ packersdk.Builder = new(Builder)
}
