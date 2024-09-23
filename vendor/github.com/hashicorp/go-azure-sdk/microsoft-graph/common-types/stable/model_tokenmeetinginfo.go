package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingInfo = TokenMeetingInfo{}

type TokenMeetingInfo struct {
	// The token used to join the call.
	Token *string `json:"token,omitempty"`

	// Fields inherited from MeetingInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TokenMeetingInfo) MeetingInfo() BaseMeetingInfoImpl {
	return BaseMeetingInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TokenMeetingInfo{}

func (s TokenMeetingInfo) MarshalJSON() ([]byte, error) {
	type wrapper TokenMeetingInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TokenMeetingInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TokenMeetingInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tokenMeetingInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TokenMeetingInfo: %+v", err)
	}

	return encoded, nil
}
