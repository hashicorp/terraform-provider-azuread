package types

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
)

type TestResource interface {
	Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error)
}

type TestResourceVerifyingRemoved interface {
	TestResource
	Destroy(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error)
}
