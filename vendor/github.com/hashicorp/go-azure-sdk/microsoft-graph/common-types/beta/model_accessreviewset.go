package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewSet{}

type AccessReviewSet struct {
	// Represents a Microsoft Entra access review decision on an instance of a review.
	Decisions *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`

	// Represents the template and scheduling for an access review.
	Definitions *[]AccessReviewScheduleDefinition `json:"definitions,omitempty"`

	// Represents a collection of access review history data and the scopes used to collect that data.
	HistoryDefinitions *[]AccessReviewHistoryDefinition `json:"historyDefinitions,omitempty"`

	// Resource that enables administrators to manage directory-level access review policies in their tenant.
	Policy *AccessReviewPolicy `json:"policy,omitempty"`

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

func (s AccessReviewSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewSet{}

func (s AccessReviewSet) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewSet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewSet: %+v", err)
	}

	return encoded, nil
}
