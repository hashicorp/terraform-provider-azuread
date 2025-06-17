package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateProductKnowledgeBaseArticle struct {
	// The unique identifier for the knowledge base article. Read-only
	ArticleId *string `json:"articleId,omitempty"`

	// The URL of the knowledge base article. Read-only
	ArticleUrl *string `json:"articleUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
