package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesSoftwareUpdateCatalogEntry = WindowsUpdatesFeatureUpdateCatalogEntry{}

type WindowsUpdatesFeatureUpdateCatalogEntry struct {
	// The build number of the feature update. Read-only.
	BuildNumber nullable.Type[string] `json:"buildNumber,omitempty"`

	// The version of the feature update. Read-only.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s WindowsUpdatesFeatureUpdateCatalogEntry) WindowsUpdatesSoftwareUpdateCatalogEntry() BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl {
	return BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s WindowsUpdatesFeatureUpdateCatalogEntry) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return BaseWindowsUpdatesCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s WindowsUpdatesFeatureUpdateCatalogEntry) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesFeatureUpdateCatalogEntry{}

func (s WindowsUpdatesFeatureUpdateCatalogEntry) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesFeatureUpdateCatalogEntry
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesFeatureUpdateCatalogEntry: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesFeatureUpdateCatalogEntry: %+v", err)
	}

	delete(decoded, "buildNumber")
	delete(decoded, "version")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.featureUpdateCatalogEntry"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesFeatureUpdateCatalogEntry: %+v", err)
	}

	return encoded, nil
}
