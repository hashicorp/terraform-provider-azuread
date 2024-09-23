package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PublishedResource{}

type PublishedResource struct {
	// List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
	AgentGroups *[]OnPremisesAgentGroup `json:"agentGroups,omitempty"`

	// Display Name of the publishedResource.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	PublishingType *OnPremisesPublishingType `json:"publishingType,omitempty"`

	// Name of the publishedResource.
	ResourceName nullable.Type[string] `json:"resourceName,omitempty"`

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

func (s PublishedResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PublishedResource{}

func (s PublishedResource) MarshalJSON() ([]byte, error) {
	type wrapper PublishedResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PublishedResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PublishedResource: %+v", err)
	}

	delete(decoded, "agentGroups")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.publishedResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PublishedResource: %+v", err)
	}

	return encoded, nil
}
