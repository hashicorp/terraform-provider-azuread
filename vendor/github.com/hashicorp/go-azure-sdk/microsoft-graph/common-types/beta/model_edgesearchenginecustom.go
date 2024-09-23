package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdgeSearchEngineBase = EdgeSearchEngineCustom{}

type EdgeSearchEngineCustom struct {
	// Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to
	// the search Engine.
	EdgeSearchEngineOpenSearchXmlUrl *string `json:"edgeSearchEngineOpenSearchXmlUrl,omitempty"`

	// Fields inherited from EdgeSearchEngineBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EdgeSearchEngineCustom) EdgeSearchEngineBase() BaseEdgeSearchEngineBaseImpl {
	return BaseEdgeSearchEngineBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdgeSearchEngineCustom{}

func (s EdgeSearchEngineCustom) MarshalJSON() ([]byte, error) {
	type wrapper EdgeSearchEngineCustom
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeSearchEngineCustom: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeSearchEngineCustom: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.edgeSearchEngineCustom"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdgeSearchEngineCustom: %+v", err)
	}

	return encoded, nil
}
