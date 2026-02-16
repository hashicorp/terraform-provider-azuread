package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamFunSettings struct {
	// If set to true, enables users to include custom memes.
	AllowCustomMemes nullable.Type[bool] `json:"allowCustomMemes,omitempty"`

	// If set to true, enables Giphy use.
	AllowGiphy nullable.Type[bool] `json:"allowGiphy,omitempty"`

	// If set to true, enables users to include stickers and memes.
	AllowStickersAndMemes nullable.Type[bool] `json:"allowStickersAndMemes,omitempty"`

	// Giphy content rating. Possible values are: moderate, strict.
	GiphyContentRating *GiphyRatingType `json:"giphyContentRating,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
