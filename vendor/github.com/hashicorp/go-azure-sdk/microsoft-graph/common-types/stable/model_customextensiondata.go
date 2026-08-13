package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionData interface {
	CustomExtensionData() BaseCustomExtensionDataImpl
}

var _ CustomExtensionData = BaseCustomExtensionDataImpl{}

type BaseCustomExtensionDataImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomExtensionDataImpl) CustomExtensionData() BaseCustomExtensionDataImpl {
	return s
}

var _ CustomExtensionData = RawCustomExtensionDataImpl{}

// RawCustomExtensionDataImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCustomExtensionDataImpl struct {
	customExtensionData BaseCustomExtensionDataImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawCustomExtensionDataImpl) CustomExtensionData() BaseCustomExtensionDataImpl {
	return s.customExtensionData
}

func (s RawCustomExtensionDataImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCustomExtensionDataImplementation(input []byte) (CustomExtensionData, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionData into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAssignmentRequestCallbackData") {
		var out AccessPackageAssignmentRequestCallbackData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAssignmentRequestCallbackData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.customTaskExtensionCallbackData") {
		var out IdentityGovernanceCustomTaskExtensionCallbackData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceCustomTaskExtensionCallbackData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.customTaskExtensionCalloutData") {
		var out IdentityGovernanceCustomTaskExtensionCalloutData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceCustomTaskExtensionCalloutData: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomExtensionDataImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomExtensionDataImpl: %+v", err)
	}

	return RawCustomExtensionDataImpl{
		customExtensionData: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
