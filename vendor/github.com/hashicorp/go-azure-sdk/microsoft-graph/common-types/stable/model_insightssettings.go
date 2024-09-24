package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InsightsSettings{}

type InsightsSettings struct {
	// The ID of a Microsoft Entra group, of which the specified type of insights are disabled for its members. The default
	// value is null. Optional.
	DisabledForGroup nullable.Type[string] `json:"disabledForGroup,omitempty"`

	// true if insights of the specified type are enabled for the organization; false if insights of the specified type are
	// disabled for all users without exceptions. The default value is true. Optional.
	IsEnabledInOrganization nullable.Type[bool] `json:"isEnabledInOrganization,omitempty"`

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

func (s InsightsSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InsightsSettings{}

func (s InsightsSettings) MarshalJSON() ([]byte, error) {
	type wrapper InsightsSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InsightsSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InsightsSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.insightsSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InsightsSettings: %+v", err)
	}

	return encoded, nil
}
