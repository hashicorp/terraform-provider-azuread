package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfiguration struct {
	// Whether the printer should collate pages when printing multiple copies of a multi-page document.
	Collate nullable.Type[bool] `json:"collate,omitempty"`

	// The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
	ColorMode *PrintColorMode `json:"colorMode,omitempty"`

	// The number of copies that should be printed. Read-only.
	Copies nullable.Type[int64] `json:"copies,omitempty"`

	// The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
	Dpi nullable.Type[int64] `json:"dpi,omitempty"`

	// The duplex mode the printer should use when printing the job. Valid values are described in the table below.
	// Read-only.
	DuplexMode *PrintDuplexMode `json:"duplexMode,omitempty"`

	// The orientation to use when feeding media into the printer. Valid values are described in the following table.
	// Read-only.
	FeedOrientation *PrinterFeedOrientation `json:"feedOrientation,omitempty"`

	// Finishing processes to use when printing.
	Finishings *[]PrintFinishing `json:"finishings,omitempty"`

	FitPdfToPage nullable.Type[bool] `json:"fitPdfToPage,omitempty"`

	// The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
	InputBin nullable.Type[string] `json:"inputBin,omitempty"`

	// The margin settings to use when printing.
	Margin *PrintMargin `json:"margin,omitempty"`

	// The media sizeto use when printing. Supports standard size names for ISO and ANSI media sizes. Valid values are
	// listed in the printerCapabilities topic.
	MediaSize nullable.Type[string] `json:"mediaSize,omitempty"`

	// The default media (such as paper) type to print the document on.
	MediaType nullable.Type[string] `json:"mediaType,omitempty"`

	// The direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the
	// following table.
	MultipageLayout *PrintMultipageLayout `json:"multipageLayout,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The orientation setting the printer should use when printing the job. Valid values are described in the following
	// table.
	Orientation *PrintOrientation `json:"orientation,omitempty"`

	// The output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
	OutputBin nullable.Type[string] `json:"outputBin,omitempty"`

	// The page ranges to print. Read-only.
	PageRanges *[]IntegerRange `json:"pageRanges,omitempty"`

	// The number of document pages to print on each sheet.
	PagesPerSheet nullable.Type[int64] `json:"pagesPerSheet,omitempty"`

	// The print quality to use when printing the job. Valid values are described in the table below. Read-only.
	Quality *PrintQuality `json:"quality,omitempty"`

	// Specifies how the printer should scale the document data to fit the requested media. Valid values are described in
	// the following table.
	Scaling *PrintScaling `json:"scaling,omitempty"`
}

var _ json.Marshaler = PrintJobConfiguration{}

func (s PrintJobConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper PrintJobConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintJobConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintJobConfiguration: %+v", err)
	}

	delete(decoded, "colorMode")
	delete(decoded, "copies")
	delete(decoded, "dpi")
	delete(decoded, "duplexMode")
	delete(decoded, "feedOrientation")
	delete(decoded, "pageRanges")
	delete(decoded, "quality")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintJobConfiguration: %+v", err)
	}

	return encoded, nil
}
