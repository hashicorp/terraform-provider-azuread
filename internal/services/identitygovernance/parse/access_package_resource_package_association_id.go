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
	// RoleOriginId is the origin ID of the resource role (for application AppRoles/OAuth2 scopes).
	// When set, AccessType should be empty.
	RoleOriginId string
}

func (id AccessPackageResourcePackageAssociationId) ID() string {
	roleIdentifier := id.AccessType
	if id.RoleOriginId != "" {
		roleIdentifier = id.RoleOriginId
	}
	return fmt.Sprintf("%s/%s/%s/%s", id.AccessPackageId, id.ResourceRoleScopeId, id.OriginId, roleIdentifier)
}

func NewAccessPackageResourcePackageAssociationID(catalogId, resourceRoleScopeId, originId, accessType string) AccessPackageResourcePackageAssociationId {
	return AccessPackageResourcePackageAssociationId{
		AccessPackageId:     catalogId,
		ResourceRoleScopeId: resourceRoleScopeId,
		OriginId:            originId,
		AccessType:          accessType,
	}
}

func NewAccessPackageResourcePackageAssociationIDWithRoleOrigin(accessPackageId, resourceRoleScopeId, originId, roleOriginId string) AccessPackageResourcePackageAssociationId {
	return AccessPackageResourcePackageAssociationId{
		AccessPackageId:     accessPackageId,
		ResourceRoleScopeId: resourceRoleScopeId,
		OriginId:            originId,
		RoleOriginId:        roleOriginId,
	}
}

// knownAccessTypes are the group-based access types stored in the ID's 4th segment.
var knownAccessTypes = map[string]bool{
	"Member":          true,
	"Owner":           true,
	"Eligible Member": true,
	"Eligible Owner":  true,
}

func AccessPackageResourcePackageAssociationID(idString string) (*AccessPackageResourcePackageAssociationId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("ID should be in the format {accessPackageId}/{resourcePackageAssociationId}/{originId}/{roleIdentifier} - but got %q", idString)
	}

	for i, p := range parts {
		if i == 0 || i == 2 {
			if _, err := uuid.ParseUUID(p); err != nil {
				return nil, fmt.Errorf("specified ID segment #%d (%q) is not a valid UUID: %s", i, p, err)
			}
		}
	}

	roleIdentifier := parts[3]
	if roleIdentifier == "" {
		return nil, fmt.Errorf("specified ID segment #3 (roleIdentifier) is empty")
	}

	result := &AccessPackageResourcePackageAssociationId{
		AccessPackageId:     parts[0],
		ResourceRoleScopeId: parts[1],
		OriginId:            parts[2],
	}

	if knownAccessTypes[roleIdentifier] {
		result.AccessType = roleIdentifier
	} else {
		result.RoleOriginId = roleIdentifier
	}

	return result, nil
}
