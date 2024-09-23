package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataUserMatchingSetting struct {
	// The RefUserMatchTarget for matching a user from the source with a Microsoft Entra user object.
	MatchTarget *IndustryDataUserMatchTargetReferenceValue `json:"matchTarget,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The priority order to apply when a user has multiple RefRole codes assigned.
	PriorityOrder *int64 `json:"priorityOrder,omitempty"`

	RoleGroup        *IndustryDataRoleGroup                    `json:"roleGroup,omitempty"`
	SourceIdentifier *IndustryDataIdentifierTypeReferenceValue `json:"sourceIdentifier,omitempty"`
}
