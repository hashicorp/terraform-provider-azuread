package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OnPremisesAgentGroup{}

type OnPremisesAgentGroup struct {
	// List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
	Agents *[]OnPremisesAgent `json:"agents,omitempty"`

	// Display name of the onPremisesAgentGroup.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default
	// onPremisesAgentGroup and is set by the system.
	IsDefault *bool `json:"isDefault,omitempty"`

	// List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
	PublishedResources *[]PublishedResource `json:"publishedResources,omitempty"`

	PublishingType *OnPremisesPublishingType `json:"publishingType,omitempty"`

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

func (s OnPremisesAgentGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPremisesAgentGroup{}

func (s OnPremisesAgentGroup) MarshalJSON() ([]byte, error) {
	type wrapper OnPremisesAgentGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremisesAgentGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremisesAgentGroup: %+v", err)
	}

	delete(decoded, "agents")
	delete(decoded, "publishedResources")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPremisesAgentGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremisesAgentGroup: %+v", err)
	}

	return encoded, nil
}
