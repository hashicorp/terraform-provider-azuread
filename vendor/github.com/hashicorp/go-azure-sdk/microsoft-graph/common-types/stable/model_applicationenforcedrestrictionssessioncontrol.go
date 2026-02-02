package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessSessionControl = ApplicationEnforcedRestrictionsSessionControl{}

type ApplicationEnforcedRestrictionsSessionControl struct {

	// Fields inherited from ConditionalAccessSessionControl

	// Specifies whether the session control is enabled.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ApplicationEnforcedRestrictionsSessionControl) ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl {
	return BaseConditionalAccessSessionControlImpl{
		IsEnabled: s.IsEnabled,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApplicationEnforcedRestrictionsSessionControl{}

func (s ApplicationEnforcedRestrictionsSessionControl) MarshalJSON() ([]byte, error) {
	type wrapper ApplicationEnforcedRestrictionsSessionControl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApplicationEnforcedRestrictionsSessionControl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationEnforcedRestrictionsSessionControl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.applicationEnforcedRestrictionsSessionControl"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApplicationEnforcedRestrictionsSessionControl: %+v", err)
	}

	return encoded, nil
}
