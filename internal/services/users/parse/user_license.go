// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type UserLicenseId struct {
	ObjectSubResourceId
	UserId string
	SkuId  string
}

func NewUserLicenseID(userId, skuId string) UserLicenseId {
	return UserLicenseId{
		ObjectSubResourceId: NewObjectSubResourceID(userId, "license", skuId),
		UserId:              userId,
		SkuId:               skuId,
	}
}

func (id UserLicenseId) ID() string {
	return id.ObjectSubResourceId.String()
}

func UserLicenseID(idString string) (*UserLicenseId, error) {
	id, err := ObjectSubResourceID(idString, "license")
	if err != nil {
		return nil, fmt.Errorf("unable to parse User License ID: %v", err)
	}

	return &UserLicenseId{
		ObjectSubResourceId: *id,
		UserId:              id.objectId,
		SkuId:               id.subId,
	}, nil
}

func ValidateUserLicenseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := UserLicenseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}
