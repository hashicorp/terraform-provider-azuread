package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GovernanceInsight = UserSignInInsight{}

type UserSignInInsight struct {
	// Indicates when the user last signed in.
	LastSignInDateTime nullable.Type[string] `json:"lastSignInDateTime,omitempty"`

	// Fields inherited from GovernanceInsight

	// Indicates when the insight was created.
	InsightCreatedDateTime nullable.Type[string] `json:"insightCreatedDateTime,omitempty"`

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

func (s UserSignInInsight) GovernanceInsight() BaseGovernanceInsightImpl {
	return BaseGovernanceInsightImpl{
		InsightCreatedDateTime: s.InsightCreatedDateTime,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s UserSignInInsight) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserSignInInsight{}

func (s UserSignInInsight) MarshalJSON() ([]byte, error) {
	type wrapper UserSignInInsight
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserSignInInsight: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSignInInsight: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userSignInInsight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserSignInInsight: %+v", err)
	}

	return encoded, nil
}
