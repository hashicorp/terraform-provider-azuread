package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesApprovalRule interface {
	WindowsUpdatesApprovalRule() BaseWindowsUpdatesApprovalRuleImpl
}

var _ WindowsUpdatesApprovalRule = BaseWindowsUpdatesApprovalRuleImpl{}

type BaseWindowsUpdatesApprovalRuleImpl struct {
	DeferralInDays *int64 `json:"deferralInDays,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesApprovalRuleImpl) WindowsUpdatesApprovalRule() BaseWindowsUpdatesApprovalRuleImpl {
	return s
}

var _ WindowsUpdatesApprovalRule = RawWindowsUpdatesApprovalRuleImpl{}

// RawWindowsUpdatesApprovalRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesApprovalRuleImpl struct {
	windowsUpdatesApprovalRule BaseWindowsUpdatesApprovalRuleImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawWindowsUpdatesApprovalRuleImpl) WindowsUpdatesApprovalRule() BaseWindowsUpdatesApprovalRuleImpl {
	return s.windowsUpdatesApprovalRule
}

func UnmarshalWindowsUpdatesApprovalRuleImplementation(input []byte) (WindowsUpdatesApprovalRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesApprovalRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.qualityUpdateApprovalRule") {
		var out WindowsUpdatesQualityUpdateApprovalRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesQualityUpdateApprovalRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesApprovalRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesApprovalRuleImpl: %+v", err)
	}

	return RawWindowsUpdatesApprovalRuleImpl{
		windowsUpdatesApprovalRule: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
