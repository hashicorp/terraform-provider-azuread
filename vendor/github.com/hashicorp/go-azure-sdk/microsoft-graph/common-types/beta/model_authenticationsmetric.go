package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationsMetric{}

type AuthenticationsMetric struct {
	// The ID of the Microsoft Entra application. Supports $filter (eq).
	Appid nullable.Type[string] `json:"appid,omitempty"`

	// The number of authentication requests made in the specified period. Supports $filter (eq).
	AttemptsCount *int64 `json:"attemptsCount,omitempty"`

	// The location where the customers authenticated from. Supports $filter (eq).
	Country nullable.Type[string] `json:"country,omitempty"`

	// The date of the user insight.
	FactDate nullable.Type[string] `json:"factDate,omitempty"`

	IdentityProvider nullable.Type[string] `json:"identityProvider,omitempty"`
	Language         nullable.Type[string] `json:"language,omitempty"`

	// The platform for the device that the customers used. Supports $filter (eq).
	Os nullable.Type[string] `json:"os,omitempty"`

	// Number of successful authentication requests. Supports $filter (eq).
	SuccessCount *int64 `json:"successCount,omitempty"`

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

func (s AuthenticationsMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationsMetric{}

func (s AuthenticationsMetric) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationsMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationsMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationsMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationsMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationsMetric: %+v", err)
	}

	return encoded, nil
}
