package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimTransformation = ExtractAlphaTransformation{}

type ExtractAlphaTransformation struct {
	Type *TransformationExtractType `json:"type,omitempty"`

	// Fields inherited from CustomClaimTransformation

	// The input attribute that provides the source for the transformation. This parameter is required if it's the first or
	// only transformation in the list of transformations to be applied. Subsequent transformations use the output of the
	// prior transformation as input.
	Input *TransformationAttribute `json:"input,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ExtractAlphaTransformation) CustomClaimTransformation() BaseCustomClaimTransformationImpl {
	return BaseCustomClaimTransformationImpl{
		Input:     s.Input,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExtractAlphaTransformation{}

func (s ExtractAlphaTransformation) MarshalJSON() ([]byte, error) {
	type wrapper ExtractAlphaTransformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExtractAlphaTransformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExtractAlphaTransformation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.extractAlphaTransformation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExtractAlphaTransformation: %+v", err)
	}

	return encoded, nil
}
