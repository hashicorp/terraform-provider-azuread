package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentProfile interface {
	Entity
	EnrollmentProfile() BaseEnrollmentProfileImpl
}

var _ EnrollmentProfile = BaseEnrollmentProfileImpl{}

type BaseEnrollmentProfileImpl struct {
	// Configuration endpoint url to use for Enrollment
	ConfigurationEndpointUrl nullable.Type[string] `json:"configurationEndpointUrl,omitempty"`

	// Description of the profile
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the profile
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
	EnableAuthenticationViaCompanyPortal *bool `json:"enableAuthenticationViaCompanyPortal,omitempty"`

	// Indicates that Company Portal is required on setup assistant enrolled devices
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`

	// Indicates if the profile requires user authentication
	RequiresUserAuthentication *bool `json:"requiresUserAuthentication,omitempty"`

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

func (s BaseEnrollmentProfileImpl) EnrollmentProfile() BaseEnrollmentProfileImpl {
	return s
}

func (s BaseEnrollmentProfileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EnrollmentProfile = RawEnrollmentProfileImpl{}

// RawEnrollmentProfileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEnrollmentProfileImpl struct {
	enrollmentProfile BaseEnrollmentProfileImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawEnrollmentProfileImpl) EnrollmentProfile() BaseEnrollmentProfileImpl {
	return s.enrollmentProfile
}

func (s RawEnrollmentProfileImpl) Entity() BaseEntityImpl {
	return s.enrollmentProfile.Entity()
}

var _ json.Marshaler = BaseEnrollmentProfileImpl{}

func (s BaseEnrollmentProfileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEnrollmentProfileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEnrollmentProfileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEnrollmentProfileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEnrollmentProfileImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalEnrollmentProfileImplementation(input []byte) (EnrollmentProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EnrollmentProfile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.depEnrollmentBaseProfile") {
		var out DepEnrollmentBaseProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepEnrollmentBaseProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depEnrollmentProfile") {
		var out DepEnrollmentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepEnrollmentProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depTvOSEnrollmentProfile") {
		var out DepTvOSEnrollmentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepTvOSEnrollmentProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.depVisionOSEnrollmentProfile") {
		var out DepVisionOSEnrollmentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DepVisionOSEnrollmentProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseEnrollmentProfileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEnrollmentProfileImpl: %+v", err)
	}

	return RawEnrollmentProfileImpl{
		enrollmentProfile: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
