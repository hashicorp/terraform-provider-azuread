package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LabelActionBase = ProtectOnlineMeetingAction{}

type ProtectOnlineMeetingAction struct {
	AllowedForwarders        *OnlineMeetingForwarders `json:"allowedForwarders,omitempty"`
	AllowedPresenters        *OnlineMeetingPresenters `json:"allowedPresenters,omitempty"`
	IsCopyToClipboardEnabled nullable.Type[bool]      `json:"isCopyToClipboardEnabled,omitempty"`
	IsLobbyEnabled           nullable.Type[bool]      `json:"isLobbyEnabled,omitempty"`
	LobbyBypassSettings      *LobbyBypassSettings     `json:"lobbyBypassSettings,omitempty"`

	// Fields inherited from LabelActionBase

	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ProtectOnlineMeetingAction) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProtectOnlineMeetingAction{}

func (s ProtectOnlineMeetingAction) MarshalJSON() ([]byte, error) {
	type wrapper ProtectOnlineMeetingAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProtectOnlineMeetingAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectOnlineMeetingAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectOnlineMeetingAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProtectOnlineMeetingAction: %+v", err)
	}

	return encoded, nil
}
