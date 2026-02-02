package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionCallbackConfiguration = IdentityGovernanceCustomTaskExtensionCallbackConfiguration{}

type IdentityGovernanceCustomTaskExtensionCallbackConfiguration struct {
	AuthorizedApps *[]Application `json:"authorizedApps,omitempty"`

	// Fields inherited from CustomExtensionCallbackConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The maximum duration in ISO 8601 format that Microsoft Entra ID will wait for a resume action for the callout it sent
	// to the logic app. The valid range for custom extensions in lifecycle workflows is five minutes to three hours. The
	// valid range for custom extensions in entitlement management is between 5 minutes and 14 days. For example, PT3H
	// refers to three hours, P3D refers to three days, PT10M refers to ten minutes.
	TimeoutDuration nullable.Type[string] `json:"timeoutDuration,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceCustomTaskExtensionCallbackConfiguration) CustomExtensionCallbackConfiguration() BaseCustomExtensionCallbackConfigurationImpl {
	return BaseCustomExtensionCallbackConfigurationImpl{
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
		TimeoutDuration: s.TimeoutDuration,
	}
}

var _ json.Marshaler = IdentityGovernanceCustomTaskExtensionCallbackConfiguration{}

func (s IdentityGovernanceCustomTaskExtensionCallbackConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceCustomTaskExtensionCallbackConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceCustomTaskExtensionCallbackConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceCustomTaskExtensionCallbackConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.customTaskExtensionCallbackConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceCustomTaskExtensionCallbackConfiguration: %+v", err)
	}

	return encoded, nil
}
