package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdateCatalogItem = WindowsQualityUpdateCatalogItem{}

type WindowsQualityUpdateCatalogItem struct {
	// Windows quality update classification
	Classification *WindowsQualityUpdateClassification `json:"classification,omitempty"`

	// Flag indicating if update qualifies for expedite
	IsExpeditable *bool `json:"isExpeditable,omitempty"`

	// Knowledge base article id
	KbArticleId *string `json:"kbArticleId,omitempty"`

	// Fields inherited from WindowsUpdateCatalogItem

	// The display name for the catalog item.
	DisplayName *string `json:"displayName,omitempty"`

	// The last supported date for a catalog item
	EndOfSupportDate nullable.Type[string] `json:"endOfSupportDate,omitempty"`

	// The date the catalog item was released
	ReleaseDateTime *string `json:"releaseDateTime,omitempty"`

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

func (s WindowsQualityUpdateCatalogItem) WindowsUpdateCatalogItem() BaseWindowsUpdateCatalogItemImpl {
	return BaseWindowsUpdateCatalogItemImpl{
		DisplayName:      s.DisplayName,
		EndOfSupportDate: s.EndOfSupportDate,
		ReleaseDateTime:  s.ReleaseDateTime,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s WindowsQualityUpdateCatalogItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsQualityUpdateCatalogItem{}

func (s WindowsQualityUpdateCatalogItem) MarshalJSON() ([]byte, error) {
	type wrapper WindowsQualityUpdateCatalogItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsQualityUpdateCatalogItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsQualityUpdateCatalogItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsQualityUpdateCatalogItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsQualityUpdateCatalogItem: %+v", err)
	}

	return encoded, nil
}
