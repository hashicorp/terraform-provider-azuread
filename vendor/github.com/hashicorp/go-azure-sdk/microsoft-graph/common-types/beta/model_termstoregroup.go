package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermStoreGroup{}

type TermStoreGroup struct {
	// Date and time of the group creation. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description that gives details on the term usage.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of the group.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// ID of the parent site of this group.
	ParentSiteId nullable.Type[string] `json:"parentSiteId,omitempty"`

	// Returns the type of the group. Possible values are: global, system, and siteCollection.
	Scope *TermStoreTermGroupScope `json:"scope,omitempty"`

	// All sets under the group in a term [store].
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

func (s TermStoreGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermStoreGroup{}

func (s TermStoreGroup) MarshalJSON() ([]byte, error) {
	type wrapper TermStoreGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermStoreGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermStoreGroup: %+v", err)
	}

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termStore.group"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermStoreGroup: %+v", err)
	}

	return encoded, nil
}
