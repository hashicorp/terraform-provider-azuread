package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyLink interface {
	Entity
	NetworkaccessPolicyLink() BaseNetworkaccessPolicyLinkImpl
}

var _ NetworkaccessPolicyLink = BaseNetworkaccessPolicyLinkImpl{}

type BaseNetworkaccessPolicyLinkImpl struct {
	Policy *NetworkaccessPolicy `json:"policy,omitempty"`
	State  *NetworkaccessStatus `json:"state,omitempty"`

	// Version.
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

func (s BaseNetworkaccessPolicyLinkImpl) NetworkaccessPolicyLink() BaseNetworkaccessPolicyLinkImpl {
	return s
}

func (s BaseNetworkaccessPolicyLinkImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NetworkaccessPolicyLink = RawNetworkaccessPolicyLinkImpl{}

// RawNetworkaccessPolicyLinkImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessPolicyLinkImpl struct {
	networkaccessPolicyLink BaseNetworkaccessPolicyLinkImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawNetworkaccessPolicyLinkImpl) NetworkaccessPolicyLink() BaseNetworkaccessPolicyLinkImpl {
	return s.networkaccessPolicyLink
}

func (s RawNetworkaccessPolicyLinkImpl) Entity() BaseEntityImpl {
	return s.networkaccessPolicyLink.Entity()
}

var _ json.Marshaler = BaseNetworkaccessPolicyLinkImpl{}

func (s BaseNetworkaccessPolicyLinkImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNetworkaccessPolicyLinkImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNetworkaccessPolicyLinkImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNetworkaccessPolicyLinkImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.policyLink"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNetworkaccessPolicyLinkImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseNetworkaccessPolicyLinkImpl{}

func (s *BaseNetworkaccessPolicyLinkImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		State     *NetworkaccessStatus `json:"state,omitempty"`
		Version   *string              `json:"version,omitempty"`
		Id        *string              `json:"id,omitempty"`
		ODataId   *string              `json:"@odata.id,omitempty"`
		ODataType *string              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.State = decoded.State
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseNetworkaccessPolicyLinkImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["policy"]; ok {
		impl, err := UnmarshalNetworkaccessPolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Policy' for 'BaseNetworkaccessPolicyLinkImpl': %+v", err)
		}
		s.Policy = &impl
	}

	return nil
}

func UnmarshalNetworkaccessPolicyLinkImplementation(input []byte) (NetworkaccessPolicyLink, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessPolicyLink into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringPolicyLink") {
		var out NetworkaccessFilteringPolicyLink
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringPolicyLink: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingPolicyLink") {
		var out NetworkaccessForwardingPolicyLink
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingPolicyLink: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessPolicyLinkImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessPolicyLinkImpl: %+v", err)
	}

	return RawNetworkaccessPolicyLinkImpl{
		networkaccessPolicyLink: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
