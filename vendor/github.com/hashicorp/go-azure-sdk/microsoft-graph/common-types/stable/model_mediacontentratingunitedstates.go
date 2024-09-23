package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedStates struct {
	// Movies rating labels in United States
	MovieRating *RatingUnitedStatesMoviesType `json:"movieRating,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// TV content rating labels in United States
	TvRating *RatingUnitedStatesTelevisionType `json:"tvRating,omitempty"`
}
