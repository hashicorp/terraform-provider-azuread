package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomClaimTransformation interface {
	CustomClaimTransformation() BaseCustomClaimTransformationImpl
}

var _ CustomClaimTransformation = BaseCustomClaimTransformationImpl{}

type BaseCustomClaimTransformationImpl struct {
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

func (s BaseCustomClaimTransformationImpl) CustomClaimTransformation() BaseCustomClaimTransformationImpl {
	return s
}

var _ CustomClaimTransformation = RawCustomClaimTransformationImpl{}

// RawCustomClaimTransformationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomClaimTransformationImpl struct {
	customClaimTransformation BaseCustomClaimTransformationImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawCustomClaimTransformationImpl) CustomClaimTransformation() BaseCustomClaimTransformationImpl {
	return s.customClaimTransformation
}

func UnmarshalCustomClaimTransformationImplementation(input []byte) (CustomClaimTransformation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaimTransformation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.containsTransformation") {
		var out ContainsTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContainsTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.endsWithTransformation") {
		var out EndsWithTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndsWithTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extractAlphaTransformation") {
		var out ExtractAlphaTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtractAlphaTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extractMailPrefixTransformation") {
		var out ExtractMailPrefixTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtractMailPrefixTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extractNumberTransformation") {
		var out ExtractNumberTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtractNumberTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.extractTransformation") {
		var out ExtractTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExtractTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ifEmptyTransformation") {
		var out IfEmptyTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IfEmptyTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ifNotEmptyTransformation") {
		var out IfNotEmptyTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IfNotEmptyTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.joinTransformation") {
		var out JoinTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JoinTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.regexReplaceTransformation") {
		var out RegexReplaceTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegexReplaceTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.startsWithTransformation") {
		var out StartsWithTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StartsWithTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.substringTransformation") {
		var out SubstringTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubstringTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.toLowercaseTransformation") {
		var out ToLowercaseTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ToLowercaseTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.toUppercaseTransformation") {
		var out ToUppercaseTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ToUppercaseTransformation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trimTransformation") {
		var out TrimTransformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrimTransformation: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomClaimTransformationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomClaimTransformationImpl: %+v", err)
	}

	return RawCustomClaimTransformationImpl{
		customClaimTransformation: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
