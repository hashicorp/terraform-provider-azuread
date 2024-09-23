package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionData = IdentityGovernanceCustomTaskExtensionCalloutData{}

type IdentityGovernanceCustomTaskExtensionCalloutData struct {
	Subject              *User                                   `json:"subject,omitempty"`
	Task                 *IdentityGovernanceTask                 `json:"task,omitempty"`
	TaskProcessingresult *IdentityGovernanceTaskProcessingResult `json:"taskProcessingresult,omitempty"`
	Workflow             *IdentityGovernanceWorkflow             `json:"workflow,omitempty"`

	// Fields inherited from CustomExtensionData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceCustomTaskExtensionCalloutData) CustomExtensionData() BaseCustomExtensionDataImpl {
	return BaseCustomExtensionDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceCustomTaskExtensionCalloutData{}

func (s IdentityGovernanceCustomTaskExtensionCalloutData) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceCustomTaskExtensionCalloutData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceCustomTaskExtensionCalloutData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceCustomTaskExtensionCalloutData: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.customTaskExtensionCalloutData"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceCustomTaskExtensionCalloutData: %+v", err)
	}

	return encoded, nil
}
