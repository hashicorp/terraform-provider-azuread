package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecurityCustomerInsight struct {
	// Details of the customer's Entra tenant MFA policy configuration and usage.
	Mfa *PartnerSecurityCustomerMfaInsight `json:"mfa,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the customer.
	TenantId *string `json:"tenantId,omitempty"`
}
