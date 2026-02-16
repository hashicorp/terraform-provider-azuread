package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessContentRequest struct {
	ActivityMetadata *ActivityMetadata `json:"activityMetadata,omitempty"`

	// A collection of content entries to be processed. Each entry contains the content itself and its metadata. Use
	// conversation metadata for content like prompts and responses and file metadata for files. Required.
	ContentEntries []ProcessContentMetadataBase `json:"contentEntries"`

	IntegratedAppMetadata IntegratedApplicationMetadata `json:"integratedAppMetadata"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Metadata about the protected application making the request. Required.
	ProtectedAppMetadata ProtectedApplicationMetadata `json:"protectedAppMetadata"`
}

var _ json.Unmarshaler = &ProcessContentRequest{}

func (s *ProcessContentRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActivityMetadata     *ActivityMetadata            `json:"activityMetadata,omitempty"`
		ODataId              *string                      `json:"@odata.id,omitempty"`
		ODataType            *string                      `json:"@odata.type,omitempty"`
		ProtectedAppMetadata ProtectedApplicationMetadata `json:"protectedAppMetadata"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActivityMetadata = decoded.ActivityMetadata
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ProtectedAppMetadata = decoded.ProtectedAppMetadata

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ProcessContentRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["contentEntries"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ContentEntries into list []json.RawMessage: %+v", err)
		}

		output := make([]ProcessContentMetadataBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalProcessContentMetadataBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ContentEntries' for 'ProcessContentRequest': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ContentEntries = output
	}

	if v, ok := temp["integratedAppMetadata"]; ok {
		impl, err := UnmarshalIntegratedApplicationMetadataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IntegratedAppMetadata' for 'ProcessContentRequest': %+v", err)
		}
		s.IntegratedAppMetadata = impl
	}

	return nil
}
