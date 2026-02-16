package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedLabel struct {
	// The display name of the label. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The unique identifier of the label.
	LabelId nullable.Type[string] `json:"labelId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = AssignedLabel{}

func (s AssignedLabel) MarshalJSON() ([]byte, error) {
	type wrapper AssignedLabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AssignedLabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AssignedLabel: %+v", err)
	}

	delete(decoded, "displayName")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AssignedLabel: %+v", err)
	}

	return encoded, nil
}
