package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QrCodeImageDetails struct {
	// The binary representation of the QR code.
	BinaryValue nullable.Type[string] `json:"binaryValue,omitempty"`

	// Specifies how much of the QRCode can be corrupted while still maintaining its readable. The possible values are: l
	// (Low), m (Medium), q (Quartile), h ( High), unknownFutureValue.
	ErrorCorrectionLevel *ErrorCorrectionLevel `json:"errorCorrectionLevel,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Base64-encoded raw content of the QR code.
	RawContent nullable.Type[string] `json:"rawContent,omitempty"`

	// Version to create QR code image.
	Version nullable.Type[int64] `json:"version,omitempty"`
}
