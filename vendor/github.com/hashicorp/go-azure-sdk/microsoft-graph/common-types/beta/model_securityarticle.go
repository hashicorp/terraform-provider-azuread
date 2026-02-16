package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityArticle{}

type SecurityArticle struct {
	Body *SecurityFormattedContent `json:"body,omitempty"`

	// The date and time when this article was created. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// URL of the header image for this article, used for display purposes.
	ImageUrl nullable.Type[string] `json:"imageUrl,omitempty"`

	// Indicators related to this article.
	Indicators *[]SecurityArticleIndicator `json:"indicators,omitempty"`

	// Indicates whether this article is currently featured by Microsoft.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The most recent date and time when this article was updated. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	Summary *SecurityFormattedContent `json:"summary,omitempty"`

	// Tags for this article, communicating keywords, or key concepts.
	Tags *[]string `json:"tags,omitempty"`

	// The title of this article.
	Title *string `json:"title,omitempty"`

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

func (s SecurityArticle) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityArticle{}

func (s SecurityArticle) MarshalJSON() ([]byte, error) {
	type wrapper SecurityArticle
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityArticle: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityArticle: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.article"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityArticle: %+v", err)
	}

	return encoded, nil
}
