package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppInstallStatus{}

type MobileAppInstallStatus struct {
	// The navigation link to the mobile app.
	App *MobileApp `json:"app,omitempty"`

	// Device ID
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Device name
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Human readable version of the application
	DisplayVersion nullable.Type[string] `json:"displayVersion,omitempty"`

	// The error code for install or uninstall failures.
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// A list of possible states for application status on an individual device. When devices contact the Intune service and
	// find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the
	// Graph API. Since the application status is identified during device interaction with the Intune service, status
	// records do not immediately appear upon application group assignment; it is created only after the assignment is
	// evaluated in the service and devices start receiving the policy during check-ins.
	InstallState *ResultantAppState `json:"installState,omitempty"`

	// Enum indicating additional details regarding why an application has a particular install state.
	InstallStateDetail *ResultantAppStateDetail `json:"installStateDetail,omitempty"`

	// Last sync date time
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// A list of possible states for application status on an individual device. When devices contact the Intune service and
	// find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the
	// Graph API. Since the application status is identified during device interaction with the Intune service, status
	// records do not immediately appear upon application group assignment; it is created only after the assignment is
	// evaluated in the service and devices start receiving the policy during check-ins.
	MobileAppInstallStatusValue *ResultantAppState `json:"mobileAppInstallStatusValue,omitempty"`

	// OS Description
	OsDescription nullable.Type[string] `json:"osDescription,omitempty"`

	// OS Version
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Device User Name
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s MobileAppInstallStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppInstallStatus{}

func (s MobileAppInstallStatus) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppInstallStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppInstallStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppInstallStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppInstallStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppInstallStatus: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MobileAppInstallStatus{}

func (s *MobileAppInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeviceId                    nullable.Type[string]    `json:"deviceId,omitempty"`
		DeviceName                  nullable.Type[string]    `json:"deviceName,omitempty"`
		DisplayVersion              nullable.Type[string]    `json:"displayVersion,omitempty"`
		ErrorCode                   *int64                   `json:"errorCode,omitempty"`
		InstallState                *ResultantAppState       `json:"installState,omitempty"`
		InstallStateDetail          *ResultantAppStateDetail `json:"installStateDetail,omitempty"`
		LastSyncDateTime            *string                  `json:"lastSyncDateTime,omitempty"`
		MobileAppInstallStatusValue *ResultantAppState       `json:"mobileAppInstallStatusValue,omitempty"`
		OsDescription               nullable.Type[string]    `json:"osDescription,omitempty"`
		OsVersion                   nullable.Type[string]    `json:"osVersion,omitempty"`
		UserName                    nullable.Type[string]    `json:"userName,omitempty"`
		UserPrincipalName           nullable.Type[string]    `json:"userPrincipalName,omitempty"`
		Id                          *string                  `json:"id,omitempty"`
		ODataId                     *string                  `json:"@odata.id,omitempty"`
		ODataType                   *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeviceId = decoded.DeviceId
	s.DeviceName = decoded.DeviceName
	s.DisplayVersion = decoded.DisplayVersion
	s.ErrorCode = decoded.ErrorCode
	s.InstallState = decoded.InstallState
	s.InstallStateDetail = decoded.InstallStateDetail
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.MobileAppInstallStatusValue = decoded.MobileAppInstallStatusValue
	s.OsDescription = decoded.OsDescription
	s.OsVersion = decoded.OsVersion
	s.UserName = decoded.UserName
	s.UserPrincipalName = decoded.UserPrincipalName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MobileAppInstallStatus into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["app"]; ok {
		impl, err := UnmarshalMobileAppImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'App' for 'MobileAppInstallStatus': %+v", err)
		}
		s.App = &impl
	}

	return nil
}
