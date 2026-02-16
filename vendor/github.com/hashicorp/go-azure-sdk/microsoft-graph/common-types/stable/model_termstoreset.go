package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermStoreSet{}

type TermStoreSet struct {
	// Children terms of set in term [store].
	Children *[]TermStoreTerm `json:"children,omitempty"`

	// Date and time of set creation. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description that gives details on the term usage.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the set for each languageTag.
	LocalizedNames *[]TermStoreLocalizedName `json:"localizedNames,omitempty"`

	ParentGroup *TermStoreGroup `json:"parentGroup,omitempty"`

	// Custom properties for the set.
	Properties *[]KeyValue `json:"properties,omitempty"`

	// Indicates which terms have been pinned or reused directly under the set.
	Relations *[]TermStoreRelation `json:"relations,omitempty"`

	// All the terms under the set.
	Terms *[]TermStoreTerm `json:"terms,omitempty"`

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

func (s TermStoreSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermStoreSet{}

func (s TermStoreSet) MarshalJSON() ([]byte, error) {
	type wrapper TermStoreSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermStoreSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermStoreSet: %+v", err)
	}

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termStore.set"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermStoreSet: %+v", err)
	}

	return encoded, nil
}
