package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ProfilePhoto{}

type ProfilePhoto struct {
	// The height of the photo. Read-only.
	Height nullable.Type[int64] `json:"height,omitempty"`

	// The width of the photo. Read-only.
	Width nullable.Type[int64] `json:"width,omitempty"`

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

func (s ProfilePhoto) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProfilePhoto{}

func (s ProfilePhoto) MarshalJSON() ([]byte, error) {
	type wrapper ProfilePhoto
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProfilePhoto: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProfilePhoto: %+v", err)
	}

	delete(decoded, "height")
	delete(decoded, "width")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.profilePhoto"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProfilePhoto: %+v", err)
	}

	return encoded, nil
}
