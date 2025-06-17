package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudiencesConfiguration struct {
	// Setting to allow or disallow creation of apps with multitenant signInAudience.
	AzureAdMultipleOrgs *AudienceRestriction `json:"azureAdMultipleOrgs,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting to allow or disallow creation of apps with personal Microsoft account signInAudience.
	PersonalMicrosoftAccount *AudienceRestriction `json:"personalMicrosoftAccount,omitempty"`
}
