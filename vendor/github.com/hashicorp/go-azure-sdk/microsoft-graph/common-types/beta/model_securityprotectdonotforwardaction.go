package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityInformationProtectionAction = SecurityProtectDoNotForwardAction{}

type SecurityProtectDoNotForwardAction struct {

	// Fields inherited from SecurityInformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityProtectDoNotForwardAction) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return BaseSecurityInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityProtectDoNotForwardAction{}

func (s SecurityProtectDoNotForwardAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityProtectDoNotForwardAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityProtectDoNotForwardAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityProtectDoNotForwardAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.protectDoNotForwardAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityProtectDoNotForwardAction: %+v", err)
	}

	return encoded, nil
}
