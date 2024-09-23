package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsSettingInstance{}

type WindowsSettingInstance struct {
	// Set by the server. Represents the dateTime in UTC when the object was created on the server.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Set by the server. The object expires at the specified dateTime in UTC, making it unavailable after that time.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Set by the server if not provided in the request from the Windows client device. Refers to the user's Windows device
	// that modified the object at the specified dateTime in UTC.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Base64-encoded JSON setting value.
	Payload *string `json:"payload,omitempty"`

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

func (s WindowsSettingInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsSettingInstance{}

func (s WindowsSettingInstance) MarshalJSON() ([]byte, error) {
	type wrapper WindowsSettingInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsSettingInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsSettingInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsSettingInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsSettingInstance: %+v", err)
	}

	return encoded, nil
}
