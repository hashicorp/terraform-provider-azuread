package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataReferenceValue = IndustryDataIdentifierTypeReferenceValue{}

type IndustryDataIdentifierTypeReferenceValue struct {

	// Fields inherited from IndustryDataReferenceValue

	// The code of the desired referenceDefinition entry.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Value *IndustryDataReferenceDefinition `json:"value,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IndustryDataIdentifierTypeReferenceValue) IndustryDataReferenceValue() BaseIndustryDataReferenceValueImpl {
	return BaseIndustryDataReferenceValueImpl{
		Code:      s.Code,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Value:     s.Value,
	}
}

var _ json.Marshaler = IndustryDataIdentifierTypeReferenceValue{}

func (s IndustryDataIdentifierTypeReferenceValue) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataIdentifierTypeReferenceValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataIdentifierTypeReferenceValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIdentifierTypeReferenceValue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.identifierTypeReferenceValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataIdentifierTypeReferenceValue: %+v", err)
	}

	return encoded, nil
}
