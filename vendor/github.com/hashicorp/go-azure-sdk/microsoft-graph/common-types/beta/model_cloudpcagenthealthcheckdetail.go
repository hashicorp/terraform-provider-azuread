package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentHealthCheckDetail struct {
	AdditionalDetails       *[]KeyValuePair                    `json:"additionalDetails,omitempty"`
	ErrorMessage            nullable.Type[string]              `json:"errorMessage,omitempty"`
	ErrorType               *CloudPCAgentHealthCheckErrorType  `json:"errorType,omitempty"`
	HealthCheckName         *string                            `json:"healthCheckName,omitempty"`
	HealthCheckResultType   *CloudPCAgentHealthCheckResultType `json:"healthCheckResultType,omitempty"`
	LastHealthCheckDateTime nullable.Type[string]              `json:"lastHealthCheckDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
