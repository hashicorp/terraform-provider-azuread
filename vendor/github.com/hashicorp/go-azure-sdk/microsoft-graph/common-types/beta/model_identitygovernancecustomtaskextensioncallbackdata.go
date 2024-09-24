package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionData = IdentityGovernanceCustomTaskExtensionCallbackData{}

type IdentityGovernanceCustomTaskExtensionCallbackData struct {
	// Operation status that's provided by the Azure Logic App indicating whenever the Azure Logic App has run successfully
	// or not. Supported values: completed, failed, unknownFutureValue.
	OperationStatus *IdentityGovernanceCustomTaskExtensionOperationStatus `json:"operationStatus,omitempty"`

	// Fields inherited from CustomExtensionData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceCustomTaskExtensionCallbackData) CustomExtensionData() BaseCustomExtensionDataImpl {
	return BaseCustomExtensionDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceCustomTaskExtensionCallbackData{}

func (s IdentityGovernanceCustomTaskExtensionCallbackData) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceCustomTaskExtensionCallbackData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceCustomTaskExtensionCallbackData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceCustomTaskExtensionCallbackData: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.customTaskExtensionCallbackData"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceCustomTaskExtensionCallbackData: %+v", err)
	}

	return encoded, nil
}
