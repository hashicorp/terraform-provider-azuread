package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesCatalogEntry interface {
	Entity
	WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl
}

var _ WindowsUpdatesCatalogEntry = BaseWindowsUpdatesCatalogEntryImpl{}

type BaseWindowsUpdatesCatalogEntryImpl struct {
	// The date on which the content is no longer available to deploy. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Read-only.
	DeployableUntilDateTime nullable.Type[string] `json:"deployableUntilDateTime,omitempty"`

	// The display name of the content. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The release date for the content. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
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

func (s BaseWindowsUpdatesCatalogEntryImpl) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return s
}

func (s BaseWindowsUpdatesCatalogEntryImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesCatalogEntry = RawWindowsUpdatesCatalogEntryImpl{}

// RawWindowsUpdatesCatalogEntryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesCatalogEntryImpl struct {
	windowsUpdatesCatalogEntry BaseWindowsUpdatesCatalogEntryImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawWindowsUpdatesCatalogEntryImpl) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return s.windowsUpdatesCatalogEntry
}

func (s RawWindowsUpdatesCatalogEntryImpl) Entity() BaseEntityImpl {
	return s.windowsUpdatesCatalogEntry.Entity()
}

var _ json.Marshaler = BaseWindowsUpdatesCatalogEntryImpl{}

func (s BaseWindowsUpdatesCatalogEntryImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesCatalogEntryImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesCatalogEntryImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesCatalogEntryImpl: %+v", err)
	}

	delete(decoded, "deployableUntilDateTime")
	delete(decoded, "displayName")
	delete(decoded, "releaseDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.catalogEntry"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesCatalogEntryImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesCatalogEntryImplementation(input []byte) (WindowsUpdatesCatalogEntry, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesCatalogEntry into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.softwareUpdateCatalogEntry") {
		var out WindowsUpdatesSoftwareUpdateCatalogEntry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesSoftwareUpdateCatalogEntry: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesCatalogEntryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesCatalogEntryImpl: %+v", err)
	}

	return RawWindowsUpdatesCatalogEntryImpl{
		windowsUpdatesCatalogEntry: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
