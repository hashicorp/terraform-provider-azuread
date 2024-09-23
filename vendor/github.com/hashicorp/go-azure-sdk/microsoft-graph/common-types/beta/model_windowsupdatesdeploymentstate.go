package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentState struct {
	EffectiveValue *WindowsUpdatesDeploymentStateValue `json:"effectiveValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the reasons the deployment has its state value. Read-only.
	Reasons *[]WindowsUpdatesDeploymentStateReason `json:"reasons,omitempty"`

	RequestedValue *WindowsUpdatesRequestedDeploymentStateValue `json:"requestedValue,omitempty"`
}

var _ json.Marshaler = WindowsUpdatesDeploymentState{}

func (s WindowsUpdatesDeploymentState) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDeploymentState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDeploymentState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDeploymentState: %+v", err)
	}

	delete(decoded, "reasons")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDeploymentState: %+v", err)
	}

	return encoded, nil
}
