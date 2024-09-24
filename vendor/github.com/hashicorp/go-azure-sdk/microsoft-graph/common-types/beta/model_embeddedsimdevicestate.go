package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EmbeddedSIMDeviceState{}

type EmbeddedSIMDeviceState struct {
	// The time the embedded SIM device status was created. Generated service side.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Device name to which the subscription was provisioned e.g. DESKTOP-JOE
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The time the embedded SIM device last checked in. Updated service side.
	LastSyncDateTime nullable.Type[string] `json:"lastSyncDateTime,omitempty"`

	// The time the embedded SIM device status was last modified. Updated service side.
	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

	// Describes the various states for an embedded SIM activation code.
	State *EmbeddedSIMDeviceStateValue `json:"state,omitempty"`

	// String description of the provisioning state.
	StateDetails nullable.Type[string] `json:"stateDetails,omitempty"`

	// The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be
	// deployed.
	UniversalIntegratedCircuitCardIdentifier nullable.Type[string] `json:"universalIntegratedCircuitCardIdentifier,omitempty"`

	// Username which the subscription was provisioned to e.g. joe@contoso.com
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s EmbeddedSIMDeviceState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EmbeddedSIMDeviceState{}

func (s EmbeddedSIMDeviceState) MarshalJSON() ([]byte, error) {
	type wrapper EmbeddedSIMDeviceState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmbeddedSIMDeviceState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmbeddedSIMDeviceState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.embeddedSIMDeviceState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmbeddedSIMDeviceState: %+v", err)
	}

	return encoded, nil
}
