package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSoftwareUpdateCatalogEntry interface {
	Entity
	WindowsUpdatesCatalogEntry
	WindowsUpdatesSoftwareUpdateCatalogEntry() BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl
}

var _ WindowsUpdatesSoftwareUpdateCatalogEntry = BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl{}

type BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl struct {

	// Fields inherited from WindowsUpdatesCatalogEntry

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

func (s BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl) WindowsUpdatesSoftwareUpdateCatalogEntry() BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl {
	return s
}

func (s BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return BaseWindowsUpdatesCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesSoftwareUpdateCatalogEntry = RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl{}

// RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl struct {
	windowsUpdatesSoftwareUpdateCatalogEntry BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl
	Type                                     string
	Values                                   map[string]interface{}
}

func (s RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl) WindowsUpdatesSoftwareUpdateCatalogEntry() BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl {
	return s.windowsUpdatesSoftwareUpdateCatalogEntry
}

func (s RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return s.windowsUpdatesSoftwareUpdateCatalogEntry.WindowsUpdatesCatalogEntry()
}

func (s RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl) Entity() BaseEntityImpl {
	return s.windowsUpdatesSoftwareUpdateCatalogEntry.Entity()
}

var _ json.Marshaler = BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl{}

func (s BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.softwareUpdateCatalogEntry"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesSoftwareUpdateCatalogEntryImplementation(input []byte) (WindowsUpdatesSoftwareUpdateCatalogEntry, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesSoftwareUpdateCatalogEntry into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.driverUpdateCatalogEntry") {
		var out WindowsUpdatesDriverUpdateCatalogEntry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesDriverUpdateCatalogEntry: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.featureUpdateCatalogEntry") {
		var out WindowsUpdatesFeatureUpdateCatalogEntry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesFeatureUpdateCatalogEntry: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntry") {
		var out WindowsUpdatesQualityUpdateCatalogEntry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesQualityUpdateCatalogEntry: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl: %+v", err)
	}

	return RawWindowsUpdatesSoftwareUpdateCatalogEntryImpl{
		windowsUpdatesSoftwareUpdateCatalogEntry: parent,
		Type:                                     value,
		Values:                                   temp,
	}, nil

}
