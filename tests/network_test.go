package tests

import (
	"fmt"
	"context"
	"testing"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-azure-example using Terratest.
func TestNetworkDeployment(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../tf/",
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	vnetName := terraform.Output(t, terraformOptions, "vnet_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	azureSubscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	fmt.Printf("Subscription ID = %s", azureSubscriptionID)
	
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		t.Fatal("Cannot get an Azure Authorizer")
	}

	vnetClient := network.NewVirtualNetworksClient(azureSubscriptionID)
	vnetClient.Authorizer = authorizer

	virtualNetwork, err := vnetClient.Get(context.Background(), resourceGroupName, vnetName, "")
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, "The virtual network does not exists")
		return
	}

	expectedSubnetsCount := 1
	actualSubnetsCount := len(*virtualNetwork.Subnets)

	if !assert.Equal(t, expectedSubnetsCount, actualSubnetsCount, "The expected number of subnet is 1") {
		return
	}

	subnet := (*virtualNetwork.Subnets)[0]
	
	expectedSubnetName := "default"
	actualSubnetName := subnet.Name

	assert.Equal(t, expectedSubnetName, *actualSubnetName, "The expected subnet name is 'default'")
}