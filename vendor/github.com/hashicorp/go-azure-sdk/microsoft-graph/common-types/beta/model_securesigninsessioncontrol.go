package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessSessionControl = SecureSignInSessionControl{}

type SecureSignInSessionControl struct {

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

func (s SecureSignInSessionControl) ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl {
	return BaseConditionalAccessSessionControlImpl{
		IsEnabled: s.IsEnabled,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecureSignInSessionControl{}

func (s SecureSignInSessionControl) MarshalJSON() ([]byte, error) {
	type wrapper SecureSignInSessionControl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecureSignInSessionControl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecureSignInSessionControl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.secureSignInSessionControl"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecureSignInSessionControl: %+v", err)
	}

	return encoded, nil
}
