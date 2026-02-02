package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationInfo struct {
	// The collection of unique identifiers that can be associated with a user and can be used to bind the Microsoft Entra
	// user to a certificate for authentication and authorization into non-Azure AD environments. The identifiers must be
	// unique in the tenant.
	CertificateUserIds *[]string `json:"certificateUserIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
