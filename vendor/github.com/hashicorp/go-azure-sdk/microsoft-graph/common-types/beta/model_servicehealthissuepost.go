package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssuePost struct {
	// The published time of the post.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The content of the service issue post. The supported value for the contentType property is html.
	Description *ItemBody `json:"description,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The post type of the service issue historical post. Possible values are: regular, quick, strategic,
	// unknownFutureValue.
	PostType *PostType `json:"postType,omitempty"`
}
