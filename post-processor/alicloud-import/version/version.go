package version

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/version"
	packerVersion "github.com/Bourne-ID/packer/version"
)

var AlicloudImportPluginVersion *version.PluginVersion

func init() {
	AlicloudImportPluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
