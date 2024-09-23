package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationCount struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Provides the registration count for your tenant.
	RegistrationCount *int64 `json:"registrationCount,omitempty"`

	RegistrationStatus *RegistrationStatusType `json:"registrationStatus,omitempty"`
}
