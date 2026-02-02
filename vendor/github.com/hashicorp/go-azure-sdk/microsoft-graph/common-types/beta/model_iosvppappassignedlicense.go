package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppAssignedLicense interface {
	Entity
	IosVppAppAssignedLicense() BaseIosVppAppAssignedLicenseImpl
}

var _ IosVppAppAssignedLicense = BaseIosVppAppAssignedLicenseImpl{}

type BaseIosVppAppAssignedLicenseImpl struct {
	// The user email address.
	UserEmailAddress nullable.Type[string] `json:"userEmailAddress,omitempty"`

	// The user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user name.
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// The user principal name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s BaseIosVppAppAssignedLicenseImpl) IosVppAppAssignedLicense() BaseIosVppAppAssignedLicenseImpl {
	return s
}

func (s BaseIosVppAppAssignedLicenseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IosVppAppAssignedLicense = RawIosVppAppAssignedLicenseImpl{}

// RawIosVppAppAssignedLicenseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIosVppAppAssignedLicenseImpl struct {
	iosVppAppAssignedLicense BaseIosVppAppAssignedLicenseImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawIosVppAppAssignedLicenseImpl) IosVppAppAssignedLicense() BaseIosVppAppAssignedLicenseImpl {
	return s.iosVppAppAssignedLicense
}

func (s RawIosVppAppAssignedLicenseImpl) Entity() BaseEntityImpl {
	return s.iosVppAppAssignedLicense.Entity()
}

var _ json.Marshaler = BaseIosVppAppAssignedLicenseImpl{}

func (s BaseIosVppAppAssignedLicenseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIosVppAppAssignedLicenseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIosVppAppAssignedLicenseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIosVppAppAssignedLicenseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosVppAppAssignedLicense"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIosVppAppAssignedLicenseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIosVppAppAssignedLicenseImplementation(input []byte) (IosVppAppAssignedLicense, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IosVppAppAssignedLicense into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppAssignedDeviceLicense") {
		var out IosVppAppAssignedDeviceLicense
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppAssignedDeviceLicense: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppAssignedUserLicense") {
		var out IosVppAppAssignedUserLicense
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppAssignedUserLicense: %+v", err)
		}
		return out, nil
	}

	var parent BaseIosVppAppAssignedLicenseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIosVppAppAssignedLicenseImpl: %+v", err)
	}

	return RawIosVppAppAssignedLicenseImpl{
		iosVppAppAssignedLicense: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
