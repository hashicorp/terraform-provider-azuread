package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessRemoteNetworkHealthEvent{}

type NetworkaccessRemoteNetworkHealthEvent struct {
	// The number of BGP routes advertised through tunnel.
	BgpRoutesAdvertisedCount nullable.Type[int64] `json:"bgpRoutesAdvertisedCount,omitempty"`

	// The time of the original event generation in UTC. Supports $filter (ge, le) and $orderby.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the event.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The IP address of the destination.
	DestinationIp nullable.Type[string] `json:"destinationIp,omitempty"`

	// The number of bytes sent from the destination to the source.
	ReceivedBytes nullable.Type[int64] `json:"receivedBytes,omitempty"`

	// A unique identifier for each remoteNetwork site. Supports $filter (eq).
	RemoteNetworkId *string `json:"remoteNetworkId,omitempty"`

	// The number of bytes sent from the source to the destination for the connection or session.
	SentBytes nullable.Type[int64] `json:"sentBytes,omitempty"`

	// The public IP address.
	SourceIp nullable.Type[string] `json:"sourceIp,omitempty"`

	Status *NetworkaccessRemoteNetworkStatus `json:"status,omitempty"`

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

func (s NetworkaccessRemoteNetworkHealthEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessRemoteNetworkHealthEvent{}

func (s NetworkaccessRemoteNetworkHealthEvent) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessRemoteNetworkHealthEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessRemoteNetworkHealthEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRemoteNetworkHealthEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.remoteNetworkHealthEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessRemoteNetworkHealthEvent: %+v", err)
	}

	return encoded, nil
}
