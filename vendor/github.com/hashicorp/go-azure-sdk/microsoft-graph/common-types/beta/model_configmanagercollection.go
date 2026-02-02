package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ConfigManagerCollection{}

type ConfigManagerCollection struct {
	// The collection identifier in SCCM.
	CollectionIdentifier nullable.Type[string] `json:"collectionIdentifier,omitempty"`

	// The created date.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The DisplayName.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The Hierarchy Identifier.
	HierarchyIdentifier nullable.Type[string] `json:"hierarchyIdentifier,omitempty"`

	// The HierarchyName.
	HierarchyName nullable.Type[string] `json:"hierarchyName,omitempty"`

	// The last modified date.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s ConfigManagerCollection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConfigManagerCollection{}

func (s ConfigManagerCollection) MarshalJSON() ([]byte, error) {
	type wrapper ConfigManagerCollection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConfigManagerCollection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConfigManagerCollection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.configManagerCollection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConfigManagerCollection: %+v", err)
	}

	return encoded, nil
}
