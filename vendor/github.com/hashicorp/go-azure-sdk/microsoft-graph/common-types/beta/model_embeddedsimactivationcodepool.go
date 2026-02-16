package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EmbeddedSIMActivationCodePool{}

type EmbeddedSIMActivationCodePool struct {
	// The total count of activation codes which belong to this pool.
	ActivationCodeCount *int64 `json:"activationCodeCount,omitempty"`

	// The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune
	// but cannot be used to read activation codes from Intune.
	ActivationCodes *[]EmbeddedSIMActivationCode `json:"activationCodes,omitempty"`

	// Navigational property to a list of targets to which this pool is assigned.
	Assignments *[]EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`

	// The time the embedded SIM activation code pool was created. Generated service side.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Navigational property to a list of device states for this pool.
	DeviceStates *[]EmbeddedSIMDeviceState `json:"deviceStates,omitempty"`

	// The admin defined name of the embedded SIM activation code pool.
	DisplayName *string `json:"displayName,omitempty"`

	// The time the embedded SIM activation code pool was last modified. Updated service side.
	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

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

func (s EmbeddedSIMActivationCodePool) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EmbeddedSIMActivationCodePool{}

func (s EmbeddedSIMActivationCodePool) MarshalJSON() ([]byte, error) {
	type wrapper EmbeddedSIMActivationCodePool
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmbeddedSIMActivationCodePool: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmbeddedSIMActivationCodePool: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.embeddedSIMActivationCodePool"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmbeddedSIMActivationCodePool: %+v", err)
	}

	return encoded, nil
}
