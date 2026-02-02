package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilities struct {
	// A list of supported bottom margins(in microns) for the printer.
	BottomMargins *[]int64 `json:"bottomMargins,omitempty"`

	// True if the printer supports collating when printing multiple copies of a multi-page document; false otherwise.
	Collation nullable.Type[bool] `json:"collation,omitempty"`

	// The color modes supported by the printer. Valid values are described in the following table.
	ColorModes *[]PrintColorMode `json:"colorModes,omitempty"`

	// A list of supported content (MIME) types that the printer supports. It isn't guaranteed that the Universal Print
	// service supports printing all of these MIME types.
	ContentTypes *[]string `json:"contentTypes,omitempty"`

	// The range of copies per job supported by the printer.
	CopiesPerJob *IntegerRange `json:"copiesPerJob,omitempty"`

	// The list of print resolutions in DPI that are supported by the printer.
	Dpis *[]int64 `json:"dpis,omitempty"`

	// The list of duplex modes that are supported by the printer. Valid values are described in the following table.
	DuplexModes *[]PrintDuplexMode `json:"duplexModes,omitempty"`

	FeedDirections *[]PrinterFeedDirection `json:"feedDirections,omitempty"`

	// The list of feed orientations that are supported by the printer.
	FeedOrientations *[]PrinterFeedOrientation `json:"feedOrientations,omitempty"`

	// Finishing processes the printer supports for a printed document.
	Finishings *[]PrintFinishing `json:"finishings,omitempty"`

	// Supported input bins for the printer.
	InputBins *[]string `json:"inputBins,omitempty"`

	// True if color printing is supported by the printer; false otherwise. Read-only.
	IsColorPrintingSupported nullable.Type[bool] `json:"isColorPrintingSupported,omitempty"`

	// True if the printer supports printing by page ranges; false otherwise.
	IsPageRangeSupported nullable.Type[bool] `json:"isPageRangeSupported,omitempty"`

	// A list of supported left margins(in microns) for the printer.
	LeftMargins *[]int64 `json:"leftMargins,omitempty"`

	// The media (that is, paper) colors supported by the printer.
	MediaColors *[]string `json:"mediaColors,omitempty"`

	// The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are
	// in the following table.
	MediaSizes *[]string `json:"mediaSizes,omitempty"`

	// The media types supported by the printer.
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// The presentation directions supported by the printer. Supported values are described in the following table.
	MultipageLayouts *[]PrintMultipageLayout `json:"multipageLayouts,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The print orientations supported by the printer. Valid values are described in the following table.
	Orientations *[]PrintOrientation `json:"orientations,omitempty"`

	// The printer's supported output bins (trays).
	OutputBins *[]string `json:"outputBins,omitempty"`

	// Supported number of Input Pages to impose upon a single Impression.
	PagesPerSheet *[]int64 `json:"pagesPerSheet,omitempty"`

	// The print qualities supported by the printer.
	Qualities *[]PrintQuality `json:"qualities,omitempty"`

	// A list of supported right margins(in microns) for the printer.
	RightMargins *[]int64 `json:"rightMargins,omitempty"`

	// Supported print scalings.
	Scalings *[]PrintScaling `json:"scalings,omitempty"`

	SupportedColorConfigurations    *[]PrintColorConfiguration    `json:"supportedColorConfigurations,omitempty"`
	SupportedCopiesPerJob           *IntegerRange                 `json:"supportedCopiesPerJob,omitempty"`
	SupportedDocumentMimeTypes      *[]string                     `json:"supportedDocumentMimeTypes,omitempty"`
	SupportedDuplexConfigurations   *[]PrintDuplexConfiguration   `json:"supportedDuplexConfigurations,omitempty"`
	SupportedFinishings             *[]PrintFinishing             `json:"supportedFinishings,omitempty"`
	SupportedMediaColors            *[]string                     `json:"supportedMediaColors,omitempty"`
	SupportedMediaSizes             *[]string                     `json:"supportedMediaSizes,omitempty"`
	SupportedMediaTypes             *[]PrintMediaType             `json:"supportedMediaTypes,omitempty"`
	SupportedOrientations           *[]PrintOrientation           `json:"supportedOrientations,omitempty"`
	SupportedOutputBins             *[]string                     `json:"supportedOutputBins,omitempty"`
	SupportedPagesPerSheet          *IntegerRange                 `json:"supportedPagesPerSheet,omitempty"`
	SupportedPresentationDirections *[]PrintPresentationDirection `json:"supportedPresentationDirections,omitempty"`
	SupportedPrintQualities         *[]PrintQuality               `json:"supportedPrintQualities,omitempty"`

	// True if the printer supports scaling PDF pages to match the print media size; false otherwise.
	SupportsFitPdfToPage nullable.Type[bool] `json:"supportsFitPdfToPage,omitempty"`

	// A list of supported top margins(in microns) for the printer.
	TopMargins *[]int64 `json:"topMargins,omitempty"`
}

var _ json.Marshaler = PrinterCapabilities{}

func (s PrinterCapabilities) MarshalJSON() ([]byte, error) {
	type wrapper PrinterCapabilities
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrinterCapabilities: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrinterCapabilities: %+v", err)
	}

	delete(decoded, "isColorPrintingSupported")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrinterCapabilities: %+v", err)
	}

	return encoded, nil
}
