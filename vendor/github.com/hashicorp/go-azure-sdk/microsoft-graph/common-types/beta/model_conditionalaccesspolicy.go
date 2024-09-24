package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPolicy interface {
	Entity
	ConditionalAccessPolicy() BaseConditionalAccessPolicyImpl
}

var _ ConditionalAccessPolicy = BaseConditionalAccessPolicyImpl{}

type BaseConditionalAccessPolicyImpl struct {
	Conditions *ConditionalAccessConditionSet `json:"conditions,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Not used.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Specifies a display name for the conditionalAccessPolicy object.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls *ConditionalAccessGrantControls `json:"grantControls"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// Specifies the session controls that are enforced after sign-in.
	SessionControls *ConditionalAccessSessionControls `json:"sessionControls"`

	State *ConditionalAccessPolicyState `json:"state,omitempty"`

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

func (s BaseConditionalAccessPolicyImpl) ConditionalAccessPolicy() BaseConditionalAccessPolicyImpl {
	return s
}

func (s BaseConditionalAccessPolicyImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ConditionalAccessPolicy = RawConditionalAccessPolicyImpl{}

// RawConditionalAccessPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConditionalAccessPolicyImpl struct {
	conditionalAccessPolicy BaseConditionalAccessPolicyImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawConditionalAccessPolicyImpl) ConditionalAccessPolicy() BaseConditionalAccessPolicyImpl {
	return s.conditionalAccessPolicy
}

func (s RawConditionalAccessPolicyImpl) Entity() BaseEntityImpl {
	return s.conditionalAccessPolicy.Entity()
}

var _ json.Marshaler = BaseConditionalAccessPolicyImpl{}

func (s BaseConditionalAccessPolicyImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseConditionalAccessPolicyImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseConditionalAccessPolicyImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseConditionalAccessPolicyImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conditionalAccessPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseConditionalAccessPolicyImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalConditionalAccessPolicyImplementation(input []byte) (ConditionalAccessPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessWhatIfPolicy") {
		var out ConditionalAccessWhatIfPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessWhatIfPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseConditionalAccessPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseConditionalAccessPolicyImpl: %+v", err)
	}

	return RawConditionalAccessPolicyImpl{
		conditionalAccessPolicy: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
