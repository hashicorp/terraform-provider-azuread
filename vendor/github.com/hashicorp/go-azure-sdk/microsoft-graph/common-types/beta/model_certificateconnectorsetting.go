package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateConnectorSetting struct {
	// Certificate expire time
	CertExpiryTime *string `json:"certExpiryTime,omitempty"`

	// Version of certificate connector
	ConnectorVersion nullable.Type[string] `json:"connectorVersion,omitempty"`

	// Certificate connector enrollment error
	EnrollmentError nullable.Type[string] `json:"enrollmentError,omitempty"`

	// Last time certificate connector connected
	LastConnectorConnectionTime *string `json:"lastConnectorConnectionTime,omitempty"`

	// Version of last uploaded certificate connector
	LastUploadVersion *int64 `json:"lastUploadVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Certificate connector status
	Status *int64 `json:"status,omitempty"`
}
