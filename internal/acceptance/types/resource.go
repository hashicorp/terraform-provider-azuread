// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"context"

	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type TestResource interface {
	Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error)
}

type TestResourceVerifyingRemoved interface {
	TestResource
	Destroy(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error)
}
