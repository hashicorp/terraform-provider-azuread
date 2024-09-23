package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserAccount struct {
	// The displayed name of the user account.
	AccountName nullable.Type[string] `json:"accountName,omitempty"`

	// The user object identifier in Microsoft Entra ID.
	AzureAdUserId nullable.Type[string] `json:"azureAdUserId,omitempty"`

	// The user display name in Microsoft Entra ID.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The name of the Active Directory domain of which the user is a member.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The user principal name of the account in Microsoft Entra ID.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The local security identifier of the user account.
	UserSid nullable.Type[string] `json:"userSid,omitempty"`
}
