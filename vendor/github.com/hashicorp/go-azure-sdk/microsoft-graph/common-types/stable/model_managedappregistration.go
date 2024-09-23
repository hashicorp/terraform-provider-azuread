package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRegistration interface {
	Entity
	ManagedAppRegistration() BaseManagedAppRegistrationImpl
}

var _ ManagedAppRegistration = BaseManagedAppRegistrationImpl{}

type BaseManagedAppRegistrationImpl struct {
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

func (s BaseManagedAppRegistrationImpl) ManagedAppRegistration() BaseManagedAppRegistrationImpl {
	return s
}

func (s BaseManagedAppRegistrationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedAppRegistration = RawManagedAppRegistrationImpl{}

// RawManagedAppRegistrationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedAppRegistrationImpl struct {
	managedAppRegistration BaseManagedAppRegistrationImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawManagedAppRegistrationImpl) ManagedAppRegistration() BaseManagedAppRegistrationImpl {
	return s.managedAppRegistration
}

func (s RawManagedAppRegistrationImpl) Entity() BaseEntityImpl {
	return s.managedAppRegistration.Entity()
}

var _ json.Marshaler = BaseManagedAppRegistrationImpl{}

func (s BaseManagedAppRegistrationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedAppRegistrationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedAppRegistrationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedAppRegistrationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedAppRegistrationImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseManagedAppRegistrationImpl{}

func (s *BaseManagedAppRegistrationImpl) UnmarshalJSON(bytes []byte) error {
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
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.ManagementSdkVersion = decoded.ManagementSdkVersion
	s.Operations = decoded.Operations
	s.PlatformVersion = decoded.PlatformVersion
	s.UserId = decoded.UserId
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseManagedAppRegistrationImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appIdentifier"]; ok {
		impl, err := UnmarshalMobileAppIdentifierImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppIdentifier' for 'BaseManagedAppRegistrationImpl': %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'AppliedPolicies' for 'BaseManagedAppRegistrationImpl': %+v", i, err)
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
				return fmt.Errorf("unmarshaling index %d field 'IntendedPolicies' for 'BaseManagedAppRegistrationImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IntendedPolicies = &output
	}

	return nil
}

func UnmarshalManagedAppRegistrationImplementation(input []byte) (ManagedAppRegistration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppRegistration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedAppRegistration") {
		var out AndroidManagedAppRegistration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedAppRegistration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosManagedAppRegistration") {
		var out IosManagedAppRegistration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosManagedAppRegistration: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedAppRegistrationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedAppRegistrationImpl: %+v", err)
	}

	return RawManagedAppRegistrationImpl{
		managedAppRegistration: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
