package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditActor struct {
	// Name of the Application.
	ApplicationDisplayName nullable.Type[string] `json:"applicationDisplayName,omitempty"`

	// AAD Application Id.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// Actor Type.
	AuditActorType nullable.Type[string] `json:"auditActorType,omitempty"`

	// IPAddress.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Remote Tenant Id
	RemoteTenantId nullable.Type[string] `json:"remoteTenantId,omitempty"`

	// Remote User Id
	RemoteUserId nullable.Type[string] `json:"remoteUserId,omitempty"`

	// Service Principal Name (SPN).
	ServicePrincipalName nullable.Type[string] `json:"servicePrincipalName,omitempty"`

	// Actor Type.
	Type nullable.Type[string] `json:"type,omitempty"`

	// User Id.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// List of user permissions when the audit was performed.
	UserPermissions *[]string `json:"userPermissions,omitempty"`

	// User Principal Name (UPN).
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// List of user scope tags when the audit was performed.
	UserRoleScopeTags *[]RoleScopeTagInfo `json:"userRoleScopeTags,omitempty"`
}
