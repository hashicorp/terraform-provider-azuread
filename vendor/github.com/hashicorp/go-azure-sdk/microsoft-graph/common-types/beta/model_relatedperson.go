package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedPerson struct {
	// Name of the person.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Possible values are: manager, colleague, directReport, dotLineReport, assistant, dotLineManager, alternateContact,
	// friend, spouse, sibling, child, parent, sponsor, emergencyContact, other, unknownFutureValue.
	Relationship *PersonRelationship `json:"relationship,omitempty"`

	// The user's directory object ID (Microsoft Entra ID or CID).
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Email address or reference to person within the organization.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
