package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTimeDeviceMembershipTargetStatus struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifiers of the targets that devices will become members of when enrolled with the asociated profile.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`

	// Represents the Validation error of the device membership target.The API will validate the device membership targets
	// specified by the admin to ensure that they exist, that they are of the proper type, and any other target requirements
	// are met such as that the Intune Device Provisioning First Party App is an owner of the target.
	TargetValidationErrorCode *EnrollmentTimeDeviceMembershipTargetValidationErrorCode `json:"targetValidationErrorCode,omitempty"`
}
