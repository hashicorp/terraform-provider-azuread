package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoReviewSettings struct {
	// Possible values: Approve, Deny, or Recommendation. If Recommendation, then accessRecommendationsEnabled in the
	// accessReviewSettings resource should also be set to true. If you want to have the system provide a decision even if
	// the reviewer does not make a choice, set the autoReviewEnabled property in the accessReviewSettings resource to true
	// and include an autoReviewSettings object with the notReviewedResult property. Then, when a review completes, based on
	// the notReviewedResult property, the decision is recorded as either Approve or Deny.
	NotReviewedResult nullable.Type[string] `json:"notReviewedResult,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
