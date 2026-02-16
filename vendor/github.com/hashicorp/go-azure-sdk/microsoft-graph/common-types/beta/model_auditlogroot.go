package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditLogRoot struct {
	// Represents a custom security attribute audit log.
	CustomSecurityAttributeAudits *[]CustomSecurityAttributeAudit `json:"customSecurityAttributeAudits,omitempty"`

	DirectoryAudits       *[]DirectoryAudit            `json:"directoryAudits,omitempty"`
	DirectoryProvisioning *[]ProvisioningObjectSummary `json:"directoryProvisioning,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents an action performed by the Microsoft Entra provisioning service and its associated properties.
	Provisioning *[]ProvisioningObjectSummary `json:"provisioning,omitempty"`

	SignIns *[]SignIn            `json:"signIns,omitempty"`
	SignUps *[]SelfServiceSignUp `json:"signUps,omitempty"`
}
