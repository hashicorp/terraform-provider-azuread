package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetResource struct {
	// Indicates the visible name defined for the resource. Typically specified when the resource is created.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and
	// unknownFutureValue
	GroupType *GroupType `json:"groupType,omitempty"`

	// Indicates the unique ID of the resource.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
	ModifiedProperties *[]ModifiedProperty `json:"modifiedProperties,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Describes the resource type. Example values include Application, Group, ServicePrincipal, and User.
	Type nullable.Type[string] `json:"type,omitempty"`

	// When type is set to User, this includes the user name that initiated the action; null for other types.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
