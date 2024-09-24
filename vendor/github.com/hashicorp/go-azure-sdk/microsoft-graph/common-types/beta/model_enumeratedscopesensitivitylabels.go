package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScopeSensitivityLabels = EnumeratedScopeSensitivityLabels{}

type EnumeratedScopeSensitivityLabels struct {
	// The sensitivity labels that are applicable to the scope type and have been preapproved. Required.
	SensitivityLabels []string `json:"sensitivityLabels"`

	// Fields inherited from ScopeSensitivityLabels

	// Indicates the kind of sensitivity label that is included. Possible values: all means all sensitivity labels are
	// allowed, or enumerated means a selected set of sensitivity labels from a single resource application are allowed.
	// Required.
	LabelKind LabelKind `json:"labelKind"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EnumeratedScopeSensitivityLabels) ScopeSensitivityLabels() BaseScopeSensitivityLabelsImpl {
	return BaseScopeSensitivityLabelsImpl{
		LabelKind: s.LabelKind,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnumeratedScopeSensitivityLabels{}

func (s EnumeratedScopeSensitivityLabels) MarshalJSON() ([]byte, error) {
	type wrapper EnumeratedScopeSensitivityLabels
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnumeratedScopeSensitivityLabels: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnumeratedScopeSensitivityLabels: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enumeratedScopeSensitivityLabels"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnumeratedScopeSensitivityLabels: %+v", err)
	}

	return encoded, nil
}
