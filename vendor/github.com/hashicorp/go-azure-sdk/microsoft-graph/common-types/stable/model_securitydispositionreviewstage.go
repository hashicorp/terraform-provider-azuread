package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityDispositionReviewStage{}

type SecurityDispositionReviewStage struct {
	// Name representing each stage within a collection.
	Name nullable.Type[string] `json:"name,omitempty"`

	// A collection of reviewers at each stage.
	ReviewersEmailAddresses *[]string `json:"reviewersEmailAddresses,omitempty"`

	// The unique sequence number for each stage of the disposition review.
	StageNumber *string `json:"stageNumber,omitempty"`

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

func (s SecurityDispositionReviewStage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityDispositionReviewStage{}

func (s SecurityDispositionReviewStage) MarshalJSON() ([]byte, error) {
	type wrapper SecurityDispositionReviewStage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityDispositionReviewStage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDispositionReviewStage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.dispositionReviewStage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityDispositionReviewStage: %+v", err)
	}

	return encoded, nil
}
