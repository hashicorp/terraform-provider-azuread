package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConfigurationValidation struct {
	// Errors in the validation result of a customAuthenticationExtension.
	Errors *[]GenericError `json:"errors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Warnings in the validation result of a customAuthenticationExtension.
	Warnings *[]GenericError `json:"warnings,omitempty"`
}

var _ json.Unmarshaler = &AuthenticationConfigurationValidation{}

func (s *AuthenticationConfigurationValidation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AuthenticationConfigurationValidation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["errors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Errors into list []json.RawMessage: %+v", err)
		}

		output := make([]GenericError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGenericErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Errors' for 'AuthenticationConfigurationValidation': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Errors = &output
	}

	if v, ok := temp["warnings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Warnings into list []json.RawMessage: %+v", err)
		}

		output := make([]GenericError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGenericErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Warnings' for 'AuthenticationConfigurationValidation': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Warnings = &output
	}

	return nil
}
