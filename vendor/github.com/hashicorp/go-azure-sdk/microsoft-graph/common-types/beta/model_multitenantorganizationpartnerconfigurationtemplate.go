package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MultiTenantOrganizationPartnerConfigurationTemplate{}

type MultiTenantOrganizationPartnerConfigurationTemplate struct {
	// Determines the partner-specific configuration for automatic user consent settings. Unless configured, the
	// inboundAllowed and outboundAllowed properties are null and inherit from the default settings, which is always false.
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

	// Determines the partner-specific configuration for trusting other Conditional Access claims from external Microsoft
	// Entra organizations.
	InboundTrust *CrossTenantAccessPolicyInboundTrust `json:"inboundTrust,omitempty"`

	TemplateApplicationLevel *TemplateApplicationLevel `json:"templateApplicationLevel,omitempty"`

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

func (s MultiTenantOrganizationPartnerConfigurationTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MultiTenantOrganizationPartnerConfigurationTemplate{}

func (s MultiTenantOrganizationPartnerConfigurationTemplate) MarshalJSON() ([]byte, error) {
	type wrapper MultiTenantOrganizationPartnerConfigurationTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MultiTenantOrganizationPartnerConfigurationTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MultiTenantOrganizationPartnerConfigurationTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.multiTenantOrganizationPartnerConfigurationTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MultiTenantOrganizationPartnerConfigurationTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MultiTenantOrganizationPartnerConfigurationTemplate{}

func (s *MultiTenantOrganizationPartnerConfigurationTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration  `json:"automaticUserConsentSettings,omitempty"`
		InboundTrust                 *CrossTenantAccessPolicyInboundTrust `json:"inboundTrust,omitempty"`
		TemplateApplicationLevel     *TemplateApplicationLevel            `json:"templateApplicationLevel,omitempty"`
		Id                           *string                              `json:"id,omitempty"`
		ODataId                      *string                              `json:"@odata.id,omitempty"`
		ODataType                    *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AutomaticUserConsentSettings = decoded.AutomaticUserConsentSettings
	s.InboundTrust = decoded.InboundTrust
	s.TemplateApplicationLevel = decoded.TemplateApplicationLevel
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MultiTenantOrganizationPartnerConfigurationTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["b2bCollaborationInbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bCollaborationInbound' for 'MultiTenantOrganizationPartnerConfigurationTemplate': %+v", err)
		}
		s.B2bCollaborationInbound = impl
	}

	if v, ok := temp["b2bCollaborationOutbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bCollaborationOutbound' for 'MultiTenantOrganizationPartnerConfigurationTemplate': %+v", err)
		}
		s.B2bCollaborationOutbound = impl
	}

	if v, ok := temp["b2bDirectConnectInbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bDirectConnectInbound' for 'MultiTenantOrganizationPartnerConfigurationTemplate': %+v", err)
		}
		s.B2bDirectConnectInbound = impl
	}

	if v, ok := temp["b2bDirectConnectOutbound"]; ok {
		impl, err := UnmarshalCrossTenantAccessPolicyB2BSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'B2bDirectConnectOutbound' for 'MultiTenantOrganizationPartnerConfigurationTemplate': %+v", err)
		}
		s.B2bDirectConnectOutbound = impl
	}

	return nil
}
