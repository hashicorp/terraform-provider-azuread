package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdgeHomeButtonConfiguration = EdgeHomeButtonOpensNewTab{}

type EdgeHomeButtonOpensNewTab struct {

	// Fields inherited from EdgeHomeButtonConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EdgeHomeButtonOpensNewTab) EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl {
	return BaseEdgeHomeButtonConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdgeHomeButtonOpensNewTab{}

func (s EdgeHomeButtonOpensNewTab) MarshalJSON() ([]byte, error) {
	type wrapper EdgeHomeButtonOpensNewTab
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeHomeButtonOpensNewTab: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeHomeButtonOpensNewTab: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.edgeHomeButtonOpensNewTab"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdgeHomeButtonOpensNewTab: %+v", err)
	}

	return encoded, nil
}
