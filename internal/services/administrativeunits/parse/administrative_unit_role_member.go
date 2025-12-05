// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type AdministrativeUnitRoleMemberId struct {
	ObjectSubResourceId
	AdministrativeUnitId   string
	ScopedRoleMembershipId string
}

func NewAdministrativeUnitRoleMemberID(adminUnitID, roleMemberId string) AdministrativeUnitRoleMemberId {
	return AdministrativeUnitRoleMemberId{
		ObjectSubResourceId:    NewObjectSubResourceID(adminUnitID, "roleMember", roleMemberId),
		AdministrativeUnitId:   adminUnitID,
		ScopedRoleMembershipId: roleMemberId,
	}
}

func AdministrativeUnitRoleMemberID(idString string) (*AdministrativeUnitRoleMemberId, error) {
	id, err := AdministrativeUnitRoleMemberObjectSubResourceID(idString, "roleMember")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Member ID: %v", err)
	}

	return &AdministrativeUnitRoleMemberId{
		ObjectSubResourceId:    *id,
		AdministrativeUnitId:   id.objectId,
		ScopedRoleMembershipId: id.subId,
	}, nil
}

func AdministrativeUnitRoleMemberObjectSubResourceID(idString, expectedType string) (*ObjectSubResourceId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Object Resource ID should be in the format {objectId}/{type}/{subId} - but got %q", idString)
	}

	id := ObjectSubResourceId{
		objectId: parts[0],
		Type:     parts[1],
		subId:    parts[2],
	}

	if _, err := uuid.ParseUUID(id.objectId); err != nil {
		return nil, fmt.Errorf("Object ID isn't a valid UUID (%q): %+v", id.objectId, err)
	}

	if id.Type == "" {
		return nil, fmt.Errorf("Type in {objectID}/{type}/{subID} should not be empty")
	}

	if id.Type != expectedType {
		return nil, fmt.Errorf("Type in {objectID}/{type}/{subID} was expected to be %s, got %s", expectedType, id.Type)
	}

	if id.subId == "" {
		return nil, fmt.Errorf("SubId in {objectID}/{type}/{subID} should not be empty")
	}

	return &id, nil
}
