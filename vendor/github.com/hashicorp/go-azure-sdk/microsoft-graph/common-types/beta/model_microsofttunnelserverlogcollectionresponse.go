package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MicrosoftTunnelServerLogCollectionResponse{}

type MicrosoftTunnelServerLogCollectionResponse struct {
	// The end time of the logs collected
	EndDateTime *string `json:"endDateTime,omitempty"`

	// The time when the log collection is expired
	ExpiryDateTime *string `json:"expiryDateTime,omitempty"`

	// The time when the log collection was requested
	RequestDateTime *string `json:"requestDateTime,omitempty"`

	// ID of the server the log collection is requested upon
	ServerId nullable.Type[string] `json:"serverId,omitempty"`

	// The size of the logs in bytes
	SizeInBytes *int64 `json:"sizeInBytes,omitempty"`

	// The start time of the logs collected
	StartDateTime *string `json:"startDateTime,omitempty"`

	// Enum type that represent the status of log collection
	Status *MicrosoftTunnelLogCollectionStatus `json:"status,omitempty"`

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

func (s MicrosoftTunnelServerLogCollectionResponse) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftTunnelServerLogCollectionResponse{}

func (s MicrosoftTunnelServerLogCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftTunnelServerLogCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftTunnelServerLogCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftTunnelServerLogCollectionResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftTunnelServerLogCollectionResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftTunnelServerLogCollectionResponse: %+v", err)
	}

	return encoded, nil
}
