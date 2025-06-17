package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OnPremisesPublishingProfile{}

type OnPremisesPublishingProfile struct {
	// List of existing onPremisesAgentGroup objects. Read-only. Nullable.
	AgentGroups *[]OnPremisesAgentGroup `json:"agentGroups,omitempty"`

	// List of existing onPremisesAgent objects. Read-only. Nullable.
	Agents *[]OnPremisesAgent `json:"agents,omitempty"`

	// Represents the segment configurations that are allowed for an on-premises non-web application published through
	// Microsoft Entra application proxy.
	ApplicationSegments *[]IPApplicationSegment `json:"applicationSegments,omitempty"`

	// List of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
	ConnectorGroups *[]ConnectorGroup `json:"connectorGroups,omitempty"`

	// List of existing connector objects for applications published through Application Proxy. Read-only. Nullable.
	Connectors *[]Connector `json:"connectors,omitempty"`

	// Represents a hybridAgentUpdaterConfiguration object.
	HybridAgentUpdaterConfiguration *HybridAgentUpdaterConfiguration `json:"hybridAgentUpdaterConfiguration,omitempty"`

	// Specifies whether default access for app proxy is enabled or disabled.
	IsDefaultAccessEnabled nullable.Type[bool] `json:"isDefaultAccessEnabled,omitempty"`

	// Represents if Microsoft Entra application proxy is enabled for the tenant.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// List of existing publishedResource objects. Read-only. Nullable.
	PublishedResources *[]PublishedResource `json:"publishedResources,omitempty"`

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

func (s OnPremisesPublishingProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPremisesPublishingProfile{}

func (s OnPremisesPublishingProfile) MarshalJSON() ([]byte, error) {
	type wrapper OnPremisesPublishingProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremisesPublishingProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremisesPublishingProfile: %+v", err)
	}

	delete(decoded, "agentGroups")
	delete(decoded, "agents")
	delete(decoded, "connectorGroups")
	delete(decoded, "connectors")
	delete(decoded, "publishedResources")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPremisesPublishingProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremisesPublishingProfile: %+v", err)
	}

	return encoded, nil
}
