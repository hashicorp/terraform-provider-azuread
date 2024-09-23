package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AdvancedThreatProtectionOnboardingDeviceSettingState{}

type AdvancedThreatProtectionOnboardingDeviceSettingState struct {
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime *string `json:"complianceGracePeriodExpirationDateTime,omitempty"`

	// The Device Id that is being reported
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The device model that is being reported
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

	// The Device Name that is being reported
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Device type.
	PlatformType *DeviceType `json:"platformType,omitempty"`

	// The setting class name and property name.
	Setting nullable.Type[string] `json:"setting,omitempty"`

	// The Setting Name that is being reported
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	State *ComplianceStatus `json:"state,omitempty"`

	// The User email address that is being reported
	UserEmail nullable.Type[string] `json:"userEmail,omitempty"`

	// The user Id that is being reported
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The User Name that is being reported
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// The User PrincipalName that is being reported
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

func (s AdvancedThreatProtectionOnboardingDeviceSettingState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AdvancedThreatProtectionOnboardingDeviceSettingState{}

func (s AdvancedThreatProtectionOnboardingDeviceSettingState) MarshalJSON() ([]byte, error) {
	type wrapper AdvancedThreatProtectionOnboardingDeviceSettingState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AdvancedThreatProtectionOnboardingDeviceSettingState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AdvancedThreatProtectionOnboardingDeviceSettingState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.advancedThreatProtectionOnboardingDeviceSettingState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AdvancedThreatProtectionOnboardingDeviceSettingState: %+v", err)
	}

	return encoded, nil
}
