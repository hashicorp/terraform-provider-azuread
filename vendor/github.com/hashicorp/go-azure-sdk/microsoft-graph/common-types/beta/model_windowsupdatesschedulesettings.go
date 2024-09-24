package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesScheduleSettings struct {
	// Settings for governing how to rollout content to devices. One of:
	// microsoft.graph.windowsUpdates.dateDrivenRolloutSettings,
	// microsoft.graph.windowsUpdates.durationDrivenRolloutSettings, or
	// microsoft.graph.windowsUpdates.rateDrivenRolloutSettings.
	GradualRollout WindowsUpdatesGradualRolloutSettings `json:"gradualRollout"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date on which devices in the deployment start receiving the update. When not set, the deployment starts as soon
	// as devices are assigned. The Timestamp type represents date and time information using ISO 8601 format and is always
	// in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}

var _ json.Unmarshaler = &WindowsUpdatesScheduleSettings{}

func (s *WindowsUpdatesScheduleSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId       *string               `json:"@odata.id,omitempty"`
		ODataType     *string               `json:"@odata.type,omitempty"`
		StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.StartDateTime = decoded.StartDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesScheduleSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["gradualRollout"]; ok {
		impl, err := UnmarshalWindowsUpdatesGradualRolloutSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GradualRollout' for 'WindowsUpdatesScheduleSettings': %+v", err)
		}
		s.GradualRollout = impl
	}

	return nil
}
