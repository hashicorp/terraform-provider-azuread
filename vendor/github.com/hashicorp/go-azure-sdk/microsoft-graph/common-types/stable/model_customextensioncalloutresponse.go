package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutResponse struct {
	// Contains the data the external system provides to the custom extension endpoint.
	Data CustomExtensionData `json:"data"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Identifies the external system or event context related to the response.
	Source nullable.Type[string] `json:"source,omitempty"`

	// Describes the type of event related to the response.
	Type nullable.Type[string] `json:"type,omitempty"`
}

var _ json.Unmarshaler = &CustomExtensionCalloutResponse{}

func (s *CustomExtensionCalloutResponse) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
		Source    nullable.Type[string] `json:"source,omitempty"`
		Type      nullable.Type[string] `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CustomExtensionCalloutResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["data"]; ok {
		impl, err := UnmarshalCustomExtensionDataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Data' for 'CustomExtensionCalloutResponse': %+v", err)
		}
		s.Data = impl
	}

	return nil
}
