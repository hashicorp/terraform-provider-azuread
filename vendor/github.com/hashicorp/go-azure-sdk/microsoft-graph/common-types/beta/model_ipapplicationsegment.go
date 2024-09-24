package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplicationSegment = IPApplicationSegment{}

type IPApplicationSegment struct {
	Application     *Application                   `json:"application,omitempty"`
	DestinationHost nullable.Type[string]          `json:"destinationHost,omitempty"`
	DestinationType *PrivateNetworkDestinationType `json:"destinationType,omitempty"`
	Port            nullable.Type[int64]           `json:"port,omitempty"`
	Ports           *[]string                      `json:"ports,omitempty"`
	Protocol        *PrivateNetworkProtocol        `json:"protocol,omitempty"`

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

func (s IPApplicationSegment) ApplicationSegment() BaseApplicationSegmentImpl {
	return BaseApplicationSegmentImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s IPApplicationSegment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IPApplicationSegment{}

func (s IPApplicationSegment) MarshalJSON() ([]byte, error) {
	type wrapper IPApplicationSegment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPApplicationSegment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPApplicationSegment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ipApplicationSegment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPApplicationSegment: %+v", err)
	}

	return encoded, nil
}
