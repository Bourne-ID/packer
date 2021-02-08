package yandex

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
)

func CommHost(state multistep.StateBag) (string, error) {
	ipAddress := state.Get("instance_ip").(string)
	return ipAddress, nil
}
