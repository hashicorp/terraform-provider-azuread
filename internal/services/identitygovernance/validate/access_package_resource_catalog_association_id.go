// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/identitygovernance/parse"
)

func AccessPackageResourceCatalogAssociationID(input string) (err error) {
	_, err = parse.AccessPackageResourceCatalogAssociationID(input)
	return
}
