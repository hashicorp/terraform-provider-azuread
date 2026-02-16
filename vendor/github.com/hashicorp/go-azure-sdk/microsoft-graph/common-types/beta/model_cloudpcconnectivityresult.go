package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityResult struct {
	// A list of failed health check items. If the status property is available, this property is empty.
	FailedHealthCheckItems *[]CloudPCHealthCheckItem `json:"failedHealthCheckItems,omitempty"`

	// The last modified time for connectivity status of the Cloud PC. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like
	// this: 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Status *CloudPCConnectivityStatus `json:"status,omitempty"`

	// Datetime when the status was updated. This property is deprecated and will no longer be supported effective August
	// 31, 2024. Use lastModifiedDateTime instead. Read-Only.
	UpdatedDateTime *string `json:"updatedDateTime,omitempty"`
}
