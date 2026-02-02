package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVmMetadata struct {
	CloudProvider *SecurityVmCloudProvider `json:"cloudProvider,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique identifier of the Azure resource.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Unique identifier of the Azure subscription the customer tenant belongs to.
	SubscriptionId nullable.Type[string] `json:"subscriptionId,omitempty"`

	// Unique identifier of the virtual machine instance.
	VmId nullable.Type[string] `json:"vmId,omitempty"`
}
