package validate

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
)

func AccessPackageResourcePackageAssociationID(input string) (err error) {
	_, err = parse.AccessPackageResourcePackageAssociationID(input)
	return
}
