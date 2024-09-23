package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCSnapshot{}

type CloudPCSnapshot struct {
	// The unique identifier for the Cloud PC.
	CloudPCId *string `json:"cloudPcId,omitempty"`

	// The date and time at which the snapshot was taken. The timestamp is shown in ISO 8601 format and Coordinated
	// Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time when the snapshot expires. The time is shown in ISO 8601 format and Coordinated Universal Time
	// (UTC) time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The date and time at which the snapshot was last used to restore the Cloud PC device. The timestamp is shown in ISO
	// 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastRestoredDateTime nullable.Type[string] `json:"lastRestoredDateTime,omitempty"`

	// The type of snapshot that indicates how to create the snapshot. Possible values are automatic, manual. Default value
	// is automatic.
	SnapshotType *CloudPCSnapshotType `json:"snapshotType,omitempty"`

	// The status of the Cloud PC snapshot. The possible values are: ready, unknownFutureValue.
	Status *CloudPCSnapshotStatus `json:"status,omitempty"`

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

func (s CloudPCSnapshot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCSnapshot{}

func (s CloudPCSnapshot) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCSnapshot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCSnapshot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCSnapshot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcSnapshot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCSnapshot: %+v", err)
	}

	return encoded, nil
}
