package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySensorSettings struct {
	// Description of the sensor.
	Description *string `json:"description,omitempty"`

	// DNS names for the domain controller
	DomainControllerDnsNames *[]string `json:"domainControllerDnsNames,omitempty"`

	// Indicates whether to delay updates for the sensor.
	IsDelayedDeploymentEnabled nullable.Type[bool] `json:"isDelayedDeploymentEnabled,omitempty"`

	NetworkAdapters *[]SecurityNetworkAdapter `json:"networkAdapters,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
