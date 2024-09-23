package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StsPolicy = TokenLifetimePolicy{}

type TokenLifetimePolicy struct {

	// Fields inherited from StsPolicy

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

func (s TokenLifetimePolicy) StsPolicy() BaseStsPolicyImpl {
	return BaseStsPolicyImpl{
		AppliesTo:             s.AppliesTo,
		AppliesTo_ODataBind:   s.AppliesTo_ODataBind,
		Definition:            s.Definition,
		IsOrganizationDefault: s.IsOrganizationDefault,
		Description:           s.Description,
		DisplayName:           s.DisplayName,
		DeletedDateTime:       s.DeletedDateTime,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s TokenLifetimePolicy) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s TokenLifetimePolicy) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s TokenLifetimePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TokenLifetimePolicy{}

func (s TokenLifetimePolicy) MarshalJSON() ([]byte, error) {
	type wrapper TokenLifetimePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TokenLifetimePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TokenLifetimePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tokenLifetimePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TokenLifetimePolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TokenLifetimePolicy{}

func (s *TokenLifetimePolicy) UnmarshalJSON(bytes []byte) error {
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
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.IsOrganizationDefault = decoded.IsOrganizationDefault
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TokenLifetimePolicy into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'AppliesTo' for 'TokenLifetimePolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppliesTo = &output
	}

	return nil
}
