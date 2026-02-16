package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesApplicableContent struct {
	// Catalog entry for the update or content.
	CatalogEntry *WindowsUpdatesCatalogEntry `json:"catalogEntry,omitempty"`

	// ID of the catalog entry for the applicable content.
	CatalogEntryId *string `json:"catalogEntryId,omitempty"`

	// Collection of devices and recommendations for applicable catalog content.
	MatchedDevices *[]WindowsUpdatesApplicableContentDeviceMatch `json:"matchedDevices,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &WindowsUpdatesApplicableContent{}

func (s *WindowsUpdatesApplicableContent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CatalogEntryId *string                                       `json:"catalogEntryId,omitempty"`
		MatchedDevices *[]WindowsUpdatesApplicableContentDeviceMatch `json:"matchedDevices,omitempty"`
		ODataId        *string                                       `json:"@odata.id,omitempty"`
		ODataType      *string                                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CatalogEntryId = decoded.CatalogEntryId
	s.MatchedDevices = decoded.MatchedDevices
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesApplicableContent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["catalogEntry"]; ok {
		impl, err := UnmarshalWindowsUpdatesCatalogEntryImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CatalogEntry' for 'WindowsUpdatesApplicableContent': %+v", err)
		}
		s.CatalogEntry = &impl
	}

	return nil
}
