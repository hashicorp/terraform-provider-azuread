package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsApp{}

type TeamsApp struct {
	// The details for each version of the app.
	AppDefinitions *[]TeamsAppDefinition `json:"appDefinitions,omitempty"`

	// The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The method of distribution for the app. Read-only.
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`

	// The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

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

func (s TeamsApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsApp{}

func (s TeamsApp) MarshalJSON() ([]byte, error) {
	type wrapper TeamsApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsApp: %+v", err)
	}

	delete(decoded, "distributionMethod")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsApp: %+v", err)
	}

	return encoded, nil
}
