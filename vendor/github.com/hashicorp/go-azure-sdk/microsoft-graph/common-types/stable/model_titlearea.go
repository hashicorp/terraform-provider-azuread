package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleArea struct {
	// Alternative text on the title area.
	AlternativeText nullable.Type[string] `json:"alternativeText,omitempty"`

	// Indicates whether the title area has a gradient effect enabled.
	EnableGradientEffect nullable.Type[bool] `json:"enableGradientEffect,omitempty"`

	// URL of the image in the title area.
	ImageWebUrl nullable.Type[string] `json:"imageWebUrl,omitempty"`

	// Enumeration value that indicates the layout of the title area. The possible values are: imageAndTitle, plain,
	// colorBlock, overlap, unknownFutureValue.
	Layout *TitleAreaLayoutType `json:"layout,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains collections of data that can be processed by server side services like search index and link fixup.
	ServerProcessedContent *ServerProcessedContent `json:"serverProcessedContent,omitempty"`

	// Indicates whether the author should be shown in title area.
	ShowAuthor nullable.Type[bool] `json:"showAuthor,omitempty"`

	// Indicates whether the published date should be shown in title area.
	ShowPublishedDate nullable.Type[bool] `json:"showPublishedDate,omitempty"`

	// Indicates whether the text block above title should be shown in title area.
	ShowTextBlockAboveTitle nullable.Type[bool] `json:"showTextBlockAboveTitle,omitempty"`

	// The text above title line.
	TextAboveTitle nullable.Type[string] `json:"textAboveTitle,omitempty"`

	// Enumeration value that indicates the text alignment of the title area. The possible values are: left, center,
	// unknownFutureValue.
	TextAlignment *TitleAreaTextAlignmentType `json:"textAlignment,omitempty"`
}
