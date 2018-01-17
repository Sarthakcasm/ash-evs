package main

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/Savasw/gophercloud/openstack/networking/v1/vpcs"
	"fmt"
)

func main() {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://iam.eu-de.otc.t-systems.com/v3",
		DomainName:       "OTC00000000001000010501",
		Username:         "lizhonghua",
		Password:         "slob@123",
		TenantID:         "87a56a48977e42068f70ad3280c50f0e",
	}

	// to get provider client with token_id
	provider, err := openstack.AuthenticatedClient(opts)
	//fmt.Println("\n\nProvider : ", provider);

	fmt.Print("\nerror : ", err);

	//creating client

	opp := gophercloud.EndpointOpts{}        // specify search criteria
	client, err1 := openstack.NewVpcV1(provider,opp)

	fmt.Print("opp", opp)
	fmt.Print("Error", err1)

	// listing VPCs

	opts1 := vpcs.ListOpts{}    //struct
	vpc , err2 := vpcs.List(client, opts1)

	fmt.Print("opts1",opts1)
	fmt.Print("error",err2)
	fmt.Println("\n\nList of existing VPC Available.\n\n")
	for _, v :=range vpc{
		fmt.Print("vpc_ID :",v.ID,"\t")
		fmt.Print("Name :",v.Name,"\t\t")
		fmt.Print("CIDR :",v.CIDR,"\t\t")
		fmt.Println("status :",v.Status)
	}

	// method to get info of single vpc

	fmt.Println("unique Id")
	fmt.Println()
	vpc_id := "3b9740a0-b44d-48f0-84ee-42eb166e54f7"
	Result := vpcs.Get(client,vpc_id)
	fmt.Println(Result.Body)



	//// creating new VPC
	//
	//createOpts := vpcs.CreateOpts{
	//	Name:         "New_born_vpc1",
	//	CIDR:         "192.168.0.0/24",
	//}
	//
	//vpc1, err := vpcs.Create(client, createOpts).Extract()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Vpc result",vpc1)
	//
	// // Example to Update a Vpc
	//
	//vpcID := ""
	//
	//updateOpts := vpcs.UpdateOpts{
	//	Name:         "vpc_updated",
	//	CIDR:         "192.168.0.0/23",
	//}
	//
	//vpc2, err1 := vpcs.Update(client, vpcID, updateOpts).Extract()
	//if err1 != nil {
	//	panic(err1)
	//}
	//
	//fmt.Println("updated vpc : ",vpc2)
	//fmt.Println("updated vpc : ",err1)
	//
	//// Example to Delete a Vpc
	//
	//vpcID1 := ""
	//err = vpcs.Delete(client, vpcID1).ExtractErr()
	//fmt.Println(err)
	//if err != nil {
	//	panic(err)
	//}
}
