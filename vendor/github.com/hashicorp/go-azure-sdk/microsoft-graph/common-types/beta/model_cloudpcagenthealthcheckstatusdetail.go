package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentHealthCheckStatusDetail struct {
	AdditionalHealthCheckMessage *string                       `json:"additionalHealthCheckMessage,omitempty"`
	CloudPCId                    nullable.Type[string]         `json:"cloudPcId,omitempty"`
	HealthCheckState             *CloudPCAgentHealthCheckState `json:"healthCheckState,omitempty"`
	LastModifiedDateTime         nullable.Type[string]         `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}
