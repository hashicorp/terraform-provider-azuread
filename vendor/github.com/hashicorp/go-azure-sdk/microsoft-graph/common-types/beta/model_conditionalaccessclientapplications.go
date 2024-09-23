package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessClientApplications struct {
	// Service principal IDs excluded from the policy scope.
	ExcludeServicePrincipals *[]string `json:"excludeServicePrincipals,omitempty"`

	// Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
	IncludeServicePrincipals *[]string `json:"includeServicePrincipals,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Filter that defines the dynamic-servicePrincipal-syntax rule to include/exclude service principals. A filter can use
	// custom security attributes to include/exclude service principals.
	ServicePrincipalFilter *ConditionalAccessFilter `json:"servicePrincipalFilter,omitempty"`
}
