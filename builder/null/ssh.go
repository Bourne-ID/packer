package null

import (
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
)

func CommHost(host string) func(multistep.StateBag) (string, error) {
	return func(state multistep.StateBag) (string, error) {
		return host, nil
	}
}
