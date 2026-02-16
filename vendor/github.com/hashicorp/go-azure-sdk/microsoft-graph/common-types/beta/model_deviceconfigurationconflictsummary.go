package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationConflictSummary{}

type DeviceConfigurationConflictSummary struct {
	// The set of policies in conflict with the given setting
	ConflictingDeviceConfigurations *[]SettingSource `json:"conflictingDeviceConfigurations,omitempty"`

	// The set of settings in conflict with the given policies
	ContributingSettings *[]string `json:"contributingSettings,omitempty"`

	// The count of checkins impacted by the conflicting policies and settings
	DeviceCheckinsImpacted *int64 `json:"deviceCheckinsImpacted,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceConfigurationConflictSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationConflictSummary{}

func (s DeviceConfigurationConflictSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationConflictSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationConflictSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationConflictSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationConflictSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationConflictSummary: %+v", err)
	}

	return encoded, nil
}
