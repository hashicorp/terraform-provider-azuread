package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentity interface {
	Entity
	ImportedAppleDeviceIdentity() BaseImportedAppleDeviceIdentityImpl
}

var _ ImportedAppleDeviceIdentity = BaseImportedAppleDeviceIdentityImpl{}

type BaseImportedAppleDeviceIdentityImpl struct {
	// Created Date Time of the device
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device
	Description nullable.Type[string] `json:"description,omitempty"`

	DiscoverySource *DiscoverySource `json:"discoverySource,omitempty"`
	EnrollmentState *EnrollmentState `json:"enrollmentState,omitempty"`

	// Indicates if the device is deleted from Apple Business Manager
	IsDeleted nullable.Type[bool] `json:"isDeleted,omitempty"`

	// Indicates if the Apple device is supervised.
	IsSupervised *bool `json:"isSupervised,omitempty"`

	// Last Contacted Date Time of the device
	LastContactedDateTime *string `json:"lastContactedDateTime,omitempty"`

	Platform *Platform `json:"platform,omitempty"`

	// The time enrollment profile was assigned to the device
	RequestedEnrollmentProfileAssignmentDateTime nullable.Type[string] `json:"requestedEnrollmentProfileAssignmentDateTime,omitempty"`

	// Enrollment profile Id admin intends to apply to the device during next enrollment
	RequestedEnrollmentProfileId nullable.Type[string] `json:"requestedEnrollmentProfileId,omitempty"`

	// Device serial number
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

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

func (s BaseImportedAppleDeviceIdentityImpl) ImportedAppleDeviceIdentity() BaseImportedAppleDeviceIdentityImpl {
	return s
}

func (s BaseImportedAppleDeviceIdentityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ImportedAppleDeviceIdentity = RawImportedAppleDeviceIdentityImpl{}

// RawImportedAppleDeviceIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawImportedAppleDeviceIdentityImpl struct {
	importedAppleDeviceIdentity BaseImportedAppleDeviceIdentityImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawImportedAppleDeviceIdentityImpl) ImportedAppleDeviceIdentity() BaseImportedAppleDeviceIdentityImpl {
	return s.importedAppleDeviceIdentity
}

func (s RawImportedAppleDeviceIdentityImpl) Entity() BaseEntityImpl {
	return s.importedAppleDeviceIdentity.Entity()
}

var _ json.Marshaler = BaseImportedAppleDeviceIdentityImpl{}

func (s BaseImportedAppleDeviceIdentityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseImportedAppleDeviceIdentityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseImportedAppleDeviceIdentityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseImportedAppleDeviceIdentityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.importedAppleDeviceIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseImportedAppleDeviceIdentityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalImportedAppleDeviceIdentityImplementation(input []byte) (ImportedAppleDeviceIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportedAppleDeviceIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.importedAppleDeviceIdentityResult") {
		var out ImportedAppleDeviceIdentityResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedAppleDeviceIdentityResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseImportedAppleDeviceIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseImportedAppleDeviceIdentityImpl: %+v", err)
	}

	return RawImportedAppleDeviceIdentityImpl{
		importedAppleDeviceIdentity: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
