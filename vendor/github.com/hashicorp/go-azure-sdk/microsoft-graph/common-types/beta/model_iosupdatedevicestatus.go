package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IosUpdateDeviceStatus{}

type IosUpdateDeviceStatus struct {
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime *string `json:"complianceGracePeriodExpirationDateTime,omitempty"`

	// Device name of the DevicePolicyStatus.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The device id that is being reported.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The device model that is being reported
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

	InstallStatus *IosUpdatesInstallStatus `json:"installStatus,omitempty"`

	// Last modified date time of the policy report.
	LastReportedDateTime *string `json:"lastReportedDateTime,omitempty"`

	// The device version that is being reported.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Platform of the device that is being reported
	Platform *int64 `json:"platform,omitempty"`

	Status *ComplianceStatus `json:"status,omitempty"`

	// The User id that is being reported.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The User Name that is being reported
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// UserPrincipalName.
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

func (s IosUpdateDeviceStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosUpdateDeviceStatus{}

func (s IosUpdateDeviceStatus) MarshalJSON() ([]byte, error) {
	type wrapper IosUpdateDeviceStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosUpdateDeviceStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosUpdateDeviceStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosUpdateDeviceStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosUpdateDeviceStatus: %+v", err)
	}

	return encoded, nil
}
