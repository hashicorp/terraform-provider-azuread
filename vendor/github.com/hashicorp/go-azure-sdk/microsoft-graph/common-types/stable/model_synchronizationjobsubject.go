package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobSubject struct {
	// Principals that you would like to provision.
	Links *SynchronizationLinkedObjects `json:"links,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier of an object to which a synchronizationJob is to be applied. Can be one of the following: An
	// onPremisesDistinguishedName for synchronization from Active Directory to Azure AD.The user ID for synchronization
	// from Microsoft Entra ID to a third-party.The Worker ID of the Workday worker for synchronization from Workday to
	// either Active Directory or Microsoft Entra ID.
	ObjectId nullable.Type[string] `json:"objectId,omitempty"`

	// The type of the object to which a synchronizationJob is to be applied. Can be one of the following: user for
	// synchronizing between Active Directory and Azure AD.User for synchronizing a user between Microsoft Entra ID and a
	// third-party application. Worker for synchronization a user between Workday and either Active Directory or Microsoft
	// Entra ID.Group for synchronizing a group between Microsoft Entra ID and a third-party application.
	ObjectTypeName nullable.Type[string] `json:"objectTypeName,omitempty"`
}
