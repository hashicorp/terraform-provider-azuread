package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserRequestsMetric{}

type UserRequestsMetric struct {
	AppId   nullable.Type[string] `json:"appId,omitempty"`
	Browser nullable.Type[string] `json:"browser,omitempty"`
	Country nullable.Type[string] `json:"country,omitempty"`

	// The date of the user insight.
	FactDate nullable.Type[string] `json:"factDate,omitempty"`

	IdentityProvider nullable.Type[string] `json:"identityProvider,omitempty"`
	Language         nullable.Type[string] `json:"language,omitempty"`

	// Number of requests to the tenant. Supports $filter (eq).
	RequestCount *int64 `json:"requestCount,omitempty"`

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

func (s UserRequestsMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserRequestsMetric{}

func (s UserRequestsMetric) MarshalJSON() ([]byte, error) {
	type wrapper UserRequestsMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserRequestsMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserRequestsMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userRequestsMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserRequestsMetric: %+v", err)
	}

	return encoded, nil
}
