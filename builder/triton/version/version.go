package version

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/version"
	packerVersion "github.com/Bourne-ID/packer/version"
)

var TritonPluginVersion *version.PluginVersion

func init() {
	TritonPluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
