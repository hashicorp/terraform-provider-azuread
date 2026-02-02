package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OnPremisesAgent{}

type OnPremisesAgent struct {
	// List of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
	AgentGroups *[]OnPremisesAgentGroup `json:"agentGroups,omitempty"`

	// The external IP address as detected by the service for the agent machine. Read-only
	ExternalIp *string `json:"externalIp,omitempty"`

	// The name of the machine that the agent is running on. Read-only
	MachineName *string `json:"machineName,omitempty"`

	Status *AgentStatus `json:"status,omitempty"`

	// Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
	SupportedPublishingTypes *[]OnPremisesPublishingType `json:"supportedPublishingTypes,omitempty"`

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

func (s OnPremisesAgent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPremisesAgent{}

func (s OnPremisesAgent) MarshalJSON() ([]byte, error) {
	type wrapper OnPremisesAgent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremisesAgent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremisesAgent: %+v", err)
	}

	delete(decoded, "agentGroups")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPremisesAgent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremisesAgent: %+v", err)
	}

	return encoded, nil
}
