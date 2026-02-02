package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversionUserDetails struct {
	ConvertedToInternalUserDateTime *string `json:"convertedToInternalUserDateTime,omitempty"`

	// Name displayed for the user.
	DisplayName *string `json:"displayName,omitempty"`

	// The SMTP address for the user.
	Mail *string `json:"mail,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The user principal name (UPN) of the user.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
