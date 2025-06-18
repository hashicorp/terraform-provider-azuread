package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAdministrationPolicyAssignment struct {
	AssignmentType *TeamsAdministrationAssignmentType `json:"assignmentType,omitempty"`

	// Represents the name of the policy.
	DisplayName *string `json:"displayName,omitempty"`

	// Represents the group identifier.
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the unique identifier for the policy.
	PolicyId *string `json:"policyId,omitempty"`
}
