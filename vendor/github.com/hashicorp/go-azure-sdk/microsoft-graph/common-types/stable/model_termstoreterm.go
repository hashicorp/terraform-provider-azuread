package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermStoreTerm{}

type TermStoreTerm struct {
	// Children of current term.
	Children *[]TermStoreTerm `json:"children,omitempty"`

	// Date and time of term creation. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description about term that is dependent on the languageTag.
	Descriptions *[]TermStoreLocalizedDescription `json:"descriptions,omitempty"`

	// Label metadata for a term.
	Labels *[]TermStoreLocalizedLabel `json:"labels,omitempty"`

	// Last date and time of term modification. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Collection of properties on the term.
	Properties *[]KeyValue `json:"properties,omitempty"`

	// To indicate which terms are related to the current term as either pinned or reused.
	Relations *[]TermStoreRelation `json:"relations,omitempty"`

	// The [set] in which the term is created.
	Set *TermStoreSet `json:"set,omitempty"`

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

func (s TermStoreTerm) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermStoreTerm{}

func (s TermStoreTerm) MarshalJSON() ([]byte, error) {
	type wrapper TermStoreTerm
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermStoreTerm: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermStoreTerm: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termStore.term"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermStoreTerm: %+v", err)
	}

	return encoded, nil
}
