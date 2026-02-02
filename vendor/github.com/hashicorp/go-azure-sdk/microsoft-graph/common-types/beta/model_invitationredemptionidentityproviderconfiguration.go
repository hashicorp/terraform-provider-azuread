package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationRedemptionIdentityProviderConfiguration interface {
	InvitationRedemptionIdentityProviderConfiguration() BaseInvitationRedemptionIdentityProviderConfigurationImpl
}

var _ InvitationRedemptionIdentityProviderConfiguration = BaseInvitationRedemptionIdentityProviderConfigurationImpl{}

type BaseInvitationRedemptionIdentityProviderConfigurationImpl struct {
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

func (s BaseInvitationRedemptionIdentityProviderConfigurationImpl) InvitationRedemptionIdentityProviderConfiguration() BaseInvitationRedemptionIdentityProviderConfigurationImpl {
	return s
}

var _ InvitationRedemptionIdentityProviderConfiguration = RawInvitationRedemptionIdentityProviderConfigurationImpl{}

// RawInvitationRedemptionIdentityProviderConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInvitationRedemptionIdentityProviderConfigurationImpl struct {
	invitationRedemptionIdentityProviderConfiguration BaseInvitationRedemptionIdentityProviderConfigurationImpl
	Type                                              string
	Values                                            map[string]interface{}
}

func (s RawInvitationRedemptionIdentityProviderConfigurationImpl) InvitationRedemptionIdentityProviderConfiguration() BaseInvitationRedemptionIdentityProviderConfigurationImpl {
	return s.invitationRedemptionIdentityProviderConfiguration
}

func UnmarshalInvitationRedemptionIdentityProviderConfigurationImplementation(input []byte) (InvitationRedemptionIdentityProviderConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InvitationRedemptionIdentityProviderConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.defaultInvitationRedemptionIdentityProviderConfiguration") {
		var out DefaultInvitationRedemptionIdentityProviderConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefaultInvitationRedemptionIdentityProviderConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseInvitationRedemptionIdentityProviderConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseInvitationRedemptionIdentityProviderConfigurationImpl: %+v", err)
	}

	return RawInvitationRedemptionIdentityProviderConfigurationImpl{
		invitationRedemptionIdentityProviderConfiguration: parent,
		Type:   value,
		Values: temp,
	}, nil

}
