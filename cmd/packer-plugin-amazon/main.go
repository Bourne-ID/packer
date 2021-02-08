package main

import (
	"fmt"
	"os"

	"github.com/Bourne-ID/packer/builder/amazon/ebs"
	"github.com/Bourne-ID/packer/builder/amazon/ebssurrogate"
	"github.com/Bourne-ID/packer/builder/amazon/ebsvolume"
	"github.com/Bourne-ID/packer/builder/osc/chroot"
	"github.com/Bourne-ID/packer/packer-plugin-sdk/plugin"
	amazonimport "github.com/Bourne-ID/packer/post-processor/amazon-import"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("ebs", new(ebs.Builder))
	pps.RegisterBuilder("chroot", new(chroot.Builder))
	pps.RegisterBuilder("ebssurrogate", new(ebssurrogate.Builder))
	pps.RegisterBuilder("ebsvolume", new(ebsvolume.Builder))
	pps.RegisterPostProcessor("import", new(amazonimport.PostProcessor))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
