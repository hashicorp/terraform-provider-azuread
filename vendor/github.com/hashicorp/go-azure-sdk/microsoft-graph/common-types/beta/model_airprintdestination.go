package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AirPrintDestination struct {
	// If true AirPrint connections are secured by Transport Layer Security (TLS). Default is false. Available in iOS 11.0
	// and later.
	ForceTls *bool `json:"forceTls,omitempty"`

	// The IP Address of the AirPrint destination.
	IPAddress *string `json:"ipAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The listening port of the AirPrint destination. If this key is not specified AirPrint will use the default port.
	// Available in iOS 11.0 and later.
	Port nullable.Type[int64] `json:"port,omitempty"`

	// The Resource Path associated with the printer. This corresponds to the rp parameter of the ipps.tcp Bonjour record.
	// For example: printers/CanonMG5300series, printers/XeroxPhaser7600, ipp/print, EpsonIPPPrinter.
	ResourcePath *string `json:"resourcePath,omitempty"`
}
