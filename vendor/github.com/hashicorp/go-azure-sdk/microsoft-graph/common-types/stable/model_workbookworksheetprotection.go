package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookWorksheetProtection{}

type WorkbookWorksheetProtection struct {
	// Worksheet protection options. Read-only.
	Options *WorkbookWorksheetProtectionOptions `json:"options,omitempty"`

	// Indicates whether the worksheet is protected. Read-only.
	Protected *bool `json:"protected,omitempty"`

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

func (s WorkbookWorksheetProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookWorksheetProtection{}

func (s WorkbookWorksheetProtection) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookWorksheetProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookWorksheetProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookWorksheetProtection: %+v", err)
	}

	delete(decoded, "options")
	delete(decoded, "protected")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookWorksheetProtection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookWorksheetProtection: %+v", err)
	}

	return encoded, nil
}
