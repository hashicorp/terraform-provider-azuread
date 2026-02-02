package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesGradualRolloutSettings = WindowsUpdatesDateDrivenRolloutSettings{}

type WindowsUpdatesDateDrivenRolloutSettings struct {
	// Specifies the date before which all devices currently in the deployment are offered the update. Devices added after
	// this date are offered immediately. When the endDateTime isn't set, all devices in the deployment are offered content
	// at the same time.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// Fields inherited from WindowsUpdatesGradualRolloutSettings

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

func (s WindowsUpdatesDateDrivenRolloutSettings) WindowsUpdatesGradualRolloutSettings() BaseWindowsUpdatesGradualRolloutSettingsImpl {
	return BaseWindowsUpdatesGradualRolloutSettingsImpl{
		DurationBetweenOffers: s.DurationBetweenOffers,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesDateDrivenRolloutSettings{}

func (s WindowsUpdatesDateDrivenRolloutSettings) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDateDrivenRolloutSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDateDrivenRolloutSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDateDrivenRolloutSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.dateDrivenRolloutSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDateDrivenRolloutSettings: %+v", err)
	}

	return encoded, nil
}
