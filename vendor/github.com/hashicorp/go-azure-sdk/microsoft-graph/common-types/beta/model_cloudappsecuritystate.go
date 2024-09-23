package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecurityState struct {
	// Destination IP Address of the connection to the cloud application/service.
	DestinationServiceIp nullable.Type[string] `json:"destinationServiceIp,omitempty"`

	// Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
	DestinationServiceName nullable.Type[string] `json:"destinationServiceName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which
	// equates to a percentage.
	RiskScore nullable.Type[string] `json:"riskScore,omitempty"`
}
