package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessProfile = NetworkaccessForwardingProfile{}

type NetworkaccessForwardingProfile struct {
	// Specifies the users, groups, devices, and remote networks whose traffic is associated with the given traffic
	// forwarding profile.
	Associations *[]NetworkaccessAssociation `json:"associations,omitempty"`

	// Profile priority.
	Priority *int64 `json:"priority,omitempty"`

	ServicePrincipal      *ServicePrincipal                   `json:"servicePrincipal,omitempty"`
	TrafficForwardingType *NetworkaccessTrafficForwardingType `json:"trafficForwardingType,omitempty"`

	// Fields inherited from NetworkaccessProfile

	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The date and time when the profile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The name of the profile.
	Name *string `json:"name,omitempty"`

	// The traffic forwarding policies associated with this profile.
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

func (s NetworkaccessForwardingProfile) NetworkaccessProfile() BaseNetworkaccessProfileImpl {
	return BaseNetworkaccessProfileImpl{
		Description:          s.Description,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		Policies:             s.Policies,
		State:                s.State,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s NetworkaccessForwardingProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessForwardingProfile{}

func (s NetworkaccessForwardingProfile) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessForwardingProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessForwardingProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessForwardingProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.forwardingProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessForwardingProfile: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessForwardingProfile{}

func (s *NetworkaccessForwardingProfile) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Priority              *int64                              `json:"priority,omitempty"`
		ServicePrincipal      *ServicePrincipal                   `json:"servicePrincipal,omitempty"`
		TrafficForwardingType *NetworkaccessTrafficForwardingType `json:"trafficForwardingType,omitempty"`
		Description           nullable.Type[string]               `json:"description,omitempty"`
		LastModifiedDateTime  *string                             `json:"lastModifiedDateTime,omitempty"`
		Name                  *string                             `json:"name,omitempty"`
		State                 *NetworkaccessStatus                `json:"state,omitempty"`
		Version               *string                             `json:"version,omitempty"`
		Id                    *string                             `json:"id,omitempty"`
		ODataId               *string                             `json:"@odata.id,omitempty"`
		ODataType             *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Priority = decoded.Priority
	s.ServicePrincipal = decoded.ServicePrincipal
	s.TrafficForwardingType = decoded.TrafficForwardingType
	s.Description = decoded.Description
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.State = decoded.State
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessForwardingProfile into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["associations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Associations into list []json.RawMessage: %+v", err)
		}

		output := make([]NetworkaccessAssociation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalNetworkaccessAssociationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Associations' for 'NetworkaccessForwardingProfile': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Associations = &output
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
				return fmt.Errorf("unmarshaling index %d field 'Policies' for 'NetworkaccessForwardingProfile': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Policies = &output
	}

	return nil
}
