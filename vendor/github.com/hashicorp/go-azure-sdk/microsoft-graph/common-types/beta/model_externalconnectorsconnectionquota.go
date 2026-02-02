package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ExternalConnectorsConnectionQuota{}

type ExternalConnectorsConnectionQuota struct {
	// The minimum of two values, one representing the items remaining in the connection and the other remaining items at
	// tenant-level. The following equation represents the formula used to calculate the minimum number: min ({max capacity
	// in the connection} – {number of items in the connection}, {tenant quota} – {number of items indexed in all
	// connections}). If the connection is not monetized, such as in a preview connector or preview content experience, then
	// this property is simply the number of remaining items in the connection.
	ItemsRemaining nullable.Type[int64] `json:"itemsRemaining,omitempty"`

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

func (s ExternalConnectorsConnectionQuota) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalConnectorsConnectionQuota{}

func (s ExternalConnectorsConnectionQuota) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnectorsConnectionQuota
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnectorsConnectionQuota: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsConnectionQuota: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.connectionQuota"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnectorsConnectionQuota: %+v", err)
	}

	return encoded, nil
}
