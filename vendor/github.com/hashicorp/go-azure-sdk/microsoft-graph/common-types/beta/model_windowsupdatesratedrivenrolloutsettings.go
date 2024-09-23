package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesGradualRolloutSettings = WindowsUpdatesRateDrivenRolloutSettings{}

type WindowsUpdatesRateDrivenRolloutSettings struct {
	// Specifies the number of devices that are offered at the same time. When not set, all devices in the deployment are
	// offered content at the same time.
	DevicesPerOffer *int64 `json:"devicesPerOffer,omitempty"`

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

func (s WindowsUpdatesRateDrivenRolloutSettings) WindowsUpdatesGradualRolloutSettings() BaseWindowsUpdatesGradualRolloutSettingsImpl {
	return BaseWindowsUpdatesGradualRolloutSettingsImpl{
		DurationBetweenOffers: s.DurationBetweenOffers,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesRateDrivenRolloutSettings{}

func (s WindowsUpdatesRateDrivenRolloutSettings) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesRateDrivenRolloutSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesRateDrivenRolloutSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesRateDrivenRolloutSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.rateDrivenRolloutSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesRateDrivenRolloutSettings: %+v", err)
	}

	return encoded, nil
}
