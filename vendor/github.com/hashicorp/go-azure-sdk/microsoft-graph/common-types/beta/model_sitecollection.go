package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteCollection struct {
	// Represents whether the site collection is recently archived, fully archived, or reactivating. Possible values are:
	// recentlyArchived, fullyArchived, reactivating, unknownFutureValue.
	ArchivalDetails *SiteArchivalDetails `json:"archivalDetails,omitempty"`

	// The geographic region code for where this site collection resides. Only present for multi-geo tenants. Read-only.
	DataLocationCode nullable.Type[string] `json:"dataLocationCode,omitempty"`

	// The hostname for the site collection. Read-only.
	Hostname nullable.Type[string] `json:"hostname,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If present, indicates that this is a root site collection in SharePoint. Read-only.
	Root *Root `json:"root,omitempty"`
}

var _ json.Marshaler = SiteCollection{}

func (s SiteCollection) MarshalJSON() ([]byte, error) {
	type wrapper SiteCollection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SiteCollection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SiteCollection: %+v", err)
	}

	delete(decoded, "dataLocationCode")
	delete(decoded, "hostname")
	delete(decoded, "root")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SiteCollection: %+v", err)
	}

	return encoded, nil
}
