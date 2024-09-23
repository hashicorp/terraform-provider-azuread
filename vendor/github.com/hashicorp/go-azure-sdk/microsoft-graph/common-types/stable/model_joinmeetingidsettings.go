package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinMeetingIdSettings struct {
	// Indicates whether a passcode is required to join a meeting when using joinMeetingId. Optional.
	IsPasscodeRequired nullable.Type[bool] `json:"isPasscodeRequired,omitempty"`

	// The meeting ID to be used to join a meeting. Optional. Read-only.
	JoinMeetingId nullable.Type[string] `json:"joinMeetingId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The passcode to join a meeting. Optional. Read-only.
	Passcode nullable.Type[string] `json:"passcode,omitempty"`
}

var _ json.Marshaler = JoinMeetingIdSettings{}

func (s JoinMeetingIdSettings) MarshalJSON() ([]byte, error) {
	type wrapper JoinMeetingIdSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling JoinMeetingIdSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling JoinMeetingIdSettings: %+v", err)
	}

	delete(decoded, "joinMeetingId")
	delete(decoded, "passcode")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling JoinMeetingIdSettings: %+v", err)
	}

	return encoded, nil
}
