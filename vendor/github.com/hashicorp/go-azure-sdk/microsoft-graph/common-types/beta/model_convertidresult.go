package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConvertIdResult struct {
	// An error object indicating the reason for the conversion failure. This value isn't present if the conversion
	// succeeded.
	ErrorDetails GenericError `json:"errorDetails"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier that was converted. This value is the original, unconverted identifier.
	SourceId nullable.Type[string] `json:"sourceId,omitempty"`

	// The converted identifier. This value isn't present if the conversion failed.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`
}

var _ json.Unmarshaler = &ConvertIdResult{}

func (s *ConvertIdResult) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
		SourceId  nullable.Type[string] `json:"sourceId,omitempty"`
		TargetId  nullable.Type[string] `json:"targetId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SourceId = decoded.SourceId
	s.TargetId = decoded.TargetId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ConvertIdResult into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["errorDetails"]; ok {
		impl, err := UnmarshalGenericErrorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ErrorDetails' for 'ConvertIdResult': %+v", err)
		}
		s.ErrorDetails = impl
	}

	return nil
}
