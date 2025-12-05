// Copyright IBM Corp. 2019, 2025
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

	for i, p := range parts {
		if _, err := uuid.ParseUUID(p); err != nil {
			return nil, fmt.Errorf("specified ID segment #%d (%q) is not a valid UUID: %s", i, p, err)
		}
	}

	return &AccessPackageResourceCatalogAssociationId{
		CatalogId: parts[0],
		OriginId:  parts[1],
	}, nil
}
