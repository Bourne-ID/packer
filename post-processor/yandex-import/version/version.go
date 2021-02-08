package version

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/version"
	packerVersion "github.com/Bourne-ID/packer/version"
)

var YandexImportPluginVersion *version.PluginVersion

func init() {
	YandexImportPluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
