package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ActiveUsersMetric{}

type ActiveUsersMetric struct {
	AppId   nullable.Type[string] `json:"appId,omitempty"`
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// The total number of users who made at least one authentication request within the specified time period.
	Count *int64 `json:"count,omitempty"`

	Country nullable.Type[string] `json:"country,omitempty"`

	// Date of the insight.
	FactDate *string `json:"factDate,omitempty"`

	Language nullable.Type[string] `json:"language,omitempty"`
	Os       nullable.Type[string] `json:"os,omitempty"`

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

func (s ActiveUsersMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ActiveUsersMetric{}

func (s ActiveUsersMetric) MarshalJSON() ([]byte, error) {
	type wrapper ActiveUsersMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ActiveUsersMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ActiveUsersMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.activeUsersMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ActiveUsersMetric: %+v", err)
	}

	return encoded, nil
}
