package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadByFilter struct {
	// Represents type of the assignment filter.
	AssignmentFilterType *DeviceAndAppManagementAssignmentFilterType `json:"assignmentFilterType,omitempty"`

	// The Azure AD security group ID
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The policy identifier
	PayloadId nullable.Type[string] `json:"payloadId,omitempty"`

	// This enum represents associated assignment payload type
	PayloadType *AssociatedAssignmentPayloadType `json:"payloadType,omitempty"`
}
