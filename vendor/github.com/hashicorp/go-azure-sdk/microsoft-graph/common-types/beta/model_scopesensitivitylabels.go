package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeSensitivityLabels interface {
	ScopeSensitivityLabels() BaseScopeSensitivityLabelsImpl
}

var _ ScopeSensitivityLabels = BaseScopeSensitivityLabelsImpl{}

type BaseScopeSensitivityLabelsImpl struct {
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

func (s BaseScopeSensitivityLabelsImpl) ScopeSensitivityLabels() BaseScopeSensitivityLabelsImpl {
	return s
}

var _ ScopeSensitivityLabels = RawScopeSensitivityLabelsImpl{}

// RawScopeSensitivityLabelsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScopeSensitivityLabelsImpl struct {
	scopeSensitivityLabels BaseScopeSensitivityLabelsImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawScopeSensitivityLabelsImpl) ScopeSensitivityLabels() BaseScopeSensitivityLabelsImpl {
	return s.scopeSensitivityLabels
}

func UnmarshalScopeSensitivityLabelsImplementation(input []byte) (ScopeSensitivityLabels, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScopeSensitivityLabels into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allScopeSensitivityLabels") {
		var out AllScopeSensitivityLabels
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllScopeSensitivityLabels: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enumeratedScopeSensitivityLabels") {
		var out EnumeratedScopeSensitivityLabels
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnumeratedScopeSensitivityLabels: %+v", err)
		}
		return out, nil
	}

	var parent BaseScopeSensitivityLabelsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScopeSensitivityLabelsImpl: %+v", err)
	}

	return RawScopeSensitivityLabelsImpl{
		scopeSensitivityLabels: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
