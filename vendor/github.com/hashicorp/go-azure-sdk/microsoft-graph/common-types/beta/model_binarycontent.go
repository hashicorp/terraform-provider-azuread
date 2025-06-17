package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ContentBase = BinaryContent{}

type BinaryContent struct {
	// The binary content, encoded as a Base64 string. Inherited from contentBase.
	Data *string `json:"data,omitempty"`

	// Fields inherited from ContentBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BinaryContent) ContentBase() BaseContentBaseImpl {
	return BaseContentBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BinaryContent{}

func (s BinaryContent) MarshalJSON() ([]byte, error) {
	type wrapper BinaryContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BinaryContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BinaryContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.binaryContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BinaryContent: %+v", err)
	}

	return encoded, nil
}
