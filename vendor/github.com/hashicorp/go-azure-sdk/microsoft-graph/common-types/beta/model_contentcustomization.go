package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentCustomization struct {
	// Represents the content options of External Identities to be customized throughout the authentication flow for a
	// tenant.
	AttributeCollection *[]KeyValue `json:"attributeCollection,omitempty"`

	// A relative URL for the content options of External Identities to be customized throughout the authentication flow for
	// a tenant.
	AttributeCollectionRelativeUrl nullable.Type[string] `json:"attributeCollectionRelativeUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents content options to customize during MFA proofup interruptions.
	RegistrationCampaign *[]KeyValue `json:"registrationCampaign,omitempty"`

	// The relative URL of the content options to customize during MFA proofup interruptions.
	RegistrationCampaignRelativeUrl nullable.Type[string] `json:"registrationCampaignRelativeUrl,omitempty"`
}
