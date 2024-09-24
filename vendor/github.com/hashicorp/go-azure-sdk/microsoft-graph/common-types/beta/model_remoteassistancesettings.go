package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RemoteAssistanceSettings{}

type RemoteAssistanceSettings struct {
	// Indicates if sessions to unenrolled devices are allowed for the account. This setting is configurable by the admin.
	// Default value is false.
	AllowSessionsToUnenrolledDevices *bool `json:"allowSessionsToUnenrolledDevices,omitempty"`

	// Indicates if sessions to block chat function. This setting is configurable by the admin. Default value is false.
	BlockChat *bool `json:"blockChat,omitempty"`

	// State of remote assistance for the account
	RemoteAssistanceState *RemoteAssistanceState `json:"remoteAssistanceState,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RemoteAssistanceSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RemoteAssistanceSettings{}

func (s RemoteAssistanceSettings) MarshalJSON() ([]byte, error) {
	type wrapper RemoteAssistanceSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RemoteAssistanceSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RemoteAssistanceSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.remoteAssistanceSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RemoteAssistanceSettings: %+v", err)
	}

	return encoded, nil
}
