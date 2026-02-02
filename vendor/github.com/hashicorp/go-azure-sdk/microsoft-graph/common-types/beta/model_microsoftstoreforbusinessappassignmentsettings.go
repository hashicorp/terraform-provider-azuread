package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = MicrosoftStoreForBusinessAppAssignmentSettings{}

type MicrosoftStoreForBusinessAppAssignmentSettings struct {
	// Whether or not to use device execution context for Microsoft Store for Business mobile app.
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MicrosoftStoreForBusinessAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftStoreForBusinessAppAssignmentSettings{}

func (s MicrosoftStoreForBusinessAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftStoreForBusinessAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftStoreForBusinessAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftStoreForBusinessAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftStoreForBusinessAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftStoreForBusinessAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
