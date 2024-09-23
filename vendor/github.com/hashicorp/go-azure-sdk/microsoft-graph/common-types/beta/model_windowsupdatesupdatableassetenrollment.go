package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdatableAssetEnrollment interface {
	WindowsUpdatesUpdatableAssetEnrollment() BaseWindowsUpdatesUpdatableAssetEnrollmentImpl
}

var _ WindowsUpdatesUpdatableAssetEnrollment = BaseWindowsUpdatesUpdatableAssetEnrollmentImpl{}

type BaseWindowsUpdatesUpdatableAssetEnrollmentImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesUpdatableAssetEnrollmentImpl) WindowsUpdatesUpdatableAssetEnrollment() BaseWindowsUpdatesUpdatableAssetEnrollmentImpl {
	return s
}

var _ WindowsUpdatesUpdatableAssetEnrollment = RawWindowsUpdatesUpdatableAssetEnrollmentImpl{}

// RawWindowsUpdatesUpdatableAssetEnrollmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesUpdatableAssetEnrollmentImpl struct {
	windowsUpdatesUpdatableAssetEnrollment BaseWindowsUpdatesUpdatableAssetEnrollmentImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawWindowsUpdatesUpdatableAssetEnrollmentImpl) WindowsUpdatesUpdatableAssetEnrollment() BaseWindowsUpdatesUpdatableAssetEnrollmentImpl {
	return s.windowsUpdatesUpdatableAssetEnrollment
}

func UnmarshalWindowsUpdatesUpdatableAssetEnrollmentImplementation(input []byte) (WindowsUpdatesUpdatableAssetEnrollment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesUpdatableAssetEnrollment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.updateManagementEnrollment") {
		var out WindowsUpdatesUpdateManagementEnrollment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesUpdateManagementEnrollment: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesUpdatableAssetEnrollmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesUpdatableAssetEnrollmentImpl: %+v", err)
	}

	return RawWindowsUpdatesUpdatableAssetEnrollmentImpl{
		windowsUpdatesUpdatableAssetEnrollment: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
