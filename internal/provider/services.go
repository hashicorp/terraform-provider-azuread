package provider

import (
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/applications"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/domains"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/groups"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/serviceprincipals"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/users"
)

func SupportedServices() []ServiceRegistration {
	return []ServiceRegistration{
		applications.Registration{},
		domains.Registration{},
		groups.Registration{},
		serviceprincipals.Registration{},
		users.Registration{},
	}
}
