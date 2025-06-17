package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedAppStatus = ManagedAppStatusRaw{}

type ManagedAppStatusRaw struct {
	// Status report content.
	Content *Json `json:"content,omitempty"`

	// Fields inherited from ManagedAppStatus

	// Friendly name of the status report.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s ManagedAppStatusRaw) ManagedAppStatus() BaseManagedAppStatusImpl {
	return BaseManagedAppStatusImpl{
		DisplayName: s.DisplayName,
		Version:     s.Version,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s ManagedAppStatusRaw) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedAppStatusRaw{}

func (s ManagedAppStatusRaw) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAppStatusRaw
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAppStatusRaw: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppStatusRaw: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppStatusRaw"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAppStatusRaw: %+v", err)
	}

	return encoded, nil
}
