package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchSensitivityLabelInfo struct {
	Color       nullable.Type[string] `json:"color,omitempty"`
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Priority           nullable.Type[int64]  `json:"priority,omitempty"`
	SensitivityLabelId nullable.Type[string] `json:"sensitivityLabelId,omitempty"`
	Tooltip            nullable.Type[string] `json:"tooltip,omitempty"`
}

var _ json.Marshaler = SearchSensitivityLabelInfo{}

func (s SearchSensitivityLabelInfo) MarshalJSON() ([]byte, error) {
	type wrapper SearchSensitivityLabelInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SearchSensitivityLabelInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SearchSensitivityLabelInfo: %+v", err)
	}

	delete(decoded, "color")
	delete(decoded, "displayName")
	delete(decoded, "priority")
	delete(decoded, "sensitivityLabelId")
	delete(decoded, "tooltip")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SearchSensitivityLabelInfo: %+v", err)
	}

	return encoded, nil
}
