package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityUserFlow = B2cIdentityUserFlow{}

type B2cIdentityUserFlow struct {
	// Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this
	// object using Get userFlowApiConnectorConfiguration.
	ApiConnectorConfiguration *UserFlowApiConnectorConfiguration `json:"apiConnectorConfiguration,omitempty"`

	// Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the
	// request. This field is RFC 5646 compliant.
	DefaultLanguageTag nullable.Type[string] `json:"defaultLanguageTag,omitempty"`

	// The identity providers included in the user flow.
	IdentityProviders *[]IdentityProvider `json:"identityProviders,omitempty"`

	// The property that determines whether language customization is enabled within the B2C user flow. Language
	// customization is not enabled by default for B2C user flows.
	IsLanguageCustomizationEnabled *bool `json:"isLanguageCustomizationEnabled,omitempty"`

	// The languages supported for customization within the user flow. Language customization is not enabled by default in
	// B2C user flows.
	Languages *[]UserFlowLanguageConfiguration `json:"languages,omitempty"`

	// The user attribute assignments included in the user flow.
	UserAttributeAssignments *[]IdentityUserFlowAttributeAssignment `json:"userAttributeAssignments,omitempty"`

	// The identity providers included in the user flow.
	UserFlowIdentityProviders *[]IdentityProviderBase `json:"userFlowIdentityProviders,omitempty"`

	// Fields inherited from IdentityUserFlow

	UserFlowType *UserFlowType `json:"userFlowType,omitempty"`

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

func (s B2cIdentityUserFlow) IdentityUserFlow() BaseIdentityUserFlowImpl {
	return BaseIdentityUserFlowImpl{
		UserFlowType: s.UserFlowType,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s B2cIdentityUserFlow) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = B2cIdentityUserFlow{}

func (s B2cIdentityUserFlow) MarshalJSON() ([]byte, error) {
	type wrapper B2cIdentityUserFlow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling B2cIdentityUserFlow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling B2cIdentityUserFlow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.b2cIdentityUserFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling B2cIdentityUserFlow: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &B2cIdentityUserFlow{}

func (s *B2cIdentityUserFlow) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApiConnectorConfiguration      *UserFlowApiConnectorConfiguration     `json:"apiConnectorConfiguration,omitempty"`
		DefaultLanguageTag             nullable.Type[string]                  `json:"defaultLanguageTag,omitempty"`
		IsLanguageCustomizationEnabled *bool                                  `json:"isLanguageCustomizationEnabled,omitempty"`
		Languages                      *[]UserFlowLanguageConfiguration       `json:"languages,omitempty"`
		UserAttributeAssignments       *[]IdentityUserFlowAttributeAssignment `json:"userAttributeAssignments,omitempty"`
		UserFlowType                   *UserFlowType                          `json:"userFlowType,omitempty"`
		Id                             *string                                `json:"id,omitempty"`
		ODataId                        *string                                `json:"@odata.id,omitempty"`
		ODataType                      *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApiConnectorConfiguration = decoded.ApiConnectorConfiguration
	s.DefaultLanguageTag = decoded.DefaultLanguageTag
	s.IsLanguageCustomizationEnabled = decoded.IsLanguageCustomizationEnabled
	s.Languages = decoded.Languages
	s.UserAttributeAssignments = decoded.UserAttributeAssignments
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.UserFlowType = decoded.UserFlowType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling B2cIdentityUserFlow into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityProviders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IdentityProviders into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProvider, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IdentityProviders' for 'B2cIdentityUserFlow': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IdentityProviders = &output
	}

	if v, ok := temp["userFlowIdentityProviders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling UserFlowIdentityProviders into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProviderBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'UserFlowIdentityProviders' for 'B2cIdentityUserFlow': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.UserFlowIdentityProviders = &output
	}

	return nil
}
