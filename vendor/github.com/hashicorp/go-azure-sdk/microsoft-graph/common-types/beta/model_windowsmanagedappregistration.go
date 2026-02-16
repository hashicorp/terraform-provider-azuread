package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedAppRegistration = WindowsManagedAppRegistration{}

type WindowsManagedAppRegistration struct {

	// Fields inherited from ManagedAppRegistration

	// The app package Identifier
	AppIdentifier MobileAppIdentifier `json:"appIdentifier"`

	// App version
	ApplicationVersion nullable.Type[string] `json:"applicationVersion,omitempty"`

	// Zero or more policys already applied on the registered app when it last synchronized with managment service.
	AppliedPolicies *[]ManagedAppPolicy `json:"appliedPolicies,omitempty"`

	// The Azure Active Directory Device identifier of the host device. Value could be empty even when the host device is
	// Azure Active Directory registered.
	AzureADDeviceId nullable.Type[string] `json:"azureADDeviceId,omitempty"`

	// Date and time of creation
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The device manufacturer for the current app registration
	DeviceManufacturer nullable.Type[string] `json:"deviceManufacturer,omitempty"`

	// The device model for the current app registration
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

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

	// Zero or more log collection requests triggered for the app.
	ManagedAppLogCollectionRequests *[]ManagedAppLogCollectionRequest `json:"managedAppLogCollectionRequests,omitempty"`

	// The Managed Device identifier of the host device. Value could be empty even when the host device is managed.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

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

func (s WindowsManagedAppRegistration) ManagedAppRegistration() BaseManagedAppRegistrationImpl {
	return BaseManagedAppRegistrationImpl{
		AppIdentifier:                   s.AppIdentifier,
		ApplicationVersion:              s.ApplicationVersion,
		AppliedPolicies:                 s.AppliedPolicies,
		AzureADDeviceId:                 s.AzureADDeviceId,
		CreatedDateTime:                 s.CreatedDateTime,
		DeviceManufacturer:              s.DeviceManufacturer,
		DeviceModel:                     s.DeviceModel,
		DeviceName:                      s.DeviceName,
		DeviceTag:                       s.DeviceTag,
		DeviceType:                      s.DeviceType,
		FlaggedReasons:                  s.FlaggedReasons,
		IntendedPolicies:                s.IntendedPolicies,
		LastSyncDateTime:                s.LastSyncDateTime,
		ManagedAppLogCollectionRequests: s.ManagedAppLogCollectionRequests,
		ManagedDeviceId:                 s.ManagedDeviceId,
		ManagementSdkVersion:            s.ManagementSdkVersion,
		Operations:                      s.Operations,
		PlatformVersion:                 s.PlatformVersion,
		UserId:                          s.UserId,
		Version:                         s.Version,
		Id:                              s.Id,
		ODataId:                         s.ODataId,
		ODataType:                       s.ODataType,
	}
}

func (s WindowsManagedAppRegistration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsManagedAppRegistration{}

func (s WindowsManagedAppRegistration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsManagedAppRegistration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsManagedAppRegistration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsManagedAppRegistration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsManagedAppRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsManagedAppRegistration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsManagedAppRegistration{}

func (s *WindowsManagedAppRegistration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicationVersion              nullable.Type[string]             `json:"applicationVersion,omitempty"`
		AzureADDeviceId                 nullable.Type[string]             `json:"azureADDeviceId,omitempty"`
		CreatedDateTime                 *string                           `json:"createdDateTime,omitempty"`
		DeviceManufacturer              nullable.Type[string]             `json:"deviceManufacturer,omitempty"`
		DeviceModel                     nullable.Type[string]             `json:"deviceModel,omitempty"`
		DeviceName                      nullable.Type[string]             `json:"deviceName,omitempty"`
		DeviceTag                       nullable.Type[string]             `json:"deviceTag,omitempty"`
		DeviceType                      nullable.Type[string]             `json:"deviceType,omitempty"`
		FlaggedReasons                  *[]ManagedAppFlaggedReason        `json:"flaggedReasons,omitempty"`
		LastSyncDateTime                *string                           `json:"lastSyncDateTime,omitempty"`
		ManagedAppLogCollectionRequests *[]ManagedAppLogCollectionRequest `json:"managedAppLogCollectionRequests,omitempty"`
		ManagedDeviceId                 nullable.Type[string]             `json:"managedDeviceId,omitempty"`
		ManagementSdkVersion            nullable.Type[string]             `json:"managementSdkVersion,omitempty"`
		Operations                      *[]ManagedAppOperation            `json:"operations,omitempty"`
		PlatformVersion                 nullable.Type[string]             `json:"platformVersion,omitempty"`
		UserId                          nullable.Type[string]             `json:"userId,omitempty"`
		Version                         nullable.Type[string]             `json:"version,omitempty"`
		Id                              *string                           `json:"id,omitempty"`
		ODataId                         *string                           `json:"@odata.id,omitempty"`
		ODataType                       *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicationVersion = decoded.ApplicationVersion
	s.AzureADDeviceId = decoded.AzureADDeviceId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeviceManufacturer = decoded.DeviceManufacturer
	s.DeviceModel = decoded.DeviceModel
	s.DeviceName = decoded.DeviceName
	s.DeviceTag = decoded.DeviceTag
	s.DeviceType = decoded.DeviceType
	s.FlaggedReasons = decoded.FlaggedReasons
	s.Id = decoded.Id
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.ManagedAppLogCollectionRequests = decoded.ManagedAppLogCollectionRequests
	s.ManagedDeviceId = decoded.ManagedDeviceId
	s.ManagementSdkVersion = decoded.ManagementSdkVersion
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Operations = decoded.Operations
	s.PlatformVersion = decoded.PlatformVersion
	s.UserId = decoded.UserId
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsManagedAppRegistration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appIdentifier"]; ok {
		impl, err := UnmarshalMobileAppIdentifierImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppIdentifier' for 'WindowsManagedAppRegistration': %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'AppliedPolicies' for 'WindowsManagedAppRegistration': %+v", i, err)
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
				return fmt.Errorf("unmarshaling index %d field 'IntendedPolicies' for 'WindowsManagedAppRegistration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IntendedPolicies = &output
	}

	return nil
}
