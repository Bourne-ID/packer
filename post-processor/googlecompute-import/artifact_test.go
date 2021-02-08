package googlecomputeimport

import (
	"testing"

	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func TestArtifact_ImplementsArtifact(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packersdk.Artifact); !ok {
		t.Fatalf("Artifact should be a Artifact")
	}
}
