package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageResourceAttributeSource = AccessPackageResourceAttributeQuestion{}

type AccessPackageResourceAttributeQuestion struct {
	// The question asked in order to get the value of the attribute
	Question AccessPackageQuestion `json:"question"`

	// Fields inherited from AccessPackageResourceAttributeSource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessPackageResourceAttributeQuestion) AccessPackageResourceAttributeSource() BaseAccessPackageResourceAttributeSourceImpl {
	return BaseAccessPackageResourceAttributeSourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageResourceAttributeQuestion{}

func (s AccessPackageResourceAttributeQuestion) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResourceAttributeQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResourceAttributeQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceAttributeQuestion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResourceAttributeQuestion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceAttributeQuestion: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageResourceAttributeQuestion{}

func (s *AccessPackageResourceAttributeQuestion) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageResourceAttributeQuestion into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["question"]; ok {
		impl, err := UnmarshalAccessPackageQuestionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Question' for 'AccessPackageResourceAttributeQuestion': %+v", err)
		}
		s.Question = impl
	}

	return nil
}
