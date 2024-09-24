package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessPolicy = ConditionalAccessWhatIfPolicy{}

type ConditionalAccessWhatIfPolicy struct {
	PolicyApplies *bool                             `json:"policyApplies,omitempty"`
	Reasons       *[]ConditionalAccessWhatIfReasons `json:"reasons,omitempty"`

	// Fields inherited from ConditionalAccessPolicy

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

func (s ConditionalAccessWhatIfPolicy) ConditionalAccessPolicy() BaseConditionalAccessPolicyImpl {
	return BaseConditionalAccessPolicyImpl{
		Conditions:       s.Conditions,
		CreatedDateTime:  s.CreatedDateTime,
		Description:      s.Description,
		DisplayName:      s.DisplayName,
		GrantControls:    s.GrantControls,
		ModifiedDateTime: s.ModifiedDateTime,
		SessionControls:  s.SessionControls,
		State:            s.State,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s ConditionalAccessWhatIfPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConditionalAccessWhatIfPolicy{}

func (s ConditionalAccessWhatIfPolicy) MarshalJSON() ([]byte, error) {
	type wrapper ConditionalAccessWhatIfPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConditionalAccessWhatIfPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessWhatIfPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conditionalAccessWhatIfPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConditionalAccessWhatIfPolicy: %+v", err)
	}

	return encoded, nil
}
