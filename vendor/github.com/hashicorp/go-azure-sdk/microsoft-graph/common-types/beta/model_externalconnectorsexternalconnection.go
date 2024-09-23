package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ExternalConnectorsExternalConnection{}

type ExternalConnectorsExternalConnection struct {
	// Collects configurable settings related to activities involving connector content.
	ActivitySettings *ExternalConnectorsActivitySettings `json:"activitySettings,omitempty"`

	ComplianceSettings *ExternalConnectorsComplianceSettings `json:"complianceSettings,omitempty"`

	// Specifies additional application IDs that are allowed to manage the connection and to index content in the
	// connection. Optional.
	Configuration *ExternalConnectorsConfiguration `json:"configuration,omitempty"`

	// The Teams App ID. Optional.
	ConnectorId nullable.Type[string] `json:"connectorId,omitempty"`

	// Description of the connection displayed in the Microsoft 365 admin center. Optional.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The list of content experiences the connection will participate in. Possible values are search.
	EnabledContentExperiences *ExternalConnectorsContentExperienceType `json:"enabledContentExperiences,omitempty"`

	Groups *[]ExternalConnectorsExternalGroup `json:"groups,omitempty"`

	// The number of items ingested into a connection. This value is refreshed every 15 minutes. If the connection state is
	// draft, then ingestedItemsCount will be null.
	IngestedItemsCount nullable.Type[int64] `json:"ingestedItemsCount,omitempty"`

	Items *[]ExternalConnectorsExternalItem `json:"items,omitempty"`

	// The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128
	// characters. Required.
	Name nullable.Type[string] `json:"name,omitempty"`

	Operations *[]ExternalConnectorsConnectionOperation `json:"operations,omitempty"`
	Quota      *ExternalConnectorsConnectionQuota       `json:"quota,omitempty"`
	Schema     *ExternalConnectorsSchema                `json:"schema,omitempty"`

	// The settings configuring the search experience for content in this connection, such as the display templates for
	// search results.
	SearchSettings *ExternalConnectorsSearchSettings `json:"searchSettings,omitempty"`

	// Indicates the current state of the connection. Possible values are draft, ready, obsolete, and limitExceeded.
	// Required.
	State *ExternalConnectorsConnectionState `json:"state,omitempty"`

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

func (s ExternalConnectorsExternalConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalConnectorsExternalConnection{}

func (s ExternalConnectorsExternalConnection) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnectorsExternalConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnectorsExternalConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsExternalConnection: %+v", err)
	}

	delete(decoded, "state")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.externalConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnectorsExternalConnection: %+v", err)
	}

	return encoded, nil
}
