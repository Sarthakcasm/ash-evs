package main

import (
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/Ashishraw/gophercloud/openstack/networking/v2/evs"
)

func main() {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://iam.eu-de.otc.t-systems.com/v3",
		DomainName:       "OTC00000000001000010501",
		Username:         "lizhonghua",
		Password:         "slob@123",
		TenantID:         "87a56a48977e42068f70ad3280c50f0e",
	}

	// To get provider client with token_Id
	provider ,err := openstack.AuthenticatedClient(opts)
	fmt.Print(err,"\t below List of Volumes\n")

	// to get Service client
	client, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{

	})

	// To get list of all volumes

	opp := evs.ListOpts{}
	lists , err := evs.List(client,opp)
	for _, v := range lists{
		fmt.Println(v)
	}

	 // To get single volume info
	 var volume_id string
	 fmt.Println("\n\nTo get details of EVS enter volume Id")
	 fmt.Scan(&volume_id)
	 Result := evs.Get(client,volume_id)
	 fmt.Println(Result)

	}


