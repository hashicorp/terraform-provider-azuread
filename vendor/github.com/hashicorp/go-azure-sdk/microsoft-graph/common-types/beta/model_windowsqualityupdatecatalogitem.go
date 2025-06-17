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
	// Windows quality update category
	Classification *WindowsQualityUpdateCategory `json:"classification,omitempty"`

	// When TRUE, indicates that the quality updates qualify for expedition. When FALSE, indicates the quality updates do
	// not quality for expedition. Default value is FALSE. Read-only
	IsExpeditable *bool `json:"isExpeditable,omitempty"`

	// Identifies the knowledge base article associated with the Windows quality update catalog item. Read-only
	KbArticleId *string `json:"kbArticleId,omitempty"`

	// The operating system product revisions that are released as part of this quality update. Read-only.
	ProductRevisions *[]WindowsQualityUpdateCatalogProductRevision `json:"productRevisions,omitempty"`

	// The publishing cadence of the quality update. Possible values are: monthly, outOfBand. This property cannot be
	// modified and is automatically populated when the catalog is created.
	QualityUpdateCadence *WindowsQualityUpdateCadence `json:"qualityUpdateCadence,omitempty"`

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

	delete(decoded, "productRevisions")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsQualityUpdateCatalogItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsQualityUpdateCatalogItem: %+v", err)
	}

	return encoded, nil
}
