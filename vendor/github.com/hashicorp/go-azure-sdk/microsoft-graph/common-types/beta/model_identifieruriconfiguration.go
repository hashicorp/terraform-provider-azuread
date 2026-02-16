package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentifierUriConfiguration struct {
	// Block new identifier URIs for applications, unless they are the 'default' URI of the format api://{appId} or
	// api://{tenantId}/{appId}.
	NonDefaultUriAddition *IdentifierUriRestriction `json:"nonDefaultUriAddition,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Block new identifier URIs for applications, unless they contain a unique tenant identifier like the tenant ID, appId
	// (client ID), or verified domain. For example, api://{tenatId}/string, api://{appId}/string,
	// {scheme}://string/{tenantId}, {scheme}://string/{appId}, https://{verified-domain.com}/path,
	// {scheme}://{subdomain}.{verified-domain.com}/path.
	UriAdditionWithoutUniqueTenantIdentifier *IdentifierUriRestriction `json:"uriAdditionWithoutUniqueTenantIdentifier,omitempty"`
}
