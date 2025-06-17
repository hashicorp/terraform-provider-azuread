package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookFormatProtection{}

type WorkbookFormatProtection struct {
	// Indicates whether Excel hides the formula for the cells in the range. A null value indicates that the entire range
	// doesn't have uniform formula hidden setting.
	FormulaHidden nullable.Type[bool] `json:"formulaHidden,omitempty"`

	// Indicates whether Excel locks the cells in the object. A null value indicates that the entire range doesn't have
	// uniform lock setting.
	Locked nullable.Type[bool] `json:"locked,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WorkbookFormatProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookFormatProtection{}

func (s WorkbookFormatProtection) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookFormatProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookFormatProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookFormatProtection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookFormatProtection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookFormatProtection: %+v", err)
	}

	return encoded, nil
}
