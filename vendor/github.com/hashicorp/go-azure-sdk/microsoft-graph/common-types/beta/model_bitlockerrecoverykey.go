package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BitlockerRecoveryKey{}

type BitlockerRecoveryKey struct {
	// The date and time when the key was originally backed up to Microsoft Entra ID.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// ID of the device the BitLocker key is originally backed up from.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The BitLocker recovery key.
	Key *string `json:"key,omitempty"`

	// Indicates the type of volume the BitLocker key is associated with. Possible values are: operatingSystemVolume,
	// fixedDataVolume, removableDataVolume, unknownFutureValue.
	VolumeType *VolumeType `json:"volumeType,omitempty"`

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

func (s BitlockerRecoveryKey) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BitlockerRecoveryKey{}

func (s BitlockerRecoveryKey) MarshalJSON() ([]byte, error) {
	type wrapper BitlockerRecoveryKey
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BitlockerRecoveryKey: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BitlockerRecoveryKey: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bitlockerRecoveryKey"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BitlockerRecoveryKey: %+v", err)
	}

	return encoded, nil
}
