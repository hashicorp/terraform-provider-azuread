package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeSummary struct {
	// A collection of the user experience analytics device scope Unique Identifiers that are enabled and finished
	// recalculating the report metric.
	CompletedDeviceScopeIds *[]string `json:"completedDeviceScopeIds,omitempty"`

	// A collection of user experience analytics device scope Unique Identitfiers that are enabled but there is insufficient
	// data to calculate results.
	InsufficientDataDeviceScopeIds *[]string `json:"insufficientDataDeviceScopeIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of user experience analytics device scopes. Valid values -2147483648 to 2147483647
	TotalDeviceScopes *int64 `json:"totalDeviceScopes,omitempty"`

	// The total number of user experience analytics device scopes that are enabled. Valid values -2147483648 to 2147483647
	TotalDeviceScopesEnabled *int64 `json:"totalDeviceScopesEnabled,omitempty"`
}
