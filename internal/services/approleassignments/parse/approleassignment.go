// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
)

const appRoleAssignment = "appRoleAssignment"

type AppRoleAssignmentId struct {
	ResourceId   string
	AssignmentId string
}

func NewAppRoleAssignmentID(objectId, keyId string) AppRoleAssignmentId {
	return AppRoleAssignmentId{
		ResourceId:   objectId,
		AssignmentId: keyId,
	}
}

func (id AppRoleAssignmentId) String() string {
	return id.ResourceId + "/" + appRoleAssignment + "/" + id.AssignmentId
}

func AppRoleAssignmentID(idString string) (*AppRoleAssignmentId, error) {
	id, err := ObjectSubResourceID(idString, appRoleAssignment)
	if err != nil {
		return nil, fmt.Errorf("unable to parse App Role Assignment ID: %v", err)
	}

	return &AppRoleAssignmentId{
		ResourceId:   id.objectId,
		AssignmentId: id.subId,
	}, nil
}
