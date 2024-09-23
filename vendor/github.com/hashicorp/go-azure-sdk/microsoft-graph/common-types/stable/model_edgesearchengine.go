package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdgeSearchEngineBase = EdgeSearchEngine{}

type EdgeSearchEngine struct {
	// Allows IT admind to set a predefined default search engine for MDM-Controlled devices
	EdgeSearchEngineType *EdgeSearchEngineType `json:"edgeSearchEngineType,omitempty"`

	// Fields inherited from EdgeSearchEngineBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EdgeSearchEngine) EdgeSearchEngineBase() BaseEdgeSearchEngineBaseImpl {
	return BaseEdgeSearchEngineBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdgeSearchEngine{}

func (s EdgeSearchEngine) MarshalJSON() ([]byte, error) {
	type wrapper EdgeSearchEngine
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeSearchEngine: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeSearchEngine: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.edgeSearchEngine"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdgeSearchEngine: %+v", err)
	}

	return encoded, nil
}
