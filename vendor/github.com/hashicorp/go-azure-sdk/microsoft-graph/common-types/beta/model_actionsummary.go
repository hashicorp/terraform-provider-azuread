package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionSummary struct {
	// This is the number of authorization system actions that have been assigned to the identity.
	Assigned *int64 `json:"assigned,omitempty"`

	// This is the number of authorization system actions that the identity has exercised in the last 90 days.
	Available *int64 `json:"available,omitempty"`

	// This is the maximum number of actions that are available in the authorization system.
	Exercised *int64 `json:"exercised,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
