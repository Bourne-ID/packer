package common

import "github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"

func IsStateCancelled(stateBag multistep.StateBag) bool {
	_, ok := stateBag.GetOk(multistep.StateCancelled)
	return ok
}
