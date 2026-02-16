package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceComplianceScriptError = DeviceComplianceScriptRuleError{}

type DeviceComplianceScriptRuleError struct {
	// Setting name for the rule with error.
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// Fields inherited from DeviceComplianceScriptError

	// Error code for rule validation.
	Code *Code `json:"code,omitempty"`

	// Error code for rule validation.
	DeviceComplianceScriptRulesValidationError *DeviceComplianceScriptRulesValidationError `json:"deviceComplianceScriptRulesValidationError,omitempty"`

	// Error message.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceComplianceScriptRuleError) DeviceComplianceScriptError() BaseDeviceComplianceScriptErrorImpl {
	return BaseDeviceComplianceScriptErrorImpl{
		Code: s.Code,
		DeviceComplianceScriptRulesValidationError: s.DeviceComplianceScriptRulesValidationError,
		Message:   s.Message,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceComplianceScriptRuleError{}

func (s DeviceComplianceScriptRuleError) MarshalJSON() ([]byte, error) {
	type wrapper DeviceComplianceScriptRuleError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceComplianceScriptRuleError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceComplianceScriptRuleError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceComplianceScriptRuleError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceComplianceScriptRuleError: %+v", err)
	}

	return encoded, nil
}
