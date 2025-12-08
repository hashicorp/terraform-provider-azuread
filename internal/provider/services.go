// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/sdk"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/administrativeunits"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/applications"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/approleassignments"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/conditionalaccess"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/directoryobjects"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/directoryroles"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/domains"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/groups"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/identitygovernance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/invitations"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/policies"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/serviceprincipals"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/synchronization"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/userflows"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/users"
)

//go:generate go run ../tools/generator-services/main.go -path=../../

func SupportedTypedServices() []sdk.TypedServiceRegistration {
	return []sdk.TypedServiceRegistration{
		applications.Registration{},
		directoryroles.Registration{},
		domains.Registration{},
		policies.Registration{},
		identitygovernance.Registration{},
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
		synchronization.Registration{},
		userflows.Registration{},
		users.Registration{},
	}
}
