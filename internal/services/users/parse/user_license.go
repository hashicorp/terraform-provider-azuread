// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type UserLicenseId struct {
	UserId string
	SkuId  string
}

func NewUserLicenseID(userId, skuId string) UserLicenseId {
	return UserLicenseId{
		UserId: userId,
		SkuId:  skuId,
	}
}

func (id UserLicenseId) String() string {
	return fmt.Sprintf("%s/license/%s", id.UserId, id.SkuId)
}

func UserLicenseID(idString string) (*UserLicenseId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("User License ID should be in the format {userId}/license/{skuId} - but got %q", idString)
	}

	if parts[1] != "license" {
		return nil, fmt.Errorf("Type in {userId}/license/{skuId} was expected to be \"license\", got %q", parts[1])
	}

	id := UserLicenseId{
		UserId: parts[0],
		SkuId:  parts[2],
	}

	if _, err := uuid.ParseUUID(id.UserId); err != nil {
		return nil, fmt.Errorf("User ID isn't a valid UUID (%q): %+v", id.UserId, err)
	}

	if _, err := uuid.ParseUUID(id.SkuId); err != nil {
		return nil, fmt.Errorf("SKU ID isn't a valid UUID (%q): %+v", id.SkuId, err)
	}

	return &id, nil
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
