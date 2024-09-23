package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingNewZealand struct {
	// Movies rating labels in New Zealand
	MovieRating *RatingNewZealandMoviesType `json:"movieRating,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// TV content rating labels in New Zealand
	TvRating *RatingNewZealandTelevisionType `json:"tvRating,omitempty"`
}
