package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunEntityCountMetric struct {
	// The count of entries for the entity marked as Active.
	Active nullable.Type[int64] `json:"active,omitempty"`

	// The count of entries for the entity marked as Inactive.
	Inactive nullable.Type[int64] `json:"inactive,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total count of the entity.
	Total nullable.Type[int64] `json:"total,omitempty"`
}

var _ json.Marshaler = IndustryDataIndustryDataRunEntityCountMetric{}

func (s IndustryDataIndustryDataRunEntityCountMetric) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataIndustryDataRunEntityCountMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataIndustryDataRunEntityCountMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataRunEntityCountMetric: %+v", err)
	}

	delete(decoded, "active")
	delete(decoded, "inactive")
	delete(decoded, "total")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataIndustryDataRunEntityCountMetric: %+v", err)
	}

	return encoded, nil
}
