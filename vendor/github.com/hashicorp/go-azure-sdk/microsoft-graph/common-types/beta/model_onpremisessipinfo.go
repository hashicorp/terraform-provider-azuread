package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesSipInfo struct {
	// Indicates whether the user is currently enabled for on-premises Skype for Business.
	IsSipEnabled *bool `json:"isSipEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates a fully qualified DNS name of the Microsoft Online Communications Server deployment.
	SipDeploymentLocation nullable.Type[string] `json:"sipDeploymentLocation,omitempty"`

	// Serves as a unique identifier for each user on the on-premises Skype for Business.
	SipPrimaryAddress nullable.Type[string] `json:"sipPrimaryAddress,omitempty"`
}
