package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosHomeScreenPage struct {
	// Name of the page
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A list of apps, folders, and web clips to appear on a page. This collection can contain a maximum of 500 elements.
	Icons *[]IosHomeScreenItem `json:"icons,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &IosHomeScreenPage{}

func (s *IosHomeScreenPage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosHomeScreenPage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["icons"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Icons into list []json.RawMessage: %+v", err)
		}

		output := make([]IosHomeScreenItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIosHomeScreenItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Icons' for 'IosHomeScreenPage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Icons = &output
	}

	return nil
}
