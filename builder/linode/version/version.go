package version

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/version"
	packerVersion "github.com/Bourne-ID/packer/version"
)

var LinodePluginVersion *version.PluginVersion

func init() {
	LinodePluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
