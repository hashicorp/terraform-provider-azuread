package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataFilter = IndustryDataBasicFilter{}

type IndustryDataBasicFilter struct {
	Attribute *IndustryDataFilterOptions `json:"attribute,omitempty"`

	// The condition to filter with.
	In *[]string `json:"in,omitempty"`

	// Fields inherited from IndustryDataFilter

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IndustryDataBasicFilter) IndustryDataFilter() BaseIndustryDataFilterImpl {
	return BaseIndustryDataFilterImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataBasicFilter{}

func (s IndustryDataBasicFilter) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataBasicFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataBasicFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataBasicFilter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.basicFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataBasicFilter: %+v", err)
	}

	return encoded, nil
}
