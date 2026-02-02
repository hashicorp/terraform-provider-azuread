package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyConfigurationPartner struct {
	// Determines the partner-specific configuration for automatic user consent settings. Unless specifically configured,
	// the inboundAllowed and outboundAllowed properties are null and inherit from the default settings, which is always
	// false.
	AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration `json:"automaticUserConsentSettings,omitempty"`

	// Defines your partner-specific configuration for users from other organizations accessing your resources via Microsoft
	// Entra B2B collaboration.
	B2bCollaborationInbound CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationInbound"`

	// Defines your partner-specific configuration for users in your organization going outbound to access resources in
	// another organization via Microsoft Entra B2B collaboration.
	B2bCollaborationOutbound CrossTenantAccessPolicyB2BSetting `json:"b2bCollaborationOutbound"`

	// Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B
	// direct connect.
	B2bDirectConnectInbound CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectInbound"`

	// Defines your partner-specific configuration for users in your organization going outbound to access resources in
	// another organization via Microsoft Entra B2B direct connect.
	B2bDirectConnectOutbound CrossTenantAccessPolicyB2BSetting `json:"b2bDirectConnectOutbound"`

	// Defines the cross-tenant policy for the synchronization of users from a partner tenant. Use this user synchronization
	// policy to streamline collaboration between users in a multitenant organization by automating the creation, update,
	// and deletion of users from one tenant to another.
	IdentitySynchronization *CrossTenantIdentitySyncPolicyPartner `json:"identitySynchronization,omitempty"`

	// Determines the partner-specific configuration for trusting other Conditional Access claims from external Microsoft
	// Entra organizations.
	InboundTrust *CrossTenantAccessPolicyInboundTrust `json:"inboundTrust,omitempty"`

	// Identifies whether a tenant is a member of a multitenant organization.
	IsInMultiTenantOrganization nullable.Type[bool] `json:"isInMultiTenantOrganization,omitempty"`

	// Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
	IsServiceProvider nullable.Type[bool] `json:"isServiceProvider,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The tenant identifier for the partner Microsoft Entra organization. Read-only. Key.
	TenantId *string `json:"tenantId,omitempty"`

	// Defines the partner-specific tenant restrictions configuration for users in your organization who access a partner
	// organization using partner supplied identities on your network or devices.
	TenantRestrictions *CrossTenantAccessPolicyTenantRestrictions `json:"tenantRestrictions,omitempty"`
}

var _ json.Marshaler = CrossTenantAccessPolicyConfigurationPartner{}

func (s CrossTenantAccessPolicyConfigurationPartner) MarshalJSON() ([]byte, error) {
	type wrapper CrossTenantAccessPolicyConfigurationPartner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CrossTenantAccessPolicyConfigurationPartner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CrossTenantAccessPolicyConfigurationPartner: %+v", err)
	}

	delete(decoded, "tenantId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CrossTenantAccessPolicyConfigurationPartner: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CrossTenantAccessPolicyConfigurationPartner{}

func (s *CrossTenantAccessPolicyConfigurationPartner) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration        `json:"automaticUserConsentSettings,omitempty"`
		IdentitySynchronization      *CrossTenantIdentitySyncPolicyPartner      `json:"identitySynchronization,omitempty"`
		InboundTrust                 *CrossTenantAccessPolicyInboundTrust       `json:"inboundTrust,omitempty"`
		IsInMultiTenantOrganization  nullable.Type[bool]                        `json:"isInMultiTenantOrganization,omitempty"`
		IsServiceProvider            nullable.Type[bool]                        `json:"isServiceProvider,omitempty"`
		ODataId                      *string                                    `json:"@odata.id,omitempty"`
		ODataType                    *string                                    `json:"@odata.type,omitempty"`
		TenantId                     *string                                    `json:"tenantId,omitempty"`
		TenantRestrictions           *CrossTenantAccessPolicyTenantRestrictions `json:"tenantRestrictions,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AutomaticUserConsentSettings = decoded.AutomaticUserConsentSettings
	s.IdentitySynchronization = decoded.IdentitySynchronization
	s.InboundTrust = decoded.InboundTrust
	s.IsInMultiTenantOrganization = decoded.IsInMultiTenantOrganization
	s.IsServiceProvider = decoded.IsServiceProvider
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TenantId = decoded.TenantId
	s.TenantRestrictions = decoded.TenantRestrictions

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CrossTenantAccessPolicyConfigurationPartner into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["b2bCollaborationInbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bCollaborationInbound' for 'CrossTenantAccessPolicyConfigurationPartner': %+v", err)
		}
		s.B2bCollaborationInbound = impl
	}

	if v, ok := temp["b2bCollaborationOutbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bCollaborationOutbound' for 'CrossTenantAccessPolicyConfigurationPartner': %+v", err)
		}
		s.B2bCollaborationOutbound = impl
	}

	if v, ok := temp["b2bDirectConnectInbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bDirectConnectInbound' for 'CrossTenantAccessPolicyConfigurationPartner': %+v", err)
		}
		s.B2bDirectConnectInbound = impl
	}

	if v, ok := temp["b2bDirectConnectOutbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bDirectConnectOutbound' for 'CrossTenantAccessPolicyConfigurationPartner': %+v", err)
		}
		s.B2bDirectConnectOutbound = impl
	}

	return nil
}
