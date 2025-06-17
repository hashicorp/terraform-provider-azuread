package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SearchSearchAnswer = SearchBookmark{}

type SearchBookmark struct {
	// Date and time when the bookmark stops appearing as a search result. Set as null for always available. The timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z.
	AvailabilityEndDateTime nullable.Type[string] `json:"availabilityEndDateTime,omitempty"`

	// Date and time when the bookmark starts to appear as a search result. Set as null for always available. The timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z.
	AvailabilityStartDateTime nullable.Type[string] `json:"availabilityStartDateTime,omitempty"`

	// Categories commonly used to describe this bookmark. For example, IT and HR.
	Categories *[]string `json:"categories,omitempty"`

	// The list of security groups that are able to view this bookmark.
	GroupIds *[]string `json:"groupIds,omitempty"`

	// True if this bookmark was suggested to the admin, by a user, or was mined and suggested by Microsoft. Read-only.
	IsSuggested nullable.Type[bool] `json:"isSuggested,omitempty"`

	// Keywords that trigger this bookmark to appear in search results.
	Keywords *SearchAnswerKeyword `json:"keywords,omitempty"`

	// A list of geographically specific language names in which this bookmark can be viewed. Each language tag value
	// follows the pattern {language}-{region}. For example, en-us is English as used in the United States. For the list of
	// possible values, see Supported language tags.
	LanguageTags *[]string `json:"languageTags,omitempty"`

	// List of devices and operating systems that are able to view this bookmark. Possible values are: android,
	// androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, unknown,
	// androidASOP, androidMobileApplicationManagement, iOSMobileApplicationManagement, unknownFutureValue.
	Platforms *[]DevicePlatformType `json:"platforms,omitempty"`

	// List of Power Apps associated with this bookmark. If users add existing Power Apps to a bookmark, they can complete
	// tasks directly on the search results page, such as entering vacation time or reporting expenses.
	PowerAppIds *[]string `json:"powerAppIds,omitempty"`

	State *SearchAnswerState `json:"state,omitempty"`

	// Variations of a bookmark for different countries/regions or devices. Use when you need to show different content to
	// users based on their device, country/region, or both. The date and group settings apply to all variations.
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

func (s SearchBookmark) SearchSearchAnswer() BaseSearchSearchAnswerImpl {
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

func (s SearchBookmark) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SearchBookmark{}

func (s SearchBookmark) MarshalJSON() ([]byte, error) {
	type wrapper SearchBookmark
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SearchBookmark: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SearchBookmark: %+v", err)
	}

	delete(decoded, "isSuggested")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.search.bookmark"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SearchBookmark: %+v", err)
	}

	return encoded, nil
}
