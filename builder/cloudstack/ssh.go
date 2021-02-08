package cloudstack

import (
	"fmt"

	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
)

func commPort(state multistep.StateBag) (int, error) {
	commPort, hasPort := state.Get("commPort").(int)
	if !hasPort {
		return 0, fmt.Errorf("Failed to retrieve communication port")
	}

	return commPort, nil
}
