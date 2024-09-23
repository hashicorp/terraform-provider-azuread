package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceAccessProfileBase interface {
	Entity
	DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl
}

var _ DeviceManagementResourceAccessProfileBase = BaseDeviceManagementResourceAccessProfileBaseImpl{}

type BaseDeviceManagementResourceAccessProfileBaseImpl struct {
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

func (s BaseDeviceManagementResourceAccessProfileBaseImpl) DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl {
	return s
}

func (s BaseDeviceManagementResourceAccessProfileBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementResourceAccessProfileBase = RawDeviceManagementResourceAccessProfileBaseImpl{}

// RawDeviceManagementResourceAccessProfileBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementResourceAccessProfileBaseImpl struct {
	deviceManagementResourceAccessProfileBase BaseDeviceManagementResourceAccessProfileBaseImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawDeviceManagementResourceAccessProfileBaseImpl) DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl {
	return s.deviceManagementResourceAccessProfileBase
}

func (s RawDeviceManagementResourceAccessProfileBaseImpl) Entity() BaseEntityImpl {
	return s.deviceManagementResourceAccessProfileBase.Entity()
}

var _ json.Marshaler = BaseDeviceManagementResourceAccessProfileBaseImpl{}

func (s BaseDeviceManagementResourceAccessProfileBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementResourceAccessProfileBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementResourceAccessProfileBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementResourceAccessProfileBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementResourceAccessProfileBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementResourceAccessProfileBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceManagementResourceAccessProfileBaseImplementation(input []byte) (DeviceManagementResourceAccessProfileBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementResourceAccessProfileBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XCertificateProfile") {
		var out Windows10XCertificateProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XCertificateProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XTrustedRootCertificate") {
		var out Windows10XTrustedRootCertificate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XTrustedRootCertificate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XVpnConfiguration") {
		var out Windows10XVpnConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XVpnConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10XWifiConfiguration") {
		var out Windows10XWifiConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10XWifiConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementResourceAccessProfileBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementResourceAccessProfileBaseImpl: %+v", err)
	}

	return RawDeviceManagementResourceAccessProfileBaseImpl{
		deviceManagementResourceAccessProfileBase: parent,
		Type:   value,
		Values: temp,
	}, nil

}
