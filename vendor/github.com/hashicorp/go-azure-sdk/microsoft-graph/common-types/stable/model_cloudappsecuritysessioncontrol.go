package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessSessionControl = CloudAppSecuritySessionControl{}

type CloudAppSecuritySessionControl struct {
	// Possible values are: mcasConfigured, monitorOnly, blockDownloads, unknownFutureValue. For more information, see
	// Deploy Conditional Access App Control for featured apps.
	CloudAppSecurityType *CloudAppSecuritySessionControlType `json:"cloudAppSecurityType,omitempty"`

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

func (s CloudAppSecuritySessionControl) ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl {
	return BaseConditionalAccessSessionControlImpl{
		IsEnabled: s.IsEnabled,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudAppSecuritySessionControl{}

func (s CloudAppSecuritySessionControl) MarshalJSON() ([]byte, error) {
	type wrapper CloudAppSecuritySessionControl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudAppSecuritySessionControl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudAppSecuritySessionControl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudAppSecuritySessionControl"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudAppSecuritySessionControl: %+v", err)
	}

	return encoded, nil
}
