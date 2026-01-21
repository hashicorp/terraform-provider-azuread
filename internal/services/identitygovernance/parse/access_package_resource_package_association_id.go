// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type AccessPackageResourcePackageAssociationId struct {
	AccessPackageId     string
	ResourceRoleScopeId string
	OriginId            string
	AccessType          string
}

func (id AccessPackageResourcePackageAssociationId) ID() string {
	return fmt.Sprintf("%s/%s/%s/%s", id.AccessPackageId, id.ResourceRoleScopeId, id.OriginId, id.AccessType)
}

func NewAccessPackageResourcePackageAssociationID(catalogId, resourceRoleScopeId, originId, accessType string) AccessPackageResourcePackageAssociationId {
	return AccessPackageResourcePackageAssociationId{
		AccessPackageId:     catalogId,
		ResourceRoleScopeId: resourceRoleScopeId,
		OriginId:            originId,
		AccessType:          accessType,
	}
}

func AccessPackageResourcePackageAssociationID(idString string) (*AccessPackageResourcePackageAssociationId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("ID should be in the format {accessPackageId}/{resourcePackageAssociationId}/{originId}/{accessType} - but got %q", idString)
	}

	for i, p := range parts {
		if i == 0 || i == 2 {
			if _, err := uuid.ParseUUID(p); err != nil {
				return nil, fmt.Errorf("specified ID segment #%d (%q) is not a valid UUID: %s", i, p, err)
			}
		}
	}

	return &AccessPackageResourcePackageAssociationId{
		AccessPackageId:     parts[0],
		ResourceRoleScopeId: parts[1],
		OriginId:            parts[2],
		AccessType:          parts[3],
	}, nil
}
