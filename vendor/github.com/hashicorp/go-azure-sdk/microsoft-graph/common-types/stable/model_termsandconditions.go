package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermsAndConditions{}

type TermsAndConditions struct {
	// Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the
	// terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
	AcceptanceStatement nullable.Type[string] `json:"acceptanceStatement,omitempty"`

	// The list of acceptance statuses for this T&C policy.
	AcceptanceStatuses *[]TermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`

	// The list of assignments for this T&C policy.
	Assignments *[]TermsAndConditionsAssignment `json:"assignments,omitempty"`

	// Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the
	// user on prompts to accept the T&C policy.
	BodyText nullable.Type[string] `json:"bodyText,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Administrator-supplied description of the T&C policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Administrator-supplied name for the T&C policy.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C
	// policy.
	Title nullable.Type[string] `json:"title,omitempty"`

	// Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms
	// and wishes to require users to re-accept the modified T&C policy.
	Version *int64 `json:"version,omitempty"`

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

func (s TermsAndConditions) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermsAndConditions{}

func (s TermsAndConditions) MarshalJSON() ([]byte, error) {
	type wrapper TermsAndConditions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermsAndConditions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermsAndConditions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termsAndConditions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermsAndConditions: %+v", err)
	}

	return encoded, nil
}
