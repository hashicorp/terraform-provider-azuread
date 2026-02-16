package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StsPolicy interface {
	Entity
	DirectoryObject
	PolicyBase
	StsPolicy() BaseStsPolicyImpl
}

var _ StsPolicy = BaseStsPolicyImpl{}

type BaseStsPolicyImpl struct {
	AppliesTo *[]DirectoryObject `json:"appliesTo,omitempty"`

	// List of OData IDs for `AppliesTo` to bind to this entity
	AppliesTo_ODataBind *[]string `json:"appliesTo@odata.bind,omitempty"`

	// A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the
	// definition differs for each derived policy type. Required.
	Definition []string `json:"definition"`

	// If set to true, activates this policy. There can be many policies for the same policy type, but only one can be
	// activated as the organization default. Optional, default value is false.
	IsOrganizationDefault nullable.Type[bool] `json:"isOrganizationDefault,omitempty"`

	// Fields inherited from PolicyBase

	// Description for this policy. Required.
	Description string `json:"description"`

	// Display name for this policy. Required.
	DisplayName string `json:"displayName"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s BaseStsPolicyImpl) StsPolicy() BaseStsPolicyImpl {
	return s
}

func (s BaseStsPolicyImpl) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseStsPolicyImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseStsPolicyImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ StsPolicy = RawStsPolicyImpl{}

// RawStsPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawStsPolicyImpl struct {
	stsPolicy BaseStsPolicyImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawStsPolicyImpl) StsPolicy() BaseStsPolicyImpl {
	return s.stsPolicy
}

func (s RawStsPolicyImpl) PolicyBase() BasePolicyBaseImpl {
	return s.stsPolicy.PolicyBase()
}

func (s RawStsPolicyImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.stsPolicy.DirectoryObject()
}

func (s RawStsPolicyImpl) Entity() BaseEntityImpl {
	return s.stsPolicy.Entity()
}

var _ json.Marshaler = BaseStsPolicyImpl{}

func (s BaseStsPolicyImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseStsPolicyImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseStsPolicyImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseStsPolicyImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.stsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseStsPolicyImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseStsPolicyImpl{}

func (s *BaseStsPolicyImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppliesTo_ODataBind   *[]string             `json:"appliesTo@odata.bind,omitempty"`
		Definition            []string              `json:"definition"`
		IsOrganizationDefault nullable.Type[bool]   `json:"isOrganizationDefault,omitempty"`
		Description           string                `json:"description"`
		DisplayName           string                `json:"displayName"`
		DeletedDateTime       nullable.Type[string] `json:"deletedDateTime,omitempty"`
		Id                    *string               `json:"id,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppliesTo_ODataBind = decoded.AppliesTo_ODataBind
	s.Definition = decoded.Definition
	s.IsOrganizationDefault = decoded.IsOrganizationDefault
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseStsPolicyImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appliesTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppliesTo into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppliesTo' for 'BaseStsPolicyImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppliesTo = &output
	}

	return nil
}

func UnmarshalStsPolicyImplementation(input []byte) (StsPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StsPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.activityBasedTimeoutPolicy") {
		var out ActivityBasedTimeoutPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityBasedTimeoutPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.claimsMappingPolicy") {
		var out ClaimsMappingPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClaimsMappingPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.homeRealmDiscoveryPolicy") {
		var out HomeRealmDiscoveryPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HomeRealmDiscoveryPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tokenIssuancePolicy") {
		var out TokenIssuancePolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TokenIssuancePolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tokenLifetimePolicy") {
		var out TokenLifetimePolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TokenLifetimePolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseStsPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseStsPolicyImpl: %+v", err)
	}

	return RawStsPolicyImpl{
		stsPolicy: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
