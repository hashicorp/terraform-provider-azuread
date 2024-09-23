package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryRedundancyDetectionSettings struct {
	// Indicates whether email threading and near duplicate detection are enabled.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Specifies the maximum number of words used for email threading and near duplicate detection. To learn more, see
	// Minimum/maximum number of words.
	MaxWords nullable.Type[int64] `json:"maxWords,omitempty"`

	// Specifies the minimum number of words used for email threading and near duplicate detection. To learn more, see
	// Minimum/maximum number of words.
	MinWords nullable.Type[int64] `json:"minWords,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the similarity level for documents to be put in the same near duplicate set. To learn more, see Document
	// and email similarity threshold.
	SimilarityThreshold nullable.Type[int64] `json:"similarityThreshold,omitempty"`
}
