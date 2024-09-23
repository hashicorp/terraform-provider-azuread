package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceMobileAppConfigurationState{}

type ManagedDeviceMobileAppConfigurationState struct {
	// The name of the policy for this policyBase
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Supported platform types for policies.
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`

	// Count of how many setting a policy holds
	SettingCount *int64 `json:"settingCount,omitempty"`

	SettingStates *[]ManagedDeviceMobileAppConfigurationSettingState `json:"settingStates,omitempty"`
	State         *ComplianceStatus                                  `json:"state,omitempty"`

	// User unique identifier, must be Guid
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// User Principal Name
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The version of the policy
	Version *int64 `json:"version,omitempty"`

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

func (s ManagedDeviceMobileAppConfigurationState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceMobileAppConfigurationState{}

func (s ManagedDeviceMobileAppConfigurationState) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceMobileAppConfigurationState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceMobileAppConfigurationState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceMobileAppConfigurationState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceMobileAppConfigurationState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceMobileAppConfigurationState: %+v", err)
	}

	return encoded, nil
}
