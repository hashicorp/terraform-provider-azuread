package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenActionResult interface {
	VppTokenActionResult() BaseVppTokenActionResultImpl
}

var _ VppTokenActionResult = BaseVppTokenActionResultImpl{}

type BaseVppTokenActionResultImpl struct {
	// Action name
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	ActionState *ActionState `json:"actionState,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time the action was initiated
	StartDateTime *string `json:"startDateTime,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseVppTokenActionResultImpl) VppTokenActionResult() BaseVppTokenActionResultImpl {
	return s
}

var _ VppTokenActionResult = RawVppTokenActionResultImpl{}

// RawVppTokenActionResultImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawVppTokenActionResultImpl struct {
	vppTokenActionResult BaseVppTokenActionResultImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawVppTokenActionResultImpl) VppTokenActionResult() BaseVppTokenActionResultImpl {
	return s.vppTokenActionResult
}

func UnmarshalVppTokenActionResultImplementation(input []byte) (VppTokenActionResult, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling VppTokenActionResult into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.vppTokenRevokeLicensesActionResult") {
		var out VppTokenRevokeLicensesActionResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VppTokenRevokeLicensesActionResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseVppTokenActionResultImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseVppTokenActionResultImpl: %+v", err)
	}

	return RawVppTokenActionResultImpl{
		vppTokenActionResult: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
