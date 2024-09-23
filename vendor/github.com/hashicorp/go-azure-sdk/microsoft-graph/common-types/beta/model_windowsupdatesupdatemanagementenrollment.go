package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesUpdatableAssetEnrollment = WindowsUpdatesUpdateManagementEnrollment{}

type WindowsUpdatesUpdateManagementEnrollment struct {
	UpdateCategory *WindowsUpdatesUpdateCategory `json:"updateCategory,omitempty"`

	// Fields inherited from WindowsUpdatesUpdatableAssetEnrollment

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesUpdateManagementEnrollment) WindowsUpdatesUpdatableAssetEnrollment() BaseWindowsUpdatesUpdatableAssetEnrollmentImpl {
	return BaseWindowsUpdatesUpdatableAssetEnrollmentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesUpdateManagementEnrollment{}

func (s WindowsUpdatesUpdateManagementEnrollment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesUpdateManagementEnrollment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesUpdateManagementEnrollment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesUpdateManagementEnrollment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.updateManagementEnrollment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesUpdateManagementEnrollment: %+v", err)
	}

	return encoded, nil
}
