package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySslCertificateEntity struct {
	// A physical address of the entity.
	Address *PhysicalAddress `json:"address,omitempty"`

	// Alternate names for this entity that are part of the certificate.
	AlternateNames *[]string `json:"alternateNames,omitempty"`

	// A common name for this entity.
	CommonName nullable.Type[string] `json:"commonName,omitempty"`

	// An email for this entity.
	Email nullable.Type[string] `json:"email,omitempty"`

	// If the entity is a person, this is the person's given name (first name).
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If the entity is an organization, this is the name of the organization.
	OrganizationName nullable.Type[string] `json:"organizationName,omitempty"`

	// If the entity is an organization, this communicates if a unit in the organization is named on the entity.
	OrganizationUnitName nullable.Type[string] `json:"organizationUnitName,omitempty"`

	// A serial number assigned to the entity; usually only available if the entity is the issuer.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// If the entity is a person, this is the person's surname (last name).
	Surname nullable.Type[string] `json:"surname,omitempty"`
}
