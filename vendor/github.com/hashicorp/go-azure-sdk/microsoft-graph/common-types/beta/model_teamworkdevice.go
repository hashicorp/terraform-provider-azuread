package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamworkDevice{}

type TeamworkDevice struct {
	// The activity properties that change based on the device usage.
	Activity *TeamworkDeviceActivity `json:"activity,omitempty"`

	// The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
	ActivityState *TeamworkDeviceActivityState `json:"activityState,omitempty"`

	// The company asset tag assigned by the admin on the device.
	CompanyAssetTag nullable.Type[string] `json:"companyAssetTag,omitempty"`

	// The configuration properties of the device.
	Configuration *TeamworkDeviceConfiguration `json:"configuration,omitempty"`

	// Identity of the user who enrolled the device to the tenant.
	CreatedBy IdentitySet `json:"createdBy"`

	// The UTC date and time when the device was enrolled to the tenant.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The signed-in user on the device.
	CurrentUser *TeamworkUserIdentity `json:"currentUser,omitempty"`

	DeviceType     *TeamworkDeviceType     `json:"deviceType,omitempty"`
	HardwareDetail *TeamworkHardwareDetail `json:"hardwareDetail,omitempty"`

	// The health properties of the device.
	Health *TeamworkDeviceHealth `json:"health,omitempty"`

	// The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy,
	// unknownFutureValue.
	HealthStatus *TeamworkDeviceHealthStatus `json:"healthStatus,omitempty"`

	// Identity of the user who last modified the device details.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The UTC date and time when the device detail was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The notes added by the admin to the device.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The async operations on the device.
	Operations *[]TeamworkDeviceOperation `json:"operations,omitempty"`

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

func (s TeamworkDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamworkDevice{}

func (s TeamworkDevice) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkDevice: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkDevice: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamworkDevice{}

func (s *TeamworkDevice) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Activity             *TeamworkDeviceActivity      `json:"activity,omitempty"`
		ActivityState        *TeamworkDeviceActivityState `json:"activityState,omitempty"`
		CompanyAssetTag      nullable.Type[string]        `json:"companyAssetTag,omitempty"`
		Configuration        *TeamworkDeviceConfiguration `json:"configuration,omitempty"`
		CreatedDateTime      nullable.Type[string]        `json:"createdDateTime,omitempty"`
		CurrentUser          *TeamworkUserIdentity        `json:"currentUser,omitempty"`
		DeviceType           *TeamworkDeviceType          `json:"deviceType,omitempty"`
		HardwareDetail       *TeamworkHardwareDetail      `json:"hardwareDetail,omitempty"`
		Health               *TeamworkDeviceHealth        `json:"health,omitempty"`
		HealthStatus         *TeamworkDeviceHealthStatus  `json:"healthStatus,omitempty"`
		LastModifiedDateTime nullable.Type[string]        `json:"lastModifiedDateTime,omitempty"`
		Notes                nullable.Type[string]        `json:"notes,omitempty"`
		Operations           *[]TeamworkDeviceOperation   `json:"operations,omitempty"`
		Id                   *string                      `json:"id,omitempty"`
		ODataId              *string                      `json:"@odata.id,omitempty"`
		ODataType            *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Activity = decoded.Activity
	s.ActivityState = decoded.ActivityState
	s.CompanyAssetTag = decoded.CompanyAssetTag
	s.Configuration = decoded.Configuration
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CurrentUser = decoded.CurrentUser
	s.DeviceType = decoded.DeviceType
	s.HardwareDetail = decoded.HardwareDetail
	s.Health = decoded.Health
	s.HealthStatus = decoded.HealthStatus
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Notes = decoded.Notes
	s.Operations = decoded.Operations
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamworkDevice into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamworkDevice': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'TeamworkDevice': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
