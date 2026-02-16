package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionEndpointConfiguration = LogicAppTriggerEndpointConfiguration{}

type LogicAppTriggerEndpointConfiguration struct {
	// The name of the logic app.
	LogicAppWorkflowName nullable.Type[string] `json:"logicAppWorkflowName,omitempty"`

	// The Azure resource group name for the logic app.
	ResourceGroupName nullable.Type[string] `json:"resourceGroupName,omitempty"`

	// Identifier of the Azure subscription for the logic app.
	SubscriptionId nullable.Type[string] `json:"subscriptionId,omitempty"`

	// The URL to the logic app endpoint that will be triggered. Only required for app-only token scenarios where app is
	// creating a customCalloutExtension without a signed-in user.
	Url nullable.Type[string] `json:"url,omitempty"`

	// Fields inherited from CustomExtensionEndpointConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s LogicAppTriggerEndpointConfiguration) CustomExtensionEndpointConfiguration() BaseCustomExtensionEndpointConfigurationImpl {
	return BaseCustomExtensionEndpointConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LogicAppTriggerEndpointConfiguration{}

func (s LogicAppTriggerEndpointConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper LogicAppTriggerEndpointConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LogicAppTriggerEndpointConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LogicAppTriggerEndpointConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.logicAppTriggerEndpointConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LogicAppTriggerEndpointConfiguration: %+v", err)
	}

	return encoded, nil
}
