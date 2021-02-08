package version

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/version"
	packerVersion "github.com/Bourne-ID/packer/version"
)

var SleepProvisionerVersion *version.PluginVersion

func init() {
	SleepProvisionerVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
