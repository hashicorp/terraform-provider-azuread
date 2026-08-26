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
	// A role identified by a value containing "/" (e.g. a SharePointOnline role URL) cannot be
	// encoded in the legacy 4-segment "/"-delimited ID. For those, originId/accessType are left
	// empty and recovered from the API on read, so the ID is the 2-segment
	// {accessPackageId}/{resourceRoleScopeId}.
	if id.OriginId == "" && id.AccessType == "" {
		return fmt.Sprintf("%s/%s", id.AccessPackageId, id.ResourceRoleScopeId)
	}
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

	switch len(parts) {
	case 2:
		// Current format: {accessPackageId}/{resourceRoleScopeId}. originId and accessType are
		// recovered from the API on read, which supports role identifiers containing "/"
		// (e.g. SharePointOnline role URLs).
		if _, err := uuid.ParseUUID(parts[0]); err != nil {
			return nil, fmt.Errorf("specified ID segment #0 (%q) is not a valid UUID: %s", parts[0], err)
		}
		return &AccessPackageResourcePackageAssociationId{
			AccessPackageId:     parts[0],
			ResourceRoleScopeId: parts[1],
		}, nil

	case 4:
		// Legacy format: {accessPackageId}/{resourcePackageAssociationId}/{originId}/{accessType},
		// still emitted for AadGroup associations whose segments never contain "/".
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

	default:
		return nil, fmt.Errorf("ID should be {accessPackageId}/{resourceRoleScopeId} or the legacy {accessPackageId}/{resourcePackageAssociationId}/{originId}/{accessType} - but got %q", idString)
	}
}
