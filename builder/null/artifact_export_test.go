package null

import (
	"testing"

	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func TestNullArtifact(t *testing.T) {
	var _ packersdk.Artifact = new(NullArtifact)
}
