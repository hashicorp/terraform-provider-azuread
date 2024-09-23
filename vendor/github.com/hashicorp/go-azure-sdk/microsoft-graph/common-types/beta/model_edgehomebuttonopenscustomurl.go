package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdgeHomeButtonConfiguration = EdgeHomeButtonOpensCustomURL{}

type EdgeHomeButtonOpensCustomURL struct {
	// The specific URL to load.
	HomeButtonCustomURL nullable.Type[string] `json:"homeButtonCustomURL,omitempty"`

	// Fields inherited from EdgeHomeButtonConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EdgeHomeButtonOpensCustomURL) EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl {
	return BaseEdgeHomeButtonConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdgeHomeButtonOpensCustomURL{}

func (s EdgeHomeButtonOpensCustomURL) MarshalJSON() ([]byte, error) {
	type wrapper EdgeHomeButtonOpensCustomURL
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeHomeButtonOpensCustomURL: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeHomeButtonOpensCustomURL: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.edgeHomeButtonOpensCustomURL"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdgeHomeButtonOpensCustomURL: %+v", err)
	}

	return encoded, nil
}
