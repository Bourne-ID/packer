package lin

import (
	"github.com/Bourne-ID/packer/builder/azure/common/constants"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
)

func SSHHost(state multistep.StateBag) (string, error) {
	host := state.Get(constants.SSHHost).(string)
	return host, nil
}
