package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserAppInstallStatus{}

type UserAppInstallStatus struct {
	// The navigation link to the mobile app.
	App *MobileApp `json:"app,omitempty"`

	// The install state of the app on devices.
	DeviceStatuses *[]MobileAppInstallStatus `json:"deviceStatuses,omitempty"`

	// Failed Device Count.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// Installed Device Count.
	InstalledDeviceCount *int64 `json:"installedDeviceCount,omitempty"`

	// Not installed device count.
	NotInstalledDeviceCount *int64 `json:"notInstalledDeviceCount,omitempty"`

	// User name.
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// User Principal Name.
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

func (s UserAppInstallStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserAppInstallStatus{}

func (s UserAppInstallStatus) MarshalJSON() ([]byte, error) {
	type wrapper UserAppInstallStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserAppInstallStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserAppInstallStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userAppInstallStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserAppInstallStatus: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UserAppInstallStatus{}

func (s *UserAppInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeviceStatuses          *[]MobileAppInstallStatus `json:"deviceStatuses,omitempty"`
		FailedDeviceCount       *int64                    `json:"failedDeviceCount,omitempty"`
		InstalledDeviceCount    *int64                    `json:"installedDeviceCount,omitempty"`
		NotInstalledDeviceCount *int64                    `json:"notInstalledDeviceCount,omitempty"`
		UserName                nullable.Type[string]     `json:"userName,omitempty"`
		UserPrincipalName       nullable.Type[string]     `json:"userPrincipalName,omitempty"`
		Id                      *string                   `json:"id,omitempty"`
		ODataId                 *string                   `json:"@odata.id,omitempty"`
		ODataType               *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeviceStatuses = decoded.DeviceStatuses
	s.FailedDeviceCount = decoded.FailedDeviceCount
	s.InstalledDeviceCount = decoded.InstalledDeviceCount
	s.NotInstalledDeviceCount = decoded.NotInstalledDeviceCount
	s.UserName = decoded.UserName
	s.UserPrincipalName = decoded.UserPrincipalName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UserAppInstallStatus into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["app"]; ok {
		impl, err := UnmarshalMobileAppImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'App' for 'UserAppInstallStatus': %+v", err)
		}
		s.App = &impl
	}

	return nil
}
