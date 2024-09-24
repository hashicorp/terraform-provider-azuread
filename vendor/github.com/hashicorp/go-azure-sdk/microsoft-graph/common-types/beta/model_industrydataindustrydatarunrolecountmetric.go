package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunRoleCountMetric struct {
	// The number of people in this role.
	Count nullable.Type[int64] `json:"count,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The role that is being measured.
	Role *string `json:"role,omitempty"`
}

var _ json.Marshaler = IndustryDataIndustryDataRunRoleCountMetric{}

func (s IndustryDataIndustryDataRunRoleCountMetric) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataIndustryDataRunRoleCountMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataIndustryDataRunRoleCountMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataRunRoleCountMetric: %+v", err)
	}

	delete(decoded, "count")
	delete(decoded, "role")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataIndustryDataRunRoleCountMetric: %+v", err)
	}

	return encoded, nil
}
