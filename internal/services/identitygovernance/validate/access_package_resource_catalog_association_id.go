package validate

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
)

func AccessPackageResourceCatalogAssociationID(input string) (err error) {
	_, err = parse.AccessPackageResourceCatalogAssociationID(input)
	return
}
