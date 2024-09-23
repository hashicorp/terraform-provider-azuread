package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfiguration struct {
	Collate         nullable.Type[bool]     `json:"collate,omitempty"`
	ColorMode       *PrintColorMode         `json:"colorMode,omitempty"`
	Copies          nullable.Type[int64]    `json:"copies,omitempty"`
	Dpi             nullable.Type[int64]    `json:"dpi,omitempty"`
	DuplexMode      *PrintDuplexMode        `json:"duplexMode,omitempty"`
	FeedDirection   *PrinterFeedDirection   `json:"feedDirection,omitempty"`
	FeedOrientation *PrinterFeedOrientation `json:"feedOrientation,omitempty"`
	Finishings      *[]PrintFinishing       `json:"finishings,omitempty"`
	FitPdfToPage    nullable.Type[bool]     `json:"fitPdfToPage,omitempty"`
	InputBin        nullable.Type[string]   `json:"inputBin,omitempty"`
	Margin          *PrintMargin            `json:"margin,omitempty"`
	MediaSize       nullable.Type[string]   `json:"mediaSize,omitempty"`
	MediaType       nullable.Type[string]   `json:"mediaType,omitempty"`
	MultipageLayout *PrintMultipageLayout   `json:"multipageLayout,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Orientation   *PrintOrientation     `json:"orientation,omitempty"`
	OutputBin     nullable.Type[string] `json:"outputBin,omitempty"`
	PageRanges    *[]IntegerRange       `json:"pageRanges,omitempty"`
	PagesPerSheet nullable.Type[int64]  `json:"pagesPerSheet,omitempty"`
	Quality       *PrintQuality         `json:"quality,omitempty"`
	Scaling       *PrintScaling         `json:"scaling,omitempty"`
}
