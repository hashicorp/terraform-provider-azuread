package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ImportedAppleDeviceIdentity = ImportedAppleDeviceIdentityResult{}

type ImportedAppleDeviceIdentityResult struct {
	// Status of imported device identity
	Status *bool `json:"status,omitempty"`

	// Fields inherited from ImportedAppleDeviceIdentity

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

func (s ImportedAppleDeviceIdentityResult) ImportedAppleDeviceIdentity() BaseImportedAppleDeviceIdentityImpl {
	return BaseImportedAppleDeviceIdentityImpl{
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
		DiscoverySource:       s.DiscoverySource,
		EnrollmentState:       s.EnrollmentState,
		IsDeleted:             s.IsDeleted,
		IsSupervised:          s.IsSupervised,
		LastContactedDateTime: s.LastContactedDateTime,
		Platform:              s.Platform,
		RequestedEnrollmentProfileAssignmentDateTime: s.RequestedEnrollmentProfileAssignmentDateTime,
		RequestedEnrollmentProfileId:                 s.RequestedEnrollmentProfileId,
		SerialNumber:                                 s.SerialNumber,
		Id:                                           s.Id,
		ODataId:                                      s.ODataId,
		ODataType:                                    s.ODataType,
	}
}

func (s ImportedAppleDeviceIdentityResult) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ImportedAppleDeviceIdentityResult{}

func (s ImportedAppleDeviceIdentityResult) MarshalJSON() ([]byte, error) {
	type wrapper ImportedAppleDeviceIdentityResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImportedAppleDeviceIdentityResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportedAppleDeviceIdentityResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.importedAppleDeviceIdentityResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImportedAppleDeviceIdentityResult: %+v", err)
	}

	return encoded, nil
}
