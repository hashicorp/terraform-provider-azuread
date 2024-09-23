package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesGradualRolloutSettings interface {
	WindowsUpdatesGradualRolloutSettings() BaseWindowsUpdatesGradualRolloutSettingsImpl
}

var _ WindowsUpdatesGradualRolloutSettings = BaseWindowsUpdatesGradualRolloutSettingsImpl{}

type BaseWindowsUpdatesGradualRolloutSettingsImpl struct {
	// The duration between each set of devices being offered the update. The value is represented in ISO 8601 format for
	// duration. Default value is P1D (one day).
	DurationBetweenOffers nullable.Type[string] `json:"durationBetweenOffers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesGradualRolloutSettingsImpl) WindowsUpdatesGradualRolloutSettings() BaseWindowsUpdatesGradualRolloutSettingsImpl {
	return s
}

var _ WindowsUpdatesGradualRolloutSettings = RawWindowsUpdatesGradualRolloutSettingsImpl{}

// RawWindowsUpdatesGradualRolloutSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesGradualRolloutSettingsImpl struct {
	windowsUpdatesGradualRolloutSettings BaseWindowsUpdatesGradualRolloutSettingsImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawWindowsUpdatesGradualRolloutSettingsImpl) WindowsUpdatesGradualRolloutSettings() BaseWindowsUpdatesGradualRolloutSettingsImpl {
	return s.windowsUpdatesGradualRolloutSettings
}

func UnmarshalWindowsUpdatesGradualRolloutSettingsImplementation(input []byte) (WindowsUpdatesGradualRolloutSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesGradualRolloutSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.dateDrivenRolloutSettings") {
		var out WindowsUpdatesDateDrivenRolloutSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDateDrivenRolloutSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.durationDrivenRolloutSettings") {
		var out WindowsUpdatesDurationDrivenRolloutSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDurationDrivenRolloutSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.rateDrivenRolloutSettings") {
		var out WindowsUpdatesRateDrivenRolloutSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesRateDrivenRolloutSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesGradualRolloutSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesGradualRolloutSettingsImpl: %+v", err)
	}

	return RawWindowsUpdatesGradualRolloutSettingsImpl{
		windowsUpdatesGradualRolloutSettings: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
