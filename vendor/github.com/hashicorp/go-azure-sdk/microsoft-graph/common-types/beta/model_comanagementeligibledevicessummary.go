package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDevicesSummary struct {
	// Count of devices already Co-Managed
	ComanagedCount *int64 `json:"comanagedCount,omitempty"`

	// Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
	EligibleButNotAzureAdJoinedCount *int64 `json:"eligibleButNotAzureAdJoinedCount,omitempty"`

	// Count of devices fully eligible for Co-Management
	EligibleCount *int64 `json:"eligibleCount,omitempty"`

	// Count of devices ineligible for Co-Management
	IneligibleCount *int64 `json:"ineligibleCount,omitempty"`

	// Count of devices that will be eligible for Co-Management after an OS update
	NeedsOsUpdateCount *int64 `json:"needsOsUpdateCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of devices scheduled for Co-Management enrollment. Valid values 0 to 9999999
	ScheduledForEnrollmentCount *int64 `json:"scheduledForEnrollmentCount,omitempty"`
}
