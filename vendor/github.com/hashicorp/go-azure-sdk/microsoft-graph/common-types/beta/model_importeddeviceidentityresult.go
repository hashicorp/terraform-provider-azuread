package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ImportedDeviceIdentity = ImportedDeviceIdentityResult{}

type ImportedDeviceIdentityResult struct {
	// Status of imported device identity
	Status *bool `json:"status,omitempty"`

	// Fields inherited from ImportedDeviceIdentity

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

	// Supported platform types for policies.
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

func (s ImportedDeviceIdentityResult) ImportedDeviceIdentity() BaseImportedDeviceIdentityImpl {
	return BaseImportedDeviceIdentityImpl{
		CreatedDateTime:            s.CreatedDateTime,
		Description:                s.Description,
		EnrollmentState:            s.EnrollmentState,
		ImportedDeviceIdentifier:   s.ImportedDeviceIdentifier,
		ImportedDeviceIdentityType: s.ImportedDeviceIdentityType,
		LastContactedDateTime:      s.LastContactedDateTime,
		LastModifiedDateTime:       s.LastModifiedDateTime,
		Platform:                   s.Platform,
		Id:                         s.Id,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
	}
}

func (s ImportedDeviceIdentityResult) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ImportedDeviceIdentityResult{}

func (s ImportedDeviceIdentityResult) MarshalJSON() ([]byte, error) {
	type wrapper ImportedDeviceIdentityResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImportedDeviceIdentityResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportedDeviceIdentityResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.importedDeviceIdentityResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImportedDeviceIdentityResult: %+v", err)
	}

	return encoded, nil
}
