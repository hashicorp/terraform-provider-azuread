package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InvitationRedemptionIdentityProviderConfiguration = DefaultInvitationRedemptionIdentityProviderConfiguration{}

type DefaultInvitationRedemptionIdentityProviderConfiguration struct {

	// Fields inherited from InvitationRedemptionIdentityProviderConfiguration

	// The fallback identity provider to be used in case no primary identity provider can be used for guest invitation
	// redemption. Possible values are: defaultConfiguredIdp, emailOneTimePasscode, or microsoftAccount.
	FallbackIdentityProvider *B2bIdentityProvidersType `json:"fallbackIdentityProvider,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of identity providers in priority order of preference to be used for guest invitation redemption. Possible
	// values are: azureActiveDirectory, externalFederation, or socialIdentityProviders.
	PrimaryIdentityProviderPrecedenceOrder *[]B2bIdentityProvidersType `json:"primaryIdentityProviderPrecedenceOrder,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DefaultInvitationRedemptionIdentityProviderConfiguration) InvitationRedemptionIdentityProviderConfiguration() BaseInvitationRedemptionIdentityProviderConfigurationImpl {
	return BaseInvitationRedemptionIdentityProviderConfigurationImpl{
		FallbackIdentityProvider:               s.FallbackIdentityProvider,
		ODataId:                                s.ODataId,
		ODataType:                              s.ODataType,
		PrimaryIdentityProviderPrecedenceOrder: s.PrimaryIdentityProviderPrecedenceOrder,
	}
}

var _ json.Marshaler = DefaultInvitationRedemptionIdentityProviderConfiguration{}

func (s DefaultInvitationRedemptionIdentityProviderConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper DefaultInvitationRedemptionIdentityProviderConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefaultInvitationRedemptionIdentityProviderConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefaultInvitationRedemptionIdentityProviderConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.defaultInvitationRedemptionIdentityProviderConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefaultInvitationRedemptionIdentityProviderConfiguration: %+v", err)
	}

	return encoded, nil
}
