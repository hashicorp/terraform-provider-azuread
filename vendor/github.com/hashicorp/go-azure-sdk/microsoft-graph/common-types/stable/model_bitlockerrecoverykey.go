package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BitlockerRecoveryKey{}

type BitlockerRecoveryKey struct {
	// The date and time when the key was originally backed up to Microsoft Entra ID. Not nullable.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Identifier of the device the BitLocker key is originally backed up from. Supports $filter (eq).
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The BitLocker recovery key. Returned only on $select. Not nullable.
	Key *string `json:"key,omitempty"`

	// Indicates the type of volume the BitLocker key is associated with. The possible values are: 1 (for
	// operatingSystemVolume), 2 (for fixedDataVolume), 3 (for removableDataVolume), and 4 (for unknownFutureValue).
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
