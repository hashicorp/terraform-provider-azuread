// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryobjects"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/domains"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/invitations"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/userflows"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/users"
)

//go:generate go run ../tools/generator-services/main.go -path=../../

func SupportedTypedServices() []sdk.TypedServiceRegistration {
	return []sdk.TypedServiceRegistration{
		applications.Registration{},
		directoryroles.Registration{},
		domains.Registration{},
		serviceprincipals.Registration{},
	}
}

func SupportedUntypedServices() []sdk.UntypedServiceRegistration {
	return []sdk.UntypedServiceRegistration{
		administrativeunits.Registration{},
		applications.Registration{},
		approleassignments.Registration{},
		conditionalaccess.Registration{},
		directoryobjects.Registration{},
		directoryroles.Registration{},
		domains.Registration{},
		groups.Registration{},
		identitygovernance.Registration{},
		invitations.Registration{},
		policies.Registration{},
		serviceprincipals.Registration{},
		userflows.Registration{},
		users.Registration{},
	}
}
