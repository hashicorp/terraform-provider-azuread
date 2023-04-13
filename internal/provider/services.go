package provider

import (
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
	"github.com/hashicorp/terraform-provider-azuread/internal/services/users"
)

func SupportedServices() []ServiceRegistration {
	return []ServiceRegistration{
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
		users.Registration{},
	}
}
