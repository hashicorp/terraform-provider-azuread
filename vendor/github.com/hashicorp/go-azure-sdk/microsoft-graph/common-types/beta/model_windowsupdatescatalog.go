package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesCatalog{}

type WindowsUpdatesCatalog struct {
	// Lists the content that you can approve for deployment. Read-only.
	Entries *[]WindowsUpdatesCatalogEntry `json:"entries,omitempty"`

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

func (s WindowsUpdatesCatalog) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesCatalog{}

func (s WindowsUpdatesCatalog) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesCatalog
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesCatalog: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesCatalog: %+v", err)
	}

	delete(decoded, "entries")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.catalog"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesCatalog: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesCatalog{}

func (s *WindowsUpdatesCatalog) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesCatalog into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["entries"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Entries into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesCatalogEntry, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesCatalogEntryImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Entries' for 'WindowsUpdatesCatalog': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Entries = &output
	}

	return nil
}
