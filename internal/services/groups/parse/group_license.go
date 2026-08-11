// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type GroupLicenseId struct {
	ObjectSubResourceId
	GroupId string
	SkuId   string
}

func NewGroupLicenseID(groupId, skuId string) GroupLicenseId {
	return GroupLicenseId{
		ObjectSubResourceId: NewObjectSubResourceID(groupId, "license", skuId),
		GroupId:             groupId,
		SkuId:               skuId,
	}
}

func (id GroupLicenseId) ID() string {
	return id.ObjectSubResourceId.String()
}

func GroupLicenseID(idString string) (*GroupLicenseId, error) {
	id, err := ObjectSubResourceID(idString, "license")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Group License ID: %v", err)
	}

	return &GroupLicenseId{
		ObjectSubResourceId: *id,
		GroupId:             id.objectId,
		SkuId:               id.subId,
	}, nil
}

func ValidateGroupLicenseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := GroupLicenseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}
