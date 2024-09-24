package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedAppRegistration = AndroidManagedAppRegistration{}

type AndroidManagedAppRegistration struct {

	// Fields inherited from ManagedAppRegistration

	// The app package Identifier
	AppIdentifier MobileAppIdentifier `json:"appIdentifier"`

	// App version
	ApplicationVersion nullable.Type[string] `json:"applicationVersion,omitempty"`

	// Zero or more policys already applied on the registered app when it last synchronized with managment service.
	AppliedPolicies *[]ManagedAppPolicy `json:"appliedPolicies,omitempty"`

	// Date and time of creation
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Host device name
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in
	// all conditions.
	DeviceTag nullable.Type[string] `json:"deviceTag,omitempty"`

	// Host device type
	DeviceType nullable.Type[string] `json:"deviceType,omitempty"`

	// Zero or more reasons an app registration is flagged. E.g. app running on rooted device
	FlaggedReasons *[]ManagedAppFlaggedReason `json:"flaggedReasons,omitempty"`

	// Zero or more policies admin intended for the app as of now.
	IntendedPolicies *[]ManagedAppPolicy `json:"intendedPolicies,omitempty"`

	// Date and time of last the app synced with management service.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// App management SDK version
	ManagementSdkVersion nullable.Type[string] `json:"managementSdkVersion,omitempty"`

	// Zero or more long running operations triggered on the app registration.
	Operations *[]ManagedAppOperation `json:"operations,omitempty"`

	// Operating System version
	PlatformVersion nullable.Type[string] `json:"platformVersion,omitempty"`

	// The user Id to who this app registration belongs.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s AndroidManagedAppRegistration) ManagedAppRegistration() BaseManagedAppRegistrationImpl {
	return BaseManagedAppRegistrationImpl{
		AppIdentifier:        s.AppIdentifier,
		ApplicationVersion:   s.ApplicationVersion,
		AppliedPolicies:      s.AppliedPolicies,
		CreatedDateTime:      s.CreatedDateTime,
		DeviceName:           s.DeviceName,
		DeviceTag:            s.DeviceTag,
		DeviceType:           s.DeviceType,
		FlaggedReasons:       s.FlaggedReasons,
		IntendedPolicies:     s.IntendedPolicies,
		LastSyncDateTime:     s.LastSyncDateTime,
		ManagementSdkVersion: s.ManagementSdkVersion,
		Operations:           s.Operations,
		PlatformVersion:      s.PlatformVersion,
		UserId:               s.UserId,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s AndroidManagedAppRegistration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidManagedAppRegistration{}

func (s AndroidManagedAppRegistration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidManagedAppRegistration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidManagedAppRegistration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidManagedAppRegistration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidManagedAppRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidManagedAppRegistration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidManagedAppRegistration{}

func (s *AndroidManagedAppRegistration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicationVersion   nullable.Type[string]      `json:"applicationVersion,omitempty"`
		CreatedDateTime      *string                    `json:"createdDateTime,omitempty"`
		DeviceName           nullable.Type[string]      `json:"deviceName,omitempty"`
		DeviceTag            nullable.Type[string]      `json:"deviceTag,omitempty"`
		DeviceType           nullable.Type[string]      `json:"deviceType,omitempty"`
		FlaggedReasons       *[]ManagedAppFlaggedReason `json:"flaggedReasons,omitempty"`
		LastSyncDateTime     *string                    `json:"lastSyncDateTime,omitempty"`
		ManagementSdkVersion nullable.Type[string]      `json:"managementSdkVersion,omitempty"`
		Operations           *[]ManagedAppOperation     `json:"operations,omitempty"`
		PlatformVersion      nullable.Type[string]      `json:"platformVersion,omitempty"`
		UserId               nullable.Type[string]      `json:"userId,omitempty"`
		Version              nullable.Type[string]      `json:"version,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicationVersion = decoded.ApplicationVersion
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeviceName = decoded.DeviceName
	s.DeviceTag = decoded.DeviceTag
	s.DeviceType = decoded.DeviceType
	s.FlaggedReasons = decoded.FlaggedReasons
	s.Id = decoded.Id
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.ManagementSdkVersion = decoded.ManagementSdkVersion
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Operations = decoded.Operations
	s.PlatformVersion = decoded.PlatformVersion
	s.UserId = decoded.UserId
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidManagedAppRegistration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appIdentifier"]; ok {
		impl, err := UnmarshalMobileAppIdentifierImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppIdentifier' for 'AndroidManagedAppRegistration': %+v", err)
		}
		s.AppIdentifier = impl
	}

	if v, ok := temp["appliedPolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppliedPolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppliedPolicies' for 'AndroidManagedAppRegistration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppliedPolicies = &output
	}

	if v, ok := temp["intendedPolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IntendedPolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedAppPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedAppPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IntendedPolicies' for 'AndroidManagedAppRegistration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IntendedPolicies = &output
	}

	return nil
}
