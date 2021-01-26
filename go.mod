module github.com/hashicorp/terraform-provider-azuread

require (
	github.com/Azure/azure-sdk-for-go v47.1.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.10
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/google/uuid v1.1.1
	github.com/hashicorp/go-azure-helpers v0.13.1
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-uuid v1.0.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.3.0
	github.com/manicminer/hamilton v0.6.0
)

go 1.16

//replace github.com/manicminer/hamilton => /Users/tom/go/src/github.com/manicminer/hamilton
