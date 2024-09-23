package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermStoreRelation{}

type TermStoreRelation struct {
	// The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the
	// relation is directly with the [set].
	FromTerm *TermStoreTerm `json:"fromTerm,omitempty"`

	// The type of relation. Possible values are: pin, reuse.
	Relationship *TermStoreRelationType `json:"relationship,omitempty"`

	// The [set] in which the relation is relevant.
	Set *TermStoreSet `json:"set,omitempty"`

	// The to [term] of the relation. The term to which the relationship is defined.
	ToTerm *TermStoreTerm `json:"toTerm,omitempty"`

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

func (s TermStoreRelation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermStoreRelation{}

func (s TermStoreRelation) MarshalJSON() ([]byte, error) {
	type wrapper TermStoreRelation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermStoreRelation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermStoreRelation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termStore.relation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermStoreRelation: %+v", err)
	}

	return encoded, nil
}
