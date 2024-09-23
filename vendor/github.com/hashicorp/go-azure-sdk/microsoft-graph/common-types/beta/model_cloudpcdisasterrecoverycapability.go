package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDisasterRecoveryCapability struct {
	// The disaster recovery action that can be performed for the Cloud PC. The possible values are: none, failover,
	// failback, unknownFutureValue.
	CapabilityType *CloudPCDisasterRecoveryCapabilityType `json:"capabilityType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The primary and mainly used region where the Cloud PC is located.
	PrimaryRegion nullable.Type[string] `json:"primaryRegion,omitempty"`

	// The secondary region to which the Cloud PC can be failed over during a regional outage.
	SecondaryRegion nullable.Type[string] `json:"secondaryRegion,omitempty"`
}
