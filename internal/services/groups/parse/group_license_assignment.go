// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type GroupLicenseAssignmentId struct {
	ObjectSubResourceId
	GroupId string
	SKUId   string
}

func NewGroupLicenseAssignmentID(groupId, skuId string) GroupLicenseAssignmentId {
	return GroupLicenseAssignmentId{
		ObjectSubResourceId: NewObjectSubResourceID(groupId, "license", skuId),
		GroupId:             groupId,
		SKUId:               skuId,
	}
}

func GroupLicenseAssignmentID(idString string) (*GroupLicenseAssignmentId, error) {
	id, err := ObjectSubResourceID(idString, "license")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Group License Assignment ID: %v", err)
	}

	return &GroupLicenseAssignmentId{
		ObjectSubResourceId: *id,
		GroupId:             id.objectId,
		SKUId:               id.subId,
	}, nil
}
