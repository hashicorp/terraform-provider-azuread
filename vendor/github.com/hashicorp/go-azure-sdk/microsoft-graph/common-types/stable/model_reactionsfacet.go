package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReactionsFacet struct {
	// Count of comments.
	CommentCount nullable.Type[int64] `json:"commentCount,omitempty"`

	// Count of likes.
	LikeCount nullable.Type[int64] `json:"likeCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of shares.
	ShareCount nullable.Type[int64] `json:"shareCount,omitempty"`
}
