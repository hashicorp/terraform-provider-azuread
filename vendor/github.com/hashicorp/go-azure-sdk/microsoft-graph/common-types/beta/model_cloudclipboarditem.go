package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudClipboardItem{}

type CloudClipboardItem struct {
	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Set by the server. DateTime in UTC when the object expires and after that the object is no longer available. The
	// default and also maximum TTL is 12 hours after the creation, but it might change for performance optimization.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Set by the server if not provided in the client's request. DateTime in UTC when the object was modified by the
	// client.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// A cloudClipboardItem can have multiple cloudClipboardItemPayload objects in the payloads. A window can place more
	// than one clipboard object on the clipboard. Each one represents the same information in a different clipboard format.
	Payloads *[]CloudClipboardItemPayload `json:"payloads,omitempty"`

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

func (s CloudClipboardItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudClipboardItem{}

func (s CloudClipboardItem) MarshalJSON() ([]byte, error) {
	type wrapper CloudClipboardItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudClipboardItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudClipboardItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudClipboardItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudClipboardItem: %+v", err)
	}

	return encoded, nil
}
