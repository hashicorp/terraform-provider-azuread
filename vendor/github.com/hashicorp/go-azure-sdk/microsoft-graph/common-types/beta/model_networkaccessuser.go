package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUser struct {
	// User display Name.
	DisplayName *string `json:"displayName,omitempty"`

	FirstAccessDateTime *string `json:"firstAccessDateTime,omitempty"`

	// The date and time of the most recent access.
	LastAccessDateTime *string `json:"lastAccessDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TotalBytesReceived *int64                    `json:"totalBytesReceived,omitempty"`
	TotalBytesSent     *int64                    `json:"totalBytesSent,omitempty"`
	TrafficType        *NetworkaccessTrafficType `json:"trafficType,omitempty"`
	TransactionCount   *int64                    `json:"transactionCount,omitempty"`

	// The ID for the user.
	UserId *string `json:"userId,omitempty"`

	// A unique identifier that is associated with a user in a system or directory. Typically, this value is an email
	// address that is used for user authentication and identification.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`

	UserType *NetworkaccessUserType `json:"userType,omitempty"`
}
