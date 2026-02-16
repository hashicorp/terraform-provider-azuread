package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessPolicyLink = NetworkaccessForwardingPolicyLink{}

type NetworkaccessForwardingPolicyLink struct {

	// Fields inherited from NetworkaccessPolicyLink

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

func (s NetworkaccessForwardingPolicyLink) NetworkaccessPolicyLink() BaseNetworkaccessPolicyLinkImpl {
	return BaseNetworkaccessPolicyLinkImpl{
		Policy:    s.Policy,
		State:     s.State,
		Version:   s.Version,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s NetworkaccessForwardingPolicyLink) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessForwardingPolicyLink{}

func (s NetworkaccessForwardingPolicyLink) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessForwardingPolicyLink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessForwardingPolicyLink: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessForwardingPolicyLink: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.forwardingPolicyLink"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessForwardingPolicyLink: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessForwardingPolicyLink{}

func (s *NetworkaccessForwardingPolicyLink) UnmarshalJSON(bytes []byte) error {
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

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.State = decoded.State
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessForwardingPolicyLink into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["policy"]; ok {
		impl, err := UnmarshalNetworkaccessPolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Policy' for 'NetworkaccessForwardingPolicyLink': %+v", err)
		}
		s.Policy = &impl
	}

	return nil
}
