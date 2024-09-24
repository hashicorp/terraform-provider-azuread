package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermsAndConditionsGroupAssignment{}

type TermsAndConditionsGroupAssignment struct {
	// Unique identifier of a group that the T&C policy is assigned to.
	TargetGroupId nullable.Type[string] `json:"targetGroupId,omitempty"`

	// Navigation link to the terms and conditions that are assigned.
	TermsAndConditions *TermsAndConditions `json:"termsAndConditions,omitempty"`

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

func (s TermsAndConditionsGroupAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermsAndConditionsGroupAssignment{}

func (s TermsAndConditionsGroupAssignment) MarshalJSON() ([]byte, error) {
	type wrapper TermsAndConditionsGroupAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermsAndConditionsGroupAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermsAndConditionsGroupAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termsAndConditionsGroupAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermsAndConditionsGroupAssignment: %+v", err)
	}

	return encoded, nil
}
