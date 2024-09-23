package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CrossTenantAccessPolicyConfigurationDefault{}

type CrossTenantAccessPolicyConfigurationDefault struct {
	// Determines the default configuration for automatic user consent settings. The inboundAllowed and outboundAllowed
	// properties are always false and can't be updated in the default configuration. Read-only.
	AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration `json:"automaticUserConsentSettings,omitempty"`

	// Defines your default configuration for users from other organizations accessing your resources via Microsoft Entra
	// B2B collaboration.
	B2bCollaborationInbound CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationInbound"`

	// Defines your default configuration for users in your organization going outbound to access resources in another
	// organization via Microsoft Entra B2B collaboration.
	B2bCollaborationOutbound CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationOutbound"`

	// Defines your default configuration for users from other organizations accessing your resources via Microsoft Entra
	// B2B direct connect.
	B2bDirectConnectInbound CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectInbound"`

	// Defines your default configuration for users in your organization going outbound to access resources in another
	// organization via Microsoft Entra B2B direct connect.
	B2bDirectConnectOutbound CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectOutbound"`

	// Determines the default configuration for trusting other Conditional Access claims from external Microsoft Entra
	// organizations.
	InboundTrust *CrossTenantAccessPolicyInboundTrust `json:"inboundTrust,omitempty"`

	// Defines the priority order based on which an identity provider is selected during invitation redemption for a guest
	// user.
	InvitationRedemptionIdentityProviderConfiguration *DefaultInvitationRedemptionIdentityProviderConfiguration `json:"invitationRedemptionIdentityProviderConfiguration,omitempty"`

	// If true, the default configuration is set to the system default configuration. If false, the default settings are
	// customized.
	IsServiceDefault nullable.Type[bool] `json:"isServiceDefault,omitempty"`

	// Defines the default tenant restrictions configuration for users in your organization who access an external
	// organization on your network or devices.
	TenantRestrictions *CrossTenantAccessPolicyTenantRestrictions `json:"tenantRestrictions,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CrossTenantAccessPolicyConfigurationDefault) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CrossTenantAccessPolicyConfigurationDefault{}

func (s CrossTenantAccessPolicyConfigurationDefault) MarshalJSON() ([]byte, error) {
	type wrapper CrossTenantAccessPolicyConfigurationDefault
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossTenantAccessPolicyConfigurationDefault: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantAccessPolicyConfigurationDefault: %+v", err)
	}

	delete(decoded, "automaticUserConsentSettings")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.crossTenantAccessPolicyConfigurationDefault"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossTenantAccessPolicyConfigurationDefault: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CrossTenantAccessPolicyConfigurationDefault{}

func (s *CrossTenantAccessPolicyConfigurationDefault) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AutomaticUserConsentSettings                      *InboundOutboundPolicyConfiguration                       `json:"automaticUserConsentSettings,omitempty"`
		InboundTrust                                      *CrossTenantAccessPolicyInboundTrust                      `json:"inboundTrust,omitempty"`
		InvitationRedemptionIdentityProviderConfiguration *DefaultInvitationRedemptionIdentityProviderConfiguration `json:"invitationRedemptionIdentityProviderConfiguration,omitempty"`
		IsServiceDefault                                  nullable.Type[bool]                                       `json:"isServiceDefault,omitempty"`
		TenantRestrictions                                *CrossTenantAccessPolicyTenantRestrictions                `json:"tenantRestrictions,omitempty"`
		Id                                                *string                                                   `json:"id,omitempty"`
		ODataId                                           *string                                                   `json:"@odata.id,omitempty"`
		ODataType                                         *string                                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AutomaticUserConsentSettings = decoded.AutomaticUserConsentSettings
	s.InboundTrust = decoded.InboundTrust
	s.InvitationRedemptionIdentityProviderConfiguration = decoded.InvitationRedemptionIdentityProviderConfiguration
	s.IsServiceDefault = decoded.IsServiceDefault
	s.TenantRestrictions = decoded.TenantRestrictions
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CrossTenantAccessPolicyConfigurationDefault into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["b2bCollaborationInbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bCollaborationInbound' for 'CrossTenantAccessPolicyConfigurationDefault': %+v", err)
		}
		s.B2bCollaborationInbound = impl
	}

	if v, ok := temp["b2bCollaborationOutbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bCollaborationOutbound' for 'CrossTenantAccessPolicyConfigurationDefault': %+v", err)
		}
		s.B2bCollaborationOutbound = impl
	}

	if v, ok := temp["b2bDirectConnectInbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bDirectConnectInbound' for 'CrossTenantAccessPolicyConfigurationDefault': %+v", err)
		}
		s.B2bDirectConnectInbound = impl
	}

	if v, ok := temp["b2bDirectConnectOutbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bDirectConnectOutbound' for 'CrossTenantAccessPolicyConfigurationDefault': %+v", err)
		}
		s.B2bDirectConnectOutbound = impl
	}

	return nil
}
