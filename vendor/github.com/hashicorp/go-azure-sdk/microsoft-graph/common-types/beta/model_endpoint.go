package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Endpoint{}

type Endpoint struct {
	// Describes the capability that is associated with this resource. (for example, Messages, Conversations, etc.) Not
	// nullable. Read-only.
	Capability *string `json:"capability,omitempty"`

	// Application id of the publishing underlying service. Not nullable. Read-only.
	ProviderId nullable.Type[string] `json:"providerId,omitempty"`

	// Name of the publishing underlying service. Read-only.
	ProviderName nullable.Type[string] `json:"providerName,omitempty"`

	// For Microsoft 365 groups, this is set to a well-known name for the resource (for example, Yammer.FeedURL etc.). Not
	// nullable. Read-only.
	ProviderResourceId nullable.Type[string] `json:"providerResourceId,omitempty"`

	// URL of the published resource. Not nullable. Read-only.
	Uri *string `json:"uri,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s Endpoint) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Endpoint) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Endpoint{}

func (s Endpoint) MarshalJSON() ([]byte, error) {
	type wrapper Endpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Endpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Endpoint: %+v", err)
	}

	delete(decoded, "capability")
	delete(decoded, "providerId")
	delete(decoded, "providerName")
	delete(decoded, "providerResourceId")
	delete(decoded, "uri")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.endpoint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Endpoint: %+v", err)
	}

	return encoded, nil
}
