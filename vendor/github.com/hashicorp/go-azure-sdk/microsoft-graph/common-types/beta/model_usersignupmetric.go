package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserSignUpMetric{}

type UserSignUpMetric struct {
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The total number of users who signed up in the specified period. Supports $filter (eq).
	Count *int64 `json:"count,omitempty"`

	Country nullable.Type[string] `json:"country,omitempty"`

	// The date of the user insight.
	FactDate nullable.Type[string] `json:"factDate,omitempty"`

	IdentityProvider nullable.Type[string] `json:"identityProvider,omitempty"`
	Language         nullable.Type[string] `json:"language,omitempty"`

	// The device plaform that the customers used. Supports $filter (eq).
	Os nullable.Type[string] `json:"os,omitempty"`

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

func (s UserSignUpMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserSignUpMetric{}

func (s UserSignUpMetric) MarshalJSON() ([]byte, error) {
	type wrapper UserSignUpMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserSignUpMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSignUpMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userSignUpMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserSignUpMetric: %+v", err)
	}

	return encoded, nil
}
