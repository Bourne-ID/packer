package main

import (
	"fmt"

	packersdk "github.com/Bourne-ID/packer/packer-plugin-sdk/packer"
)

func openTTY() (packersdk.TTY, error) {
	return nil, fmt.Errorf("no TTY available on solaris")
}
