package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Office365GroupsActivityStorage{}

type Office365GroupsActivityStorage struct {
	// The storage used in group mailbox.
	MailboxStorageUsedInBytes nullable.Type[int64] `json:"mailboxStorageUsedInBytes,omitempty"`

	// The snapshot date for Exchange and SharePoint used storage.
	ReportDate nullable.Type[string] `json:"reportDate,omitempty"`

	// The number of days the report covers.
	ReportPeriod nullable.Type[string] `json:"reportPeriod,omitempty"`

	// The latest date of the content.
	ReportRefreshDate nullable.Type[string] `json:"reportRefreshDate,omitempty"`

	// The storage used in SharePoint document library.
	SiteStorageUsedInBytes nullable.Type[int64] `json:"siteStorageUsedInBytes,omitempty"`

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

func (s Office365GroupsActivityStorage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Office365GroupsActivityStorage{}

func (s Office365GroupsActivityStorage) MarshalJSON() ([]byte, error) {
	type wrapper Office365GroupsActivityStorage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Office365GroupsActivityStorage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Office365GroupsActivityStorage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.office365GroupsActivityStorage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Office365GroupsActivityStorage: %+v", err)
	}

	return encoded, nil
}
