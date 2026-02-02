package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SearchSearchAnswer = SearchAcronym{}

type SearchAcronym struct {
	// What the acronym stands for.
	StandsFor nullable.Type[string] `json:"standsFor,omitempty"`

	State *SearchAnswerState `json:"state,omitempty"`

	// Fields inherited from SearchSearchAnswer

	// The search answer description that is shown on the search results page.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The search answer name that is displayed in search results.
	DisplayName *string `json:"displayName,omitempty"`

	// Details of the user who created or last modified the search answer. Read-only.
	LastModifiedBy *SearchIdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time when the search answer was created or last edited. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The URL link for the search answer. When users select this search answer from the search results, they are directed
	// to the specified URL.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s SearchAcronym) SearchSearchAnswer() BaseSearchSearchAnswerImpl {
	return BaseSearchSearchAnswerImpl{
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		WebUrl:               s.WebUrl,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SearchAcronym) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SearchAcronym{}

func (s SearchAcronym) MarshalJSON() ([]byte, error) {
	type wrapper SearchAcronym
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SearchAcronym: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SearchAcronym: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.search.acronym"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SearchAcronym: %+v", err)
	}

	return encoded, nil
}
