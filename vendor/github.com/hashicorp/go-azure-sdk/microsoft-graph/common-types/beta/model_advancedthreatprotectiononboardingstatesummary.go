package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AdvancedThreatProtectionOnboardingStateSummary{}

type AdvancedThreatProtectionOnboardingStateSummary struct {
	AdvancedThreatProtectionOnboardingDeviceSettingStates *[]AdvancedThreatProtectionOnboardingDeviceSettingState `json:"advancedThreatProtectionOnboardingDeviceSettingStates,omitempty"`

	// Number of compliant devices
	CompliantDeviceCount *int64 `json:"compliantDeviceCount,omitempty"`

	// Number of conflict devices
	ConflictDeviceCount *int64 `json:"conflictDeviceCount,omitempty"`

	// Number of error devices
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Number of NonCompliant devices
	NonCompliantDeviceCount *int64 `json:"nonCompliantDeviceCount,omitempty"`

	// Number of not applicable devices
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Number of not assigned devices
	NotAssignedDeviceCount *int64 `json:"notAssignedDeviceCount,omitempty"`

	// Number of remediated devices
	RemediatedDeviceCount *int64 `json:"remediatedDeviceCount,omitempty"`

	// Number of unknown devices
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`

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

func (s AdvancedThreatProtectionOnboardingStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AdvancedThreatProtectionOnboardingStateSummary{}

func (s AdvancedThreatProtectionOnboardingStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper AdvancedThreatProtectionOnboardingStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AdvancedThreatProtectionOnboardingStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AdvancedThreatProtectionOnboardingStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.advancedThreatProtectionOnboardingStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AdvancedThreatProtectionOnboardingStateSummary: %+v", err)
	}

	return encoded, nil
}
