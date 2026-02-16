package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityBaselineDeviceState{}

type SecurityBaselineDeviceState struct {
	// Display name of the device
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// Last modified date time of the policy report
	LastReportedDateTime *string `json:"lastReportedDateTime,omitempty"`

	// Intune device id
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// Security Baseline Compliance State
	State *SecurityBaselineComplianceState `json:"state,omitempty"`

	// User Principal Name
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s SecurityBaselineDeviceState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityBaselineDeviceState{}

func (s SecurityBaselineDeviceState) MarshalJSON() ([]byte, error) {
	type wrapper SecurityBaselineDeviceState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityBaselineDeviceState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityBaselineDeviceState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityBaselineDeviceState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityBaselineDeviceState: %+v", err)
	}

	return encoded, nil
}
