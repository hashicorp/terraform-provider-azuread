package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaults struct {
	// The default color mode to use when printing the document. Valid values are described in the following table.
	ColorMode *PrintColorMode `json:"colorMode,omitempty"`

	// The default content (MIME) type to use when processing documents.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// The default number of copies printed per job.
	CopiesPerJob nullable.Type[int64] `json:"copiesPerJob,omitempty"`

	DocumentMimeType nullable.Type[string] `json:"documentMimeType,omitempty"`

	// The default resolution in DPI to use when printing the job.
	Dpi nullable.Type[int64] `json:"dpi,omitempty"`

	DuplexConfiguration *PrintDuplexConfiguration `json:"duplexConfiguration,omitempty"`

	// The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the
	// following table.
	DuplexMode *PrintDuplexMode `json:"duplexMode,omitempty"`

	// The default set of finishings to apply to print jobs. Valid values are described in the following table.
	Finishings *[]PrintFinishing `json:"finishings,omitempty"`

	// The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let
	// the printer decide how to lay out impressions.
	FitPdfToPage nullable.Type[bool] `json:"fitPdfToPage,omitempty"`

	// The default input bin that serves as the paper source.
	InputBin nullable.Type[string] `json:"inputBin,omitempty"`

	// The default media (such as paper) color to print the document on.
	MediaColor nullable.Type[string] `json:"mediaColor,omitempty"`

	// The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in
	// the printerCapabilities topic.
	MediaSize nullable.Type[string] `json:"mediaSize,omitempty"`

	// The default media (such as paper) type to print the document on.
	MediaType nullable.Type[string] `json:"mediaType,omitempty"`

	// The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in
	// the following table.
	MultipageLayout *PrintMultipageLayout `json:"multipageLayout,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The default orientation to use when printing the document. Valid values are described in the following table.
	Orientation *PrintOrientation `json:"orientation,omitempty"`

	// The default output bin to place completed prints into. See the printer's capabilities for a list of supported output
	// bins.
	OutputBin nullable.Type[string] `json:"outputBin,omitempty"`

	// The default number of document pages to print on each sheet.
	PagesPerSheet nullable.Type[int64] `json:"pagesPerSheet,omitempty"`

	PdfFitToPage            nullable.Type[bool]         `json:"pdfFitToPage,omitempty"`
	PresentationDirection   *PrintPresentationDirection `json:"presentationDirection,omitempty"`
	PrintColorConfiguration *PrintColorConfiguration    `json:"printColorConfiguration,omitempty"`
	PrintQuality            *PrintQuality               `json:"printQuality,omitempty"`

	// The default quality to use when printing the document. Valid values are described in the following table.
	Quality *PrintQuality `json:"quality,omitempty"`

	// Specifies how the printer scales the document data to fit the requested media. Valid values are described in the
	// following table.
	Scaling *PrintScaling `json:"scaling,omitempty"`
}
