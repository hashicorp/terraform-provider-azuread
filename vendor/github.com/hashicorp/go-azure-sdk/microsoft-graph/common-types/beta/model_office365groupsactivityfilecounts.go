package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365GroupsActivityFileCounts{}

type Office365GroupsActivityFileCounts struct {
	// The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
	Active nullable.Type[int64] `json:"active,omitempty"`

	// The date on which a number of files were active in the group's SharePoint site.
	ReportDate nullable.Type[string] `json:"reportDate,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The total number of files in the group's SharePoint document library.
	Total nullable.Type[int64] `json:"total,omitempty"`

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

func (s Office365GroupsActivityFileCounts) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365GroupsActivityFileCounts{}

func (s Office365GroupsActivityFileCounts) MarshalJSON() ([]byte, error) {
	type wrapper Office365GroupsActivityFileCounts
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365GroupsActivityFileCounts: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365GroupsActivityFileCounts: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365GroupsActivityFileCounts"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365GroupsActivityFileCounts: %+v", err)
	}

	return encoded, nil
}
