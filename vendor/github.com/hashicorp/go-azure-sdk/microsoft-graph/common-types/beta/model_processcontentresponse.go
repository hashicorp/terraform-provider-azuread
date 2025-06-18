package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessContentResponse struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A collection of policy actions (like DLP actions) triggered by the processed content. NOTE: Currently, the only
	// policyAction supported in processContentResponse is '_restrictAccess'
	PolicyActions *[]DlpActionInfo `json:"policyActions,omitempty"`

	// A collection of errors encountered during the content processing.
	ProcessingErrors *[]ProcessingError `json:"processingErrors,omitempty"`

	ProtectionScopeState *ProtectionScopeState `json:"protectionScopeState,omitempty"`
}

var _ json.Unmarshaler = &ProcessContentResponse{}

func (s *ProcessContentResponse) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
		ProcessingErrors     *[]ProcessingError    `json:"processingErrors,omitempty"`
		ProtectionScopeState *ProtectionScopeState `json:"protectionScopeState,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ProcessingErrors = decoded.ProcessingErrors
	s.ProtectionScopeState = decoded.ProtectionScopeState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ProcessContentResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["policyActions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PolicyActions into list []json.RawMessage: %+v", err)
		}

		output := make([]DlpActionInfo, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDlpActionInfoImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PolicyActions' for 'ProcessContentResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PolicyActions = &output
	}

	return nil
}
