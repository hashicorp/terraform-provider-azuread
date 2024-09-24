package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XCertificateProfile interface {
	Entity
	DeviceManagementResourceAccessProfileBase
	Windows10XCertificateProfile() BaseWindows10XCertificateProfileImpl
}

var _ Windows10XCertificateProfile = BaseWindows10XCertificateProfileImpl{}

type BaseWindows10XCertificateProfileImpl struct {

	// Fields inherited from DeviceManagementResourceAccessProfileBase

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceManagementResourceAccessProfileAssignment `json:"assignments,omitempty"`

	// DateTime profile was created
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`

	// Profile description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Profile display name
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime profile was last modified
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Scope Tags
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Version of the profile
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

func (s BaseWindows10XCertificateProfileImpl) Windows10XCertificateProfile() BaseWindows10XCertificateProfileImpl {
	return s
}

func (s BaseWindows10XCertificateProfileImpl) DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl {
	return BaseDeviceManagementResourceAccessProfileBaseImpl{
		Assignments:          s.Assignments,
		CreationDateTime:     s.CreationDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseWindows10XCertificateProfileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Windows10XCertificateProfile = RawWindows10XCertificateProfileImpl{}

// RawWindows10XCertificateProfileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindows10XCertificateProfileImpl struct {
	windows10XCertificateProfile BaseWindows10XCertificateProfileImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawWindows10XCertificateProfileImpl) Windows10XCertificateProfile() BaseWindows10XCertificateProfileImpl {
	return s.windows10XCertificateProfile
}

func (s RawWindows10XCertificateProfileImpl) DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl {
	return s.windows10XCertificateProfile.DeviceManagementResourceAccessProfileBase()
}

func (s RawWindows10XCertificateProfileImpl) Entity() BaseEntityImpl {
	return s.windows10XCertificateProfile.Entity()
}

var _ json.Marshaler = BaseWindows10XCertificateProfileImpl{}

func (s BaseWindows10XCertificateProfileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindows10XCertificateProfileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindows10XCertificateProfileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindows10XCertificateProfileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10XCertificateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindows10XCertificateProfileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindows10XCertificateProfileImplementation(input []byte) (Windows10XCertificateProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10XCertificateProfile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XSCEPCertificateProfile") {
		var out Windows10XSCEPCertificateProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XSCEPCertificateProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindows10XCertificateProfileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindows10XCertificateProfileImpl: %+v", err)
	}

	return RawWindows10XCertificateProfileImpl{
		windows10XCertificateProfile: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
