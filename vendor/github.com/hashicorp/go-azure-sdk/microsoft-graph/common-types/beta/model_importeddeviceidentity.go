package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentity interface {
	Entity
	ImportedDeviceIdentity() BaseImportedDeviceIdentityImpl
}

var _ ImportedDeviceIdentity = BaseImportedDeviceIdentityImpl{}

type BaseImportedDeviceIdentityImpl struct {
	// Created Date Time of the device
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device
	Description nullable.Type[string] `json:"description,omitempty"`

	EnrollmentState *EnrollmentState `json:"enrollmentState,omitempty"`

	// Imported Device Identifier
	ImportedDeviceIdentifier nullable.Type[string] `json:"importedDeviceIdentifier,omitempty"`

	ImportedDeviceIdentityType *ImportedDeviceIdentityType `json:"importedDeviceIdentityType,omitempty"`

	// Last Contacted Date Time of the device
	LastContactedDateTime *string `json:"lastContactedDateTime,omitempty"`

	// Last Modified DateTime of the description
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	Platform *Platform `json:"platform,omitempty"`

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

func (s BaseImportedDeviceIdentityImpl) ImportedDeviceIdentity() BaseImportedDeviceIdentityImpl {
	return s
}

func (s BaseImportedDeviceIdentityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ImportedDeviceIdentity = RawImportedDeviceIdentityImpl{}

// RawImportedDeviceIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawImportedDeviceIdentityImpl struct {
	importedDeviceIdentity BaseImportedDeviceIdentityImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawImportedDeviceIdentityImpl) ImportedDeviceIdentity() BaseImportedDeviceIdentityImpl {
	return s.importedDeviceIdentity
}

func (s RawImportedDeviceIdentityImpl) Entity() BaseEntityImpl {
	return s.importedDeviceIdentity.Entity()
}

var _ json.Marshaler = BaseImportedDeviceIdentityImpl{}

func (s BaseImportedDeviceIdentityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseImportedDeviceIdentityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseImportedDeviceIdentityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseImportedDeviceIdentityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.importedDeviceIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseImportedDeviceIdentityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalImportedDeviceIdentityImplementation(input []byte) (ImportedDeviceIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportedDeviceIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.importedDeviceIdentityResult") {
		var out ImportedDeviceIdentityResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportedDeviceIdentityResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseImportedDeviceIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseImportedDeviceIdentityImpl: %+v", err)
	}

	return RawImportedDeviceIdentityImpl{
		importedDeviceIdentity: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
