package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesDeployableContent = WindowsUpdatesCatalogContent{}

type WindowsUpdatesCatalogContent struct {
	CatalogEntry *WindowsUpdatesCatalogEntry `json:"catalogEntry,omitempty"`

	// Fields inherited from WindowsUpdatesDeployableContent

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesCatalogContent) WindowsUpdatesDeployableContent() BaseWindowsUpdatesDeployableContentImpl {
	return BaseWindowsUpdatesDeployableContentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesCatalogContent{}

func (s WindowsUpdatesCatalogContent) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesCatalogContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesCatalogContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesCatalogContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.catalogContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesCatalogContent: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesCatalogContent{}

func (s *WindowsUpdatesCatalogContent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesCatalogContent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["catalogEntry"]; ok {
		impl, err := UnmarshalWindowsUpdatesCatalogEntryImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CatalogEntry' for 'WindowsUpdatesCatalogContent': %+v", err)
		}
		s.CatalogEntry = &impl
	}

	return nil
}
