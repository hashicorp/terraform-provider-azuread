package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosSingleSignOnExtension = IosRedirectSingleSignOnExtension{}

type IosRedirectSingleSignOnExtension struct {
	// Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain
	// a maximum of 500 elements.
	Configurations *[]KeyTypedValuePair `json:"configurations,omitempty"`

	// Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
	ExtensionIdentifier *string `json:"extensionIdentifier,omitempty"`

	// Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
	TeamIdentifier nullable.Type[string] `json:"teamIdentifier,omitempty"`

	// One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must
	// begin with http:// or https://. All URL prefixes must be unique for all profiles.
	UrlPrefixes *[]string `json:"urlPrefixes,omitempty"`

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosRedirectSingleSignOnExtension) IosSingleSignOnExtension() BaseIosSingleSignOnExtensionImpl {
	return BaseIosSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s IosRedirectSingleSignOnExtension) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosRedirectSingleSignOnExtension{}

func (s IosRedirectSingleSignOnExtension) MarshalJSON() ([]byte, error) {
	type wrapper IosRedirectSingleSignOnExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosRedirectSingleSignOnExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosRedirectSingleSignOnExtension: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosRedirectSingleSignOnExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosRedirectSingleSignOnExtension: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosRedirectSingleSignOnExtension{}

func (s *IosRedirectSingleSignOnExtension) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExtensionIdentifier *string               `json:"extensionIdentifier,omitempty"`
		TeamIdentifier      nullable.Type[string] `json:"teamIdentifier,omitempty"`
		UrlPrefixes         *[]string             `json:"urlPrefixes,omitempty"`
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExtensionIdentifier = decoded.ExtensionIdentifier
	s.TeamIdentifier = decoded.TeamIdentifier
	s.UrlPrefixes = decoded.UrlPrefixes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosRedirectSingleSignOnExtension into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["configurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Configurations into list []json.RawMessage: %+v", err)
		}

		output := make([]KeyTypedValuePair, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalKeyTypedValuePairImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Configurations' for 'IosRedirectSingleSignOnExtension': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Configurations = &output
	}

	return nil
}
