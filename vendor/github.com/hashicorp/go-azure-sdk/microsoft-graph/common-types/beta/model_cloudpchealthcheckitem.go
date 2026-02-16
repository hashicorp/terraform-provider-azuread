package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCHealthCheckItem struct {
	// Additional message for this health check.
	AdditionalDetails nullable.Type[string] `json:"additionalDetails,omitempty"`

	// The connectivity health check item name.
	DisplayName *string `json:"displayName,omitempty"`

	// Timestamp when the last check occurs. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC).
	// For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
	LastHealthCheckDateTime nullable.Type[string] `json:"lastHealthCheckDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Result *CloudPCConnectivityEventResult `json:"result,omitempty"`
}
