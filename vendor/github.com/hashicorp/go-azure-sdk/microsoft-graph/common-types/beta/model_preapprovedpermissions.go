package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreApprovedPermissions interface {
	PreApprovedPermissions() BasePreApprovedPermissionsImpl
}

var _ PreApprovedPermissions = BasePreApprovedPermissionsImpl{}

type BasePreApprovedPermissionsImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the scope of permissions that are included in this condition set. Possible values: all for all permissions,
	// enumerated for a given list of permissions, or allPermissionsOnResourceApp for all permissions from a given API.
	// Required.
	PermissionKind PermissionKind `json:"permissionKind"`

	// The type of permission being granted. Possible values: application for application permissions, or delegated for
	// delegated permissions. Required.
	PermissionType PermissionType `json:"permissionType"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePreApprovedPermissionsImpl) PreApprovedPermissions() BasePreApprovedPermissionsImpl {
	return s
}

var _ PreApprovedPermissions = RawPreApprovedPermissionsImpl{}

// RawPreApprovedPermissionsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPreApprovedPermissionsImpl struct {
	preApprovedPermissions BasePreApprovedPermissionsImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawPreApprovedPermissionsImpl) PreApprovedPermissions() BasePreApprovedPermissionsImpl {
	return s.preApprovedPermissions
}

func UnmarshalPreApprovedPermissionsImplementation(input []byte) (PreApprovedPermissions, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PreApprovedPermissions into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allPreApprovedPermissions") {
		var out AllPreApprovedPermissions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllPreApprovedPermissions: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.allPreApprovedPermissionsOnResourceApp") {
		var out AllPreApprovedPermissionsOnResourceApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllPreApprovedPermissionsOnResourceApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enumeratedPreApprovedPermissions") {
		var out EnumeratedPreApprovedPermissions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnumeratedPreApprovedPermissions: %+v", err)
		}
		return out, nil
	}

	var parent BasePreApprovedPermissionsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePreApprovedPermissionsImpl: %+v", err)
	}

	return RawPreApprovedPermissionsImpl{
		preApprovedPermissions: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
