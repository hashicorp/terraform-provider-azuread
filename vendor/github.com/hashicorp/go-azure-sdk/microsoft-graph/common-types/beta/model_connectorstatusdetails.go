package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetails struct {
	// Connector Instance Id
	ConnectorInstanceId nullable.Type[string] `json:"connectorInstanceId,omitempty"`

	// Connectors name for connector status
	ConnectorName *ConnectorName `json:"connectorName,omitempty"`

	// Event datetime
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Connector health state for connector status
	Status *ConnectorHealthState `json:"status,omitempty"`
}
