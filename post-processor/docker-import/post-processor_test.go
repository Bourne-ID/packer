package dockerimport

import (
	"testing"

	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func TestPostProcessor_ImplementsPostProcessor(t *testing.T) {
	var _ packersdk.PostProcessor = new(PostProcessor)
}
