package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementExportJob{}

type DeviceManagementExportJob struct {
	// Time that the exported report expires.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Filters applied on the report. The maximum length allowed for this property is 2000 characters.
	Filter nullable.Type[string] `json:"filter,omitempty"`

	// Possible values for the file format of a report to be exported.
	Format *DeviceManagementReportFileFormat `json:"format,omitempty"`

	// Configures how the requested export job is localized.
	LocalizationType *DeviceManagementExportJobLocalizationType `json:"localizationType,omitempty"`

	// Name of the report. The maximum length allowed for this property is 2000 characters.
	ReportName *string `json:"reportName,omitempty"`

	// Time that the exported report was requested.
	RequestDateTime *string `json:"requestDateTime,omitempty"`

	// Configures a search term to filter the data. The maximum length allowed for this property is 100 characters.
	Search nullable.Type[string] `json:"search,omitempty"`

	// Columns selected from the report. The maximum number of allowed columns names is 256. The maximum length allowed for
	// each column name in this property is 1000 characters.
	Select *[]string `json:"select,omitempty"`

	// A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or
	// CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied
	// to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a
	// CachedReportConfiguration id. The maximum length allowed for this property is 128 characters.
	SnapshotId nullable.Type[string] `json:"snapshotId,omitempty"`

	// Possible statuses associated with a generated report.
	Status *DeviceManagementReportStatus `json:"status,omitempty"`

	// Temporary location of the exported report.
	Url nullable.Type[string] `json:"url,omitempty"`

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

func (s DeviceManagementExportJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementExportJob{}

func (s DeviceManagementExportJob) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementExportJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementExportJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementExportJob: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementExportJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementExportJob: %+v", err)
	}

	return encoded, nil
}
