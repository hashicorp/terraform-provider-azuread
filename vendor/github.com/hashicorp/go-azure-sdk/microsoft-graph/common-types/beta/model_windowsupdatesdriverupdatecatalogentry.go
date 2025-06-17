package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesSoftwareUpdateCatalogEntry = WindowsUpdatesDriverUpdateCatalogEntry{}

type WindowsUpdatesDriverUpdateCatalogEntry struct {
	// The description of the content.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The classification of the driver.
	DriverClass nullable.Type[string] `json:"driverClass,omitempty"`

	// The manufacturer of the driver.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The provider of the driver.
	Provider nullable.Type[string] `json:"provider,omitempty"`

	// The setup information file of the driver.
	SetupInformationFile nullable.Type[string] `json:"setupInformationFile,omitempty"`

	// The unique version of the content.
	Version nullable.Type[string] `json:"version,omitempty"`

	// The date and time when a new version of content was created. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	VersionDateTime nullable.Type[string] `json:"versionDateTime,omitempty"`

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

func (s WindowsUpdatesDriverUpdateCatalogEntry) WindowsUpdatesSoftwareUpdateCatalogEntry() BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl {
	return BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s WindowsUpdatesDriverUpdateCatalogEntry) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return BaseWindowsUpdatesCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s WindowsUpdatesDriverUpdateCatalogEntry) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesDriverUpdateCatalogEntry{}

func (s WindowsUpdatesDriverUpdateCatalogEntry) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDriverUpdateCatalogEntry
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDriverUpdateCatalogEntry: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDriverUpdateCatalogEntry: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.driverUpdateCatalogEntry"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDriverUpdateCatalogEntry: %+v", err)
	}

	return encoded, nil
}
