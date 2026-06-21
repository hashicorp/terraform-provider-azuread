// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type AccessPackageResourceCatalogAssociationId struct {
	CatalogId string
	OriginId  string
}

func (id AccessPackageResourceCatalogAssociationId) ID() string {
	return fmt.Sprintf("%s/%s", id.CatalogId, id.OriginId)
}

func NewAccessPackageResourceCatalogAssociationID(catalogId, originId string) AccessPackageResourceCatalogAssociationId {
	return AccessPackageResourceCatalogAssociationId{
		CatalogId: catalogId,
		OriginId:  originId,
	}
}

func AccessPackageResourceCatalogAssociationID(idString string) (*AccessPackageResourceCatalogAssociationId, error) {
	parts := strings.SplitN(idString, "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("ID should be in the format {catalogId}/{originId} - but got %q", idString)
	}

	// Only the catalog id is a UUID. originId is the resource's origin identifier, whose format
	// depends on the origin system (a UUID for AadGroup/AadApplication, but a site URL for
	// SharePointOnline), so it is validated only as non-empty.
	if _, err := uuid.ParseUUID(parts[0]); err != nil {
		return nil, fmt.Errorf("specified ID segment #0 (%q) is not a valid UUID: %s", parts[0], err)
	}
	if parts[1] == "" {
		return nil, fmt.Errorf("specified ID segment #1 (originId) must not be empty - got %q", idString)
	}

	return &AccessPackageResourceCatalogAssociationId{
		CatalogId: parts[0],
		OriginId:  parts[1],
	}, nil
}
