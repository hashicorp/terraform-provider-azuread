package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentStatusDetail struct {
	AgentVersionNumber              nullable.Type[string]             `json:"agentVersionNumber,omitempty"`
	CloudPCId                       nullable.Type[string]             `json:"cloudPcId,omitempty"`
	DiagnosticResultMessage         nullable.Type[string]             `json:"diagnosticResultMessage,omitempty"`
	DiagnosticResultType            *CloudPCAgentDiagnosticResultType `json:"diagnosticResultType,omitempty"`
	HealthCheckSummary              *CloudPCAgentHealthCheckSummary   `json:"healthCheckSummary,omitempty"`
	HealthStatus                    *CloudPCAgentHealthStatus         `json:"healthStatus,omitempty"`
	LastHealthStatusCheckedDateTime nullable.Type[string]             `json:"lastHealthStatusCheckedDateTime,omitempty"`
	ManagedDeviceName               nullable.Type[string]             `json:"managedDeviceName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
