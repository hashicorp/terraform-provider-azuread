package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365GroupsActivityGroupCounts{}

type Office365GroupsActivityGroupCounts struct {
	// The number of active groups. A group is considered active if any of the following occurred: group mailbox received
	// email, or a user viewed, edited, shared, or synced files in SharePoint document library, or a user viewed SharePoint
	// pages, or a user posted, read, or liked messages in Yammer groups.
	Active nullable.Type[int64] `json:"active,omitempty"`

	// The date on which groups were active.
	ReportDate nullable.Type[string] `json:"reportDate,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The total number of groups.
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

func (s Office365GroupsActivityGroupCounts) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365GroupsActivityGroupCounts{}

func (s Office365GroupsActivityGroupCounts) MarshalJSON() ([]byte, error) {
	type wrapper Office365GroupsActivityGroupCounts
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365GroupsActivityGroupCounts: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365GroupsActivityGroupCounts: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365GroupsActivityGroupCounts"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365GroupsActivityGroupCounts: %+v", err)
	}

	return encoded, nil
}
