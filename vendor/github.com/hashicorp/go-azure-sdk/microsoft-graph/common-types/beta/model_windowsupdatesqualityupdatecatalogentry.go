package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesSoftwareUpdateCatalogEntry = WindowsUpdatesQualityUpdateCatalogEntry{}

type WindowsUpdatesQualityUpdateCatalogEntry struct {
	// The catalog name of the content. Read-only.
	CatalogName nullable.Type[string] `json:"catalogName,omitempty"`

	// Severity information of the Common Vulnerabilities and Exposures associated with the content.
	CveSeverityInformation *WindowsUpdatesQualityUpdateCveSeverityInformation `json:"cveSeverityInformation,omitempty"`

	// Indicates whether the content can be deployed as an expedited quality update. Read-only.
	IsExpeditable *bool `json:"isExpeditable,omitempty"`

	// The operating system product revisions that are released as part of this quality update.
	ProductRevisions *[]WindowsUpdatesProductRevision `json:"productRevisions,omitempty"`

	// The publishing cadence of the quality update. Possible values are: monthly, outOfBand, unknownFutureValue. Read-only.
	QualityUpdateCadence *WindowsUpdatesQualityUpdateCadence `json:"qualityUpdateCadence,omitempty"`

	QualityUpdateClassification *WindowsUpdatesQualityUpdateClassification `json:"qualityUpdateClassification,omitempty"`

	// The short name of the content. Read-only.
	ShortName nullable.Type[string] `json:"shortName,omitempty"`

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

func (s WindowsUpdatesQualityUpdateCatalogEntry) WindowsUpdatesSoftwareUpdateCatalogEntry() BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl {
	return BaseWindowsUpdatesSoftwareUpdateCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s WindowsUpdatesQualityUpdateCatalogEntry) WindowsUpdatesCatalogEntry() BaseWindowsUpdatesCatalogEntryImpl {
	return BaseWindowsUpdatesCatalogEntryImpl{
		DeployableUntilDateTime: s.DeployableUntilDateTime,
		DisplayName:             s.DisplayName,
		ReleaseDateTime:         s.ReleaseDateTime,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s WindowsUpdatesQualityUpdateCatalogEntry) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesQualityUpdateCatalogEntry{}

func (s WindowsUpdatesQualityUpdateCatalogEntry) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesQualityUpdateCatalogEntry
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesQualityUpdateCatalogEntry: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesQualityUpdateCatalogEntry: %+v", err)
	}

	delete(decoded, "catalogName")
	delete(decoded, "isExpeditable")
	delete(decoded, "qualityUpdateCadence")
	delete(decoded, "shortName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntry"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesQualityUpdateCatalogEntry: %+v", err)
	}

	return encoded, nil
}
