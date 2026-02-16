package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationResult struct {
	// The string containing the reason for why the rule passed or not. Read-only. Not nullable.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The string containing the name of the password validation rule that the action was validated against. Read-only. Not
	// nullable.
	RuleName nullable.Type[string] `json:"ruleName,omitempty"`

	// Whether the password passed or failed the validation rule. Read-only. Not nullable.
	ValidationPassed nullable.Type[bool] `json:"validationPassed,omitempty"`
}

var _ json.Marshaler = ValidationResult{}

func (s ValidationResult) MarshalJSON() ([]byte, error) {
	type wrapper ValidationResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidationResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidationResult: %+v", err)
	}

	delete(decoded, "message")
	delete(decoded, "ruleName")
	delete(decoded, "validationPassed")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidationResult: %+v", err)
	}

	return encoded, nil
}
