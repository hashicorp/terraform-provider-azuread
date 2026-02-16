package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = LearningContent{}

type LearningContent struct {
	// Keywords, topics, and other tags associated with the learning content. Optional.
	AdditionalTags *[]string `json:"additionalTags,omitempty"`

	// The content web URL for the learning content. Required.
	ContentWebUrl string `json:"contentWebUrl"`

	// The authors, creators, or contributors of the learning content. Optional.
	Contributors *[]string `json:"contributors,omitempty"`

	// The date and time when the learning content was created. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Optional.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description or summary for the learning content. Optional.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The duration of the learning content in seconds. The value is represented in ISO 8601 format for durations. Optional.
	Duration nullable.Type[string] `json:"duration,omitempty"`

	// Unique external content ID for the learning content. Required.
	ExternalId string `json:"externalId"`

	// The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
	Format nullable.Type[string] `json:"format,omitempty"`

	// Indicates whether the content is active or not. Inactive content doesn't show up in the UI. The default value is
	// true. Optional.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The
	// default value is false. Optional.
	IsPremium nullable.Type[bool] `json:"isPremium,omitempty"`

	// Indicates whether the learning content is searchable or not. The default value is true. Optional.
	IsSearchable nullable.Type[bool] `json:"isSearchable,omitempty"`

	// The language of the learning content, for example, en-us or fr-fr. Required.
	LanguageTag string `json:"languageTag"`

	// The date and time when the learning content was last modified. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Optional.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The difficulty level of the learning content. Possible values are: Beginner, Intermediate, Advanced,
	// unknownFutureValue. Optional.
	Level *Level `json:"level,omitempty"`

	// The number of pages of the learning content, for example, 9. Optional.
	NumberOfPages nullable.Type[int64] `json:"numberOfPages,omitempty"`

	// The skills tags associated with the learning content. Optional.
	SkillTags *[]string `json:"skillTags,omitempty"`

	// The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
	SourceName nullable.Type[string] `json:"sourceName,omitempty"`

	// The URL of learning content thumbnail image. Optional.
	ThumbnailWebUrl nullable.Type[string] `json:"thumbnailWebUrl,omitempty"`

	// The title of the learning content. Required.
	Title string `json:"title"`

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

func (s LearningContent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LearningContent{}

func (s LearningContent) MarshalJSON() ([]byte, error) {
	type wrapper LearningContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LearningContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LearningContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.learningContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LearningContent: %+v", err)
	}

	return encoded, nil
}
