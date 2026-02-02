package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingIreland struct {
	// Movies rating labels in Ireland
	MovieRating *RatingIrelandMoviesType `json:"movieRating,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// TV content rating labels in Ireland
	TvRating *RatingIrelandTelevisionType `json:"tvRating,omitempty"`
}
