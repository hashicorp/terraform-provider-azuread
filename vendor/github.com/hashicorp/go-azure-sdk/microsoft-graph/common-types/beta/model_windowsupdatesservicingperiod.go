package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesServicingPeriod struct {
	// The date and time when the servicing period ends. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// The name of the servicing period. For example, Modern Lifecycle.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The start date and time of the servicing period. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	StartDateTime *string `json:"startDateTime,omitempty"`
}

var _ json.Marshaler = WindowsUpdatesServicingPeriod{}

func (s WindowsUpdatesServicingPeriod) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesServicingPeriod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesServicingPeriod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesServicingPeriod: %+v", err)
	}

	delete(decoded, "endDateTime")
	delete(decoded, "startDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesServicingPeriod: %+v", err)
	}

	return encoded, nil
}
