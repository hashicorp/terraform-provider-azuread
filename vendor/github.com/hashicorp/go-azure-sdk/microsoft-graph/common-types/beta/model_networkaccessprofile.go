package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessProfile interface {
	Entity
	NetworkaccessProfile() BaseNetworkaccessProfileImpl
}

var _ NetworkaccessProfile = BaseNetworkaccessProfileImpl{}

type BaseNetworkaccessProfileImpl struct {
	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Profile last modified time.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Profile name.
	Name *string `json:"name,omitempty"`

	// Traffic forwarding policies associated with this profile.
	Policies *[]NetworkaccessPolicyLink `json:"policies,omitempty"`

	State *NetworkaccessStatus `json:"state,omitempty"`

	// Profile version.
	Version *string `json:"version,omitempty"`

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

func (s BaseNetworkaccessProfileImpl) NetworkaccessProfile() BaseNetworkaccessProfileImpl {
	return s
}

func (s BaseNetworkaccessProfileImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NetworkaccessProfile = RawNetworkaccessProfileImpl{}

// RawNetworkaccessProfileImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessProfileImpl struct {
	networkaccessProfile BaseNetworkaccessProfileImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawNetworkaccessProfileImpl) NetworkaccessProfile() BaseNetworkaccessProfileImpl {
	return s.networkaccessProfile
}

func (s RawNetworkaccessProfileImpl) Entity() BaseEntityImpl {
	return s.networkaccessProfile.Entity()
}

var _ json.Marshaler = BaseNetworkaccessProfileImpl{}

func (s BaseNetworkaccessProfileImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNetworkaccessProfileImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNetworkaccessProfileImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNetworkaccessProfileImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.profile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNetworkaccessProfileImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseNetworkaccessProfileImpl{}

func (s *BaseNetworkaccessProfileImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description          nullable.Type[string] `json:"description,omitempty"`
		LastModifiedDateTime *string               `json:"lastModifiedDateTime,omitempty"`
		Name                 *string               `json:"name,omitempty"`
		State                *NetworkaccessStatus  `json:"state,omitempty"`
		Version              *string               `json:"version,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.State = decoded.State
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseNetworkaccessProfileImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["policies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Policies into list []json.RawMessage: %+v", err)
		}

		output := make([]NetworkaccessPolicyLink, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalNetworkaccessPolicyLinkImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Policies' for 'BaseNetworkaccessProfileImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Policies = &output
	}

	return nil
}

func UnmarshalNetworkaccessProfileImplementation(input []byte) (NetworkaccessProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessProfile into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringProfile") {
		var out NetworkaccessFilteringProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingProfile") {
		var out NetworkaccessForwardingProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessProfileImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessProfileImpl: %+v", err)
	}

	return RawNetworkaccessProfileImpl{
		networkaccessProfile: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
