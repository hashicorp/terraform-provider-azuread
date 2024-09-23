package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SearchSearchAnswer = SearchQna{}

type SearchQna struct {
	// Date and time when the QnA stops appearing as a search result. Set as null for always available. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014 is 2014-01-01T00:00:00Z.
	AvailabilityEndDateTime nullable.Type[string] `json:"availabilityEndDateTime,omitempty"`

	// Date and time when the QnA starts to appear as a search result. Set as null for always available. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014 is 2014-01-01T00:00:00Z.
	AvailabilityStartDateTime nullable.Type[string] `json:"availabilityStartDateTime,omitempty"`

	// The list of security groups that are able to view this QnA.
	GroupIds *[]string `json:"groupIds,omitempty"`

	// True if a user or Microsoft suggested this QnA to the admin. Read-only.
	IsSuggested nullable.Type[bool] `json:"isSuggested,omitempty"`

	// Keywords that trigger this QnA to appear in search results.
	Keywords *SearchAnswerKeyword `json:"keywords,omitempty"`

	// A list of geographically specific language names in which this QnA can be viewed. Each language tag value follows the
	// pattern {language}-{region}. For example, en-us is English as used in the United States. For the list of possible
	// values, see Supported language tags.
	LanguageTags *[]string `json:"languageTags,omitempty"`

	// List of devices and operating systems that are able to view this QnA. Possible values are: android, androidForWork,
	// ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, unknown, androidASOP,
	// androidMobileApplicationManagement, iOSMobileApplicationManagement, unknownFutureValue.
	Platforms *[]DevicePlatformType `json:"platforms,omitempty"`

	State *SearchAnswerState `json:"state,omitempty"`

	// Variations of a QnA for different countries or devices. Use when you need to show different content to users based on
	// their device, country/region, or both. The date and group settings apply to all variations.
	TargetedVariations *[]SearchAnswerVariant `json:"targetedVariations,omitempty"`

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

func (s SearchQna) SearchSearchAnswer() BaseSearchSearchAnswerImpl {
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

func (s SearchQna) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SearchQna{}

func (s SearchQna) MarshalJSON() ([]byte, error) {
	type wrapper SearchQna
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SearchQna: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SearchQna: %+v", err)
	}

	delete(decoded, "isSuggested")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.search.qna"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SearchQna: %+v", err)
	}

	return encoded, nil
}
