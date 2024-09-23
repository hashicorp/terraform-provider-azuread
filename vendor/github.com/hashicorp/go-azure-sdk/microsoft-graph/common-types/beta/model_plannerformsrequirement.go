package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerFormsRequirement struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Read-only. A collection of keys from the plannerFormsDictionary that identify the plannerFormReference objects that
	// specify the requirements to complete the plannerTask.
	RequiredForms *[]string `json:"requiredForms,omitempty"`
}

var _ json.Marshaler = PlannerFormsRequirement{}

func (s PlannerFormsRequirement) MarshalJSON() ([]byte, error) {
	type wrapper PlannerFormsRequirement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerFormsRequirement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerFormsRequirement: %+v", err)
	}

	delete(decoded, "requiredForms")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerFormsRequirement: %+v", err)
	}

	return encoded, nil
}
