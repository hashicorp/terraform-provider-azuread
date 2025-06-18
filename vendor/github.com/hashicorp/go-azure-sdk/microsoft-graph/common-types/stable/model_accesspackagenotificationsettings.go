package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageNotificationSettings struct {
	// Indicates if notification emails for an access package are disabled within an access package assignment policy.
	IsAssignmentNotificationDisabled *bool `json:"isAssignmentNotificationDisabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
