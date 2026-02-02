package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityContainer{}

type IdentityContainer struct {
	// Represents entry point for API connectors.
	ApiConnectors *[]IdentityApiConnector `json:"apiConnectors,omitempty"`

	// Represents listeners for custom authentication extension events in Azure AD for workforce and customers.
	AuthenticationEventListeners *[]AuthenticationEventListener `json:"authenticationEventListeners,omitempty"`

	// Represents the entry point for self-service sign-up and sign-in user flows in both Microsoft Entra workforce and
	// external tenants.
	AuthenticationEventsFlows *[]AuthenticationEventsFlow `json:"authenticationEventsFlows,omitempty"`

	// Represents entry point for B2X/self-service sign-up identity userflows.
	B2xUserFlows *[]B2xIdentityUserFlow `json:"b2xUserFlows,omitempty"`

	// the entry point for the Conditional Access (CA) object model.
	ConditionalAccess *ConditionalAccessRoot `json:"conditionalAccess,omitempty"`

	// Represents custom extensions to authentication flows in Azure AD for workforce and customers.
	CustomAuthenticationExtensions *[]CustomAuthenticationExtension `json:"customAuthenticationExtensions,omitempty"`

	IdentityProviders *[]IdentityProviderBase `json:"identityProviders,omitempty"`

	// Represents entry point for identity userflow attributes.
	UserFlowAttributes *[]IdentityUserFlowAttribute `json:"userFlowAttributes,omitempty"`

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

func (s IdentityContainer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityContainer{}

func (s IdentityContainer) MarshalJSON() ([]byte, error) {
	type wrapper IdentityContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityContainer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityContainer: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityContainer{}

func (s *IdentityContainer) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApiConnectors     *[]IdentityApiConnector `json:"apiConnectors,omitempty"`
		B2xUserFlows      *[]B2xIdentityUserFlow  `json:"b2xUserFlows,omitempty"`
		ConditionalAccess *ConditionalAccessRoot  `json:"conditionalAccess,omitempty"`
		Id                *string                 `json:"id,omitempty"`
		ODataId           *string                 `json:"@odata.id,omitempty"`
		ODataType         *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApiConnectors = decoded.ApiConnectors
	s.B2xUserFlows = decoded.B2xUserFlows
	s.ConditionalAccess = decoded.ConditionalAccess
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IdentityContainer into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationEventListeners"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AuthenticationEventListeners into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationEventListener, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationEventListenerImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AuthenticationEventListeners' for 'IdentityContainer': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AuthenticationEventListeners = &output
	}

	if v, ok := temp["authenticationEventsFlows"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AuthenticationEventsFlows into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationEventsFlow, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationEventsFlowImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AuthenticationEventsFlows' for 'IdentityContainer': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AuthenticationEventsFlows = &output
	}

	if v, ok := temp["customAuthenticationExtensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CustomAuthenticationExtensions into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomAuthenticationExtension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCustomAuthenticationExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CustomAuthenticationExtensions' for 'IdentityContainer': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CustomAuthenticationExtensions = &output
	}

	if v, ok := temp["identityProviders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IdentityProviders into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProviderBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IdentityProviders' for 'IdentityContainer': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IdentityProviders = &output
	}

	if v, ok := temp["userFlowAttributes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling UserFlowAttributes into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityUserFlowAttribute, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityUserFlowAttributeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'UserFlowAttributes' for 'IdentityContainer': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.UserFlowAttributes = &output
	}

	return nil
}
