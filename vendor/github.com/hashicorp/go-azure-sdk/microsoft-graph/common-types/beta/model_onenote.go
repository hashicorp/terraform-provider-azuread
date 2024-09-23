package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Onenote{}

type Onenote struct {
	// The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
	Notebooks *[]Notebook `json:"notebooks,omitempty"`

	// The status of OneNote operations. Getting an operations collection isn't supported, but you can get the status of
	// long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable.
	Operations *[]OnenoteOperation `json:"operations,omitempty"`

	// The pages in all OneNote notebooks that are owned by the user or group. Read-only. Nullable.
	Pages *[]OnenotePage `json:"pages,omitempty"`

	// The image and other file resources in OneNote pages. Getting a resources collection isn't supported, but you can get
	// the binary content of a specific resource. Read-only. Nullable.
	Resources *[]OnenoteResource `json:"resources,omitempty"`

	// The section groups in all OneNote notebooks that are owned by the user or group. Read-only. Nullable.
	SectionGroups *[]SectionGroup `json:"sectionGroups,omitempty"`

	// The sections in all OneNote notebooks that are owned by the user or group. Read-only. Nullable.
	Sections *[]OnenoteSection `json:"sections,omitempty"`

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

func (s Onenote) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Onenote{}

func (s Onenote) MarshalJSON() ([]byte, error) {
	type wrapper Onenote
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Onenote: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Onenote: %+v", err)
	}

	delete(decoded, "notebooks")
	delete(decoded, "operations")
	delete(decoded, "pages")
	delete(decoded, "resources")
	delete(decoded, "sectionGroups")
	delete(decoded, "sections")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onenote"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Onenote: %+v", err)
	}

	return encoded, nil
}
