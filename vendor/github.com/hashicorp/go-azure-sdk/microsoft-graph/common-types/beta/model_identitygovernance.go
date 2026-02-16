package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernance struct {
	AccessReviews         *AccessReviewSet                               `json:"accessReviews,omitempty"`
	AppConsent            *AppConsentApprovalRoute                       `json:"appConsent,omitempty"`
	EntitlementManagement *EntitlementManagement                         `json:"entitlementManagement,omitempty"`
	LifecycleWorkflows    *IdentityGovernanceLifecycleWorkflowsContainer `json:"lifecycleWorkflows,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PermissionsAnalytics  *PermissionsAnalyticsAggregation `json:"permissionsAnalytics,omitempty"`
	PermissionsManagement *PermissionsManagement           `json:"permissionsManagement,omitempty"`
	PrivilegedAccess      *PrivilegedAccessRoot            `json:"privilegedAccess,omitempty"`
	RoleManagementAlerts  *RoleManagementAlert             `json:"roleManagementAlerts,omitempty"`
	TermsOfUse            *TermsOfUseContainer             `json:"termsOfUse,omitempty"`
}
