package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermStoreStore{}

type TermStoreStore struct {
	// Default language of the term store.
	DefaultLanguageTag *string `json:"defaultLanguageTag,omitempty"`

	// Collection of all groups available in the term store.
	Groups *[]TermStoreGroup `json:"groups,omitempty"`

	// List of languages for the term store.
	LanguageTags *[]string `json:"languageTags,omitempty"`

	// Collection of all sets available in the term store.
	Sets *[]TermStoreSet `json:"sets,omitempty"`

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

func (s TermStoreStore) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermStoreStore{}

func (s TermStoreStore) MarshalJSON() ([]byte, error) {
	type wrapper TermStoreStore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermStoreStore: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermStoreStore: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termStore.store"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermStoreStore: %+v", err)
	}

	return encoded, nil
}
