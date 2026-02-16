package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateCatalogItem interface {
	Entity
	WindowsUpdateCatalogItem() BaseWindowsUpdateCatalogItemImpl
}

var _ WindowsUpdateCatalogItem = BaseWindowsUpdateCatalogItemImpl{}

type BaseWindowsUpdateCatalogItemImpl struct {
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

func (s BaseWindowsUpdateCatalogItemImpl) WindowsUpdateCatalogItem() BaseWindowsUpdateCatalogItemImpl {
	return s
}

func (s BaseWindowsUpdateCatalogItemImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdateCatalogItem = RawWindowsUpdateCatalogItemImpl{}

// RawWindowsUpdateCatalogItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdateCatalogItemImpl struct {
	windowsUpdateCatalogItem BaseWindowsUpdateCatalogItemImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawWindowsUpdateCatalogItemImpl) WindowsUpdateCatalogItem() BaseWindowsUpdateCatalogItemImpl {
	return s.windowsUpdateCatalogItem
}

func (s RawWindowsUpdateCatalogItemImpl) Entity() BaseEntityImpl {
	return s.windowsUpdateCatalogItem.Entity()
}

var _ json.Marshaler = BaseWindowsUpdateCatalogItemImpl{}

func (s BaseWindowsUpdateCatalogItemImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdateCatalogItemImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdateCatalogItemImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdateCatalogItemImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdateCatalogItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdateCatalogItemImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdateCatalogItemImplementation(input []byte) (WindowsUpdateCatalogItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdateCatalogItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsFeatureUpdateCatalogItem") {
		var out WindowsFeatureUpdateCatalogItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsFeatureUpdateCatalogItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsQualityUpdateCatalogItem") {
		var out WindowsQualityUpdateCatalogItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsQualityUpdateCatalogItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdateCatalogItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdateCatalogItemImpl: %+v", err)
	}

	return RawWindowsUpdateCatalogItemImpl{
		windowsUpdateCatalogItem: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
