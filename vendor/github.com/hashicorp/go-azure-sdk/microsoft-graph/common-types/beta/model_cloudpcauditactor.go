package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditActor struct {
	// Name of the application.
	ApplicationDisplayName nullable.Type[string] `json:"applicationDisplayName,omitempty"`

	// Microsoft Entra application ID.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// IP address.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The delegated partner tenant ID.
	RemoteTenantId nullable.Type[string] `json:"remoteTenantId,omitempty"`

	// The delegated partner user ID.
	RemoteUserId nullable.Type[string] `json:"remoteUserId,omitempty"`

	// Service Principal Name (SPN).
	ServicePrincipalName nullable.Type[string] `json:"servicePrincipalName,omitempty"`

	Type *CloudPCAuditActorType `json:"type,omitempty"`

	// Microsoft Entra user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// List of user permissions and application permissions when the audit event was performed.
	UserPermissions *[]string `json:"userPermissions,omitempty"`

	// User Principal Name (UPN).
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// List of role scope tags.
	UserRoleScopeTags *[]CloudPCUserRoleScopeTagInfo `json:"userRoleScopeTags,omitempty"`
}
