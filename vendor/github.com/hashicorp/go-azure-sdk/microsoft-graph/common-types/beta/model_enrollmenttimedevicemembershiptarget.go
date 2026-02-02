package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentTimeDeviceMembershipTarget struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifiers of the targets that devices will become members of when enrolled with the asociated profile.
	TargetId nullable.Type[string] `json:"targetId,omitempty"`

	// Represents the type of the targets that devices will become members of when enrolled with the associated profile.
	// Possible values are staticSecurityGroup.
	TargetType *EnrollmentTimeDeviceMembershipTargetType `json:"targetType,omitempty"`
}
