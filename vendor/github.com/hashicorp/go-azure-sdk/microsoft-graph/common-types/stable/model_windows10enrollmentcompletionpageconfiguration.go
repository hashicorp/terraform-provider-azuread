package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceEnrollmentConfiguration = Windows10EnrollmentCompletionPageConfiguration{}

type Windows10EnrollmentCompletionPageConfiguration struct {
	// When TRUE, ESP (Enrollment Status Page) installs all required apps targeted during technician phase and ignores any
	// failures for non-blocking apps. When FALSE, ESP fails on any error during app install. The default is false.
	AllowNonBlockingAppInstallation *bool `json:"allowNonBlockingAppInstallation,omitempty"`

	// Fields inherited from DeviceEnrollmentConfiguration

	// The list of group assignments for the device configuration profile
	Assignments *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`

	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device enrollment configuration
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the device enrollment configuration
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject
	// only to the configuration with the lowest priority value.
	Priority *int64 `json:"priority,omitempty"`

	// The version of the device enrollment configuration
	Version *int64 `json:"version,omitempty"`

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

func (s Windows10EnrollmentCompletionPageConfiguration) DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl {
	return BaseDeviceEnrollmentConfigurationImpl{
		Assignments:          s.Assignments,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Priority:             s.Priority,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Windows10EnrollmentCompletionPageConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10EnrollmentCompletionPageConfiguration{}

func (s Windows10EnrollmentCompletionPageConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10EnrollmentCompletionPageConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10EnrollmentCompletionPageConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10EnrollmentCompletionPageConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10EnrollmentCompletionPageConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10EnrollmentCompletionPageConfiguration: %+v", err)
	}

	return encoded, nil
}
