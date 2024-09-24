package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingJapan struct {
	// Movies rating labels in Japan
	MovieRating *RatingJapanMoviesType `json:"movieRating,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// TV content rating labels in Japan
	TvRating *RatingJapanTelevisionType `json:"tvRating,omitempty"`
}
