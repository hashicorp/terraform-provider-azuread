package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesResourceConnection = WindowsUpdatesOperationalInsightsConnection{}

type WindowsUpdatesOperationalInsightsConnection struct {
	// The name of the Azure resource group that contains the Log Analytics workspace.
	AzureResourceGroupName nullable.Type[string] `json:"azureResourceGroupName,omitempty"`

	// The Azure subscription ID that contains the Log Analytics workspace.
	AzureSubscriptionId nullable.Type[string] `json:"azureSubscriptionId,omitempty"`

	// The name of the Log Analytics workspace.
	WorkspaceName nullable.Type[string] `json:"workspaceName,omitempty"`

	// Fields inherited from WindowsUpdatesResourceConnection

	// The state of the connection. The possible values are: connected, notAuthorized, notFound, unknownFutureValue.
	State *WindowsUpdatesResourceConnectionState `json:"state,omitempty"`

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

func (s WindowsUpdatesOperationalInsightsConnection) WindowsUpdatesResourceConnection() BaseWindowsUpdatesResourceConnectionImpl {
	return BaseWindowsUpdatesResourceConnectionImpl{
		State:     s.State,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsUpdatesOperationalInsightsConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesOperationalInsightsConnection{}

func (s WindowsUpdatesOperationalInsightsConnection) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesOperationalInsightsConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesOperationalInsightsConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesOperationalInsightsConnection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.operationalInsightsConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesOperationalInsightsConnection: %+v", err)
	}

	return encoded, nil
}
