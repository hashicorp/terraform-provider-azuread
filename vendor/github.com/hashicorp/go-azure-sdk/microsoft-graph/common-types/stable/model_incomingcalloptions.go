package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallOptions = IncomingCallOptions{}

type IncomingCallOptions struct {

	// Fields inherited from CallOptions

	// Indicates whether to hide the app after the call is escalated.
	HideBotAfterEscalation nullable.Type[bool] `json:"hideBotAfterEscalation,omitempty"`

	// Indicates whether content sharing notifications should be enabled for the call.
	IsContentSharingNotificationEnabled nullable.Type[bool] `json:"isContentSharingNotificationEnabled,omitempty"`

	// Indicates whether delta roster is enabled for the call.
	IsDeltaRosterEnabled nullable.Type[bool] `json:"isDeltaRosterEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IncomingCallOptions) CallOptions() BaseCallOptionsImpl {
	return BaseCallOptionsImpl{
		HideBotAfterEscalation:              s.HideBotAfterEscalation,
		IsContentSharingNotificationEnabled: s.IsContentSharingNotificationEnabled,
		IsDeltaRosterEnabled:                s.IsDeltaRosterEnabled,
		ODataId:                             s.ODataId,
		ODataType:                           s.ODataType,
	}
}

var _ json.Marshaler = IncomingCallOptions{}

func (s IncomingCallOptions) MarshalJSON() ([]byte, error) {
	type wrapper IncomingCallOptions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IncomingCallOptions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IncomingCallOptions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.incomingCallOptions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IncomingCallOptions: %+v", err)
	}

	return encoded, nil
}
