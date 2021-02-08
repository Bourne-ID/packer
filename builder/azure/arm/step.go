package arm

import (
	"github.com/Bourne-ID/packer/builder/azure/common/constants"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/multistep"
)

func processStepResult(
	err error, sayError func(error), state multistep.StateBag) multistep.StepAction {

	if err != nil {
		state.Put(constants.Error, err)
		sayError(err)

		return multistep.ActionHalt
	}

	return multistep.ActionContinue

}
