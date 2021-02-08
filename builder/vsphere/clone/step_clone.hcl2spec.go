// Code generated by "mapstructure-to-hcl2 -type CloneConfig,vAppConfig"; DO NOT EDIT.
package clone

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/Bourne-ID/packer/builder/vsphere/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatCloneConfig is an auto-generated flat version of CloneConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCloneConfig struct {
	Template           *string                 `mapstructure:"template" cty:"template" hcl:"template"`
	DiskSize           *int64                  `mapstructure:"disk_size" cty:"disk_size" hcl:"disk_size"`
	LinkedClone        *bool                   `mapstructure:"linked_clone" cty:"linked_clone" hcl:"linked_clone"`
	Network            *string                 `mapstructure:"network" cty:"network" hcl:"network"`
	MacAddress         *string                 `mapstructure:"mac_address" cty:"mac_address" hcl:"mac_address"`
	Notes              *string                 `mapstructure:"notes" cty:"notes" hcl:"notes"`
	VAppConfig         *FlatvAppConfig         `mapstructure:"vapp" cty:"vapp" hcl:"vapp"`
	DiskControllerType []string                `mapstructure:"disk_controller_type" cty:"disk_controller_type" hcl:"disk_controller_type"`
	Storage            []common.FlatDiskConfig `mapstructure:"storage" cty:"storage" hcl:"storage"`
}

// FlatMapstructure returns a new FlatCloneConfig.
// FlatCloneConfig is an auto-generated flat version of CloneConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*CloneConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCloneConfig)
}

// HCL2Spec returns the hcl spec of a CloneConfig.
// This spec is used by HCL to read the fields of CloneConfig.
// The decoded values from this spec will then be applied to a FlatCloneConfig.
func (*FlatCloneConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"template":             &hcldec.AttrSpec{Name: "template", Type: cty.String, Required: false},
		"disk_size":            &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"linked_clone":         &hcldec.AttrSpec{Name: "linked_clone", Type: cty.Bool, Required: false},
		"network":              &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"mac_address":          &hcldec.AttrSpec{Name: "mac_address", Type: cty.String, Required: false},
		"notes":                &hcldec.AttrSpec{Name: "notes", Type: cty.String, Required: false},
		"vapp":                 &hcldec.BlockSpec{TypeName: "vapp", Nested: hcldec.ObjectSpec((*FlatvAppConfig)(nil).HCL2Spec())},
		"disk_controller_type": &hcldec.AttrSpec{Name: "disk_controller_type", Type: cty.List(cty.String), Required: false},
		"storage":              &hcldec.BlockListSpec{TypeName: "storage", Nested: hcldec.ObjectSpec((*common.FlatDiskConfig)(nil).HCL2Spec())},
	}
	return s
}

// FlatvAppConfig is an auto-generated flat version of vAppConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatvAppConfig struct {
	Properties map[string]string `mapstructure:"properties" cty:"properties" hcl:"properties"`
}

// FlatMapstructure returns a new FlatvAppConfig.
// FlatvAppConfig is an auto-generated flat version of vAppConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*vAppConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatvAppConfig)
}

// HCL2Spec returns the hcl spec of a vAppConfig.
// This spec is used by HCL to read the fields of vAppConfig.
// The decoded values from this spec will then be applied to a FlatvAppConfig.
func (*FlatvAppConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"properties": &hcldec.AttrSpec{Name: "properties", Type: cty.Map(cty.String), Required: false},
	}
	return s
}
