package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementCachedReportConfiguration{}

type DeviceManagementCachedReportConfiguration struct {
	// Time that the cached report expires.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Filters applied on report creation.
	Filter nullable.Type[string] `json:"filter,omitempty"`

	// Time that the cached report was last refreshed.
	LastRefreshDateTime *string `json:"lastRefreshDateTime,omitempty"`

	// Caller-managed metadata associated with the report.
	Metadata nullable.Type[string] `json:"metadata,omitempty"`

	// Ordering of columns in the report.
	OrderBy *[]string `json:"orderBy,omitempty"`

	// Name of the report.
	ReportName nullable.Type[string] `json:"reportName,omitempty"`

	// Columns selected from the report.
	Select *[]string `json:"select,omitempty"`

	// Possible statuses associated with a generated report.
	Status *DeviceManagementReportStatus `json:"status,omitempty"`

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

func (s DeviceManagementCachedReportConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementCachedReportConfiguration{}

func (s DeviceManagementCachedReportConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementCachedReportConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementCachedReportConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementCachedReportConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementCachedReportConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementCachedReportConfiguration: %+v", err)
	}

	return encoded, nil
}
