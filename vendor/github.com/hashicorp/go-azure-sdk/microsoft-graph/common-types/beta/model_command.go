package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Command{}

type Command struct {
	AppServiceName    nullable.Type[string] `json:"appServiceName,omitempty"`
	Error             nullable.Type[string] `json:"error,omitempty"`
	PackageFamilyName nullable.Type[string] `json:"packageFamilyName,omitempty"`
	Payload           *PayloadRequest       `json:"payload,omitempty"`
	PermissionTicket  nullable.Type[string] `json:"permissionTicket,omitempty"`
	PostBackUri       nullable.Type[string] `json:"postBackUri,omitempty"`
	Responsepayload   *PayloadResponse      `json:"responsepayload,omitempty"`
	Status            nullable.Type[string] `json:"status,omitempty"`
	Type              nullable.Type[string] `json:"type,omitempty"`

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

func (s Command) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Command{}

func (s Command) MarshalJSON() ([]byte, error) {
	type wrapper Command
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Command: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Command: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.command"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Command: %+v", err)
	}

	return encoded, nil
}
