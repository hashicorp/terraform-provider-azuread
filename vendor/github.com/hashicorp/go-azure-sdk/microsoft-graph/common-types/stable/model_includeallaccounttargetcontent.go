package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccountTargetContent = IncludeAllAccountTargetContent{}

type IncludeAllAccountTargetContent struct {

	// Fields inherited from AccountTargetContent

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of account target content. Possible values are: unknown, includeAll, addressBook, unknownFutureValue.
	Type *AccountTargetContentType `json:"type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IncludeAllAccountTargetContent) AccountTargetContent() BaseAccountTargetContentImpl {
	return BaseAccountTargetContentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Type:      s.Type,
	}
}

var _ json.Marshaler = IncludeAllAccountTargetContent{}

func (s IncludeAllAccountTargetContent) MarshalJSON() ([]byte, error) {
	type wrapper IncludeAllAccountTargetContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IncludeAllAccountTargetContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IncludeAllAccountTargetContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.includeAllAccountTargetContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IncludeAllAccountTargetContent: %+v", err)
	}

	return encoded, nil
}
