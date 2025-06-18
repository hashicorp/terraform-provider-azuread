package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessLogs{}

type NetworkaccessLogs struct {
	// An aggregated log entry that contains comprehensive information about network traffic events.
	Connections *[]NetworkaccessConnection `json:"connections,omitempty"`

	// A collection of remote network health events.
	RemoteNetworks *[]NetworkaccessRemoteNetworkHealthEvent `json:"remoteNetworks,omitempty"`

	// A network access traffic log entry that contains comprehensive information about network traffic events.
	Traffic *[]NetworkaccessNetworkAccessTraffic `json:"traffic,omitempty"`

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

func (s NetworkaccessLogs) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessLogs{}

func (s NetworkaccessLogs) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessLogs
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessLogs: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessLogs: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.logs"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessLogs: %+v", err)
	}

	return encoded, nil
}
