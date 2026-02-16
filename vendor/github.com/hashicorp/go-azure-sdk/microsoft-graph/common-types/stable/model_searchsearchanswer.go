package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchSearchAnswer interface {
	Entity
	SearchSearchAnswer() BaseSearchSearchAnswerImpl
}

var _ SearchSearchAnswer = BaseSearchSearchAnswerImpl{}

type BaseSearchSearchAnswerImpl struct {
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

func (s BaseSearchSearchAnswerImpl) SearchSearchAnswer() BaseSearchSearchAnswerImpl {
	return s
}

func (s BaseSearchSearchAnswerImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SearchSearchAnswer = RawSearchSearchAnswerImpl{}

// RawSearchSearchAnswerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSearchSearchAnswerImpl struct {
	searchSearchAnswer BaseSearchSearchAnswerImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawSearchSearchAnswerImpl) SearchSearchAnswer() BaseSearchSearchAnswerImpl {
	return s.searchSearchAnswer
}

func (s RawSearchSearchAnswerImpl) Entity() BaseEntityImpl {
	return s.searchSearchAnswer.Entity()
}

var _ json.Marshaler = BaseSearchSearchAnswerImpl{}

func (s BaseSearchSearchAnswerImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSearchSearchAnswerImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSearchSearchAnswerImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSearchSearchAnswerImpl: %+v", err)
	}

	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.search.searchAnswer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSearchSearchAnswerImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSearchSearchAnswerImplementation(input []byte) (SearchSearchAnswer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SearchSearchAnswer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.search.acronym") {
		var out SearchAcronym
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchAcronym: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.bookmark") {
		var out SearchBookmark
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchBookmark: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.search.qna") {
		var out SearchQna
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchQna: %+v", err)
		}
		return out, nil
	}

	var parent BaseSearchSearchAnswerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSearchSearchAnswerImpl: %+v", err)
	}

	return RawSearchSearchAnswerImpl{
		searchSearchAnswer: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
