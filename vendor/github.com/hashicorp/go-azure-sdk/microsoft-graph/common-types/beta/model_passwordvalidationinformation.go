package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordValidationInformation struct {
	// Specifies whether the password is valid based on the calculation of the results in the validationResults property.
	// Not nullable. Read-only.
	IsValid nullable.Type[bool] `json:"isValid,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
	ValidationResults *[]ValidationResult `json:"validationResults,omitempty"`
}

var _ json.Marshaler = PasswordValidationInformation{}

func (s PasswordValidationInformation) MarshalJSON() ([]byte, error) {
	type wrapper PasswordValidationInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PasswordValidationInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PasswordValidationInformation: %+v", err)
	}

	delete(decoded, "isValid")
	delete(decoded, "validationResults")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PasswordValidationInformation: %+v", err)
	}

	return encoded, nil
}
