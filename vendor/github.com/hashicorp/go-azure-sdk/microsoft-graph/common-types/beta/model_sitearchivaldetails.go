package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteArchivalDetails struct {
	// Represents the current archive status of the site collection. Returned only on $select.
	ArchiveStatus *SiteArchiveStatus `json:"archiveStatus,omitempty"`

	ArchivedBy       IdentitySet           `json:"archivedBy"`
	ArchivedDateTime nullable.Type[string] `json:"archivedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &SiteArchivalDetails{}

func (s *SiteArchivalDetails) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ArchiveStatus    *SiteArchiveStatus    `json:"archiveStatus,omitempty"`
		ArchivedDateTime nullable.Type[string] `json:"archivedDateTime,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ArchiveStatus = decoded.ArchiveStatus
	s.ArchivedDateTime = decoded.ArchivedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SiteArchivalDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["archivedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ArchivedBy' for 'SiteArchivalDetails': %+v", err)
		}
		s.ArchivedBy = impl
	}

	return nil
}
