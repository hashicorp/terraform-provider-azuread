package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesComplianceChangeRule interface {
	WindowsUpdatesComplianceChangeRule() BaseWindowsUpdatesComplianceChangeRuleImpl
}

var _ WindowsUpdatesComplianceChangeRule = BaseWindowsUpdatesComplianceChangeRuleImpl{}

type BaseWindowsUpdatesComplianceChangeRuleImpl struct {
	// The date and time when the rule was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time when the rule was last evaluated.
	LastEvaluatedDateTime nullable.Type[string] `json:"lastEvaluatedDateTime,omitempty"`

	// The date and time when the rule was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesComplianceChangeRuleImpl) WindowsUpdatesComplianceChangeRule() BaseWindowsUpdatesComplianceChangeRuleImpl {
	return s
}

var _ WindowsUpdatesComplianceChangeRule = RawWindowsUpdatesComplianceChangeRuleImpl{}

// RawWindowsUpdatesComplianceChangeRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesComplianceChangeRuleImpl struct {
	windowsUpdatesComplianceChangeRule BaseWindowsUpdatesComplianceChangeRuleImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawWindowsUpdatesComplianceChangeRuleImpl) WindowsUpdatesComplianceChangeRule() BaseWindowsUpdatesComplianceChangeRuleImpl {
	return s.windowsUpdatesComplianceChangeRule
}

func UnmarshalWindowsUpdatesComplianceChangeRuleImplementation(input []byte) (WindowsUpdatesComplianceChangeRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesComplianceChangeRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.contentApprovalRule") {
		var out WindowsUpdatesContentApprovalRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesContentApprovalRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesComplianceChangeRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesComplianceChangeRuleImpl: %+v", err)
	}

	return RawWindowsUpdatesComplianceChangeRuleImpl{
		windowsUpdatesComplianceChangeRule: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
