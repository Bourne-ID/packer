package chroot

import "github.com/Bourne-ID/packer/builder/azure/common/client"

func withMetadataStub(f func()) {
	mdc := client.DefaultMetadataClient
	defer func() { client.DefaultMetadataClient = mdc }()
	client.DefaultMetadataClient = client.MetadataClientStub{
		ComputeInfo: client.ComputeInfo{
			SubscriptionID:    "testSubscriptionID",
			ResourceGroupName: "testResourceGroup",
			Name:              "testVM",
			Location:          "testLocation",
		},
	}

	f()
}
