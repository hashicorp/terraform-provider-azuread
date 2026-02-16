package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityResponseAction = SecurityDisableUserResponseAction{}

type SecurityDisableUserResponseAction struct {
	Identifier *SecurityDisableUserEntityIdentifier `json:"identifier,omitempty"`

	// Fields inherited from SecurityResponseAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityDisableUserResponseAction) SecurityResponseAction() BaseSecurityResponseActionImpl {
	return BaseSecurityResponseActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityDisableUserResponseAction{}

func (s SecurityDisableUserResponseAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityDisableUserResponseAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityDisableUserResponseAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDisableUserResponseAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.disableUserResponseAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityDisableUserResponseAction: %+v", err)
	}

	return encoded, nil
}
