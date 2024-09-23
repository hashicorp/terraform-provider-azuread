package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosSingleSignOnExtension = IosAzureAdSingleSignOnExtension{}

type IosAzureAdSingleSignOnExtension struct {
	// An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
	BundleIdAccessControlList *[]string `json:"bundleIdAccessControlList,omitempty"`

	// Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain
	// a maximum of 500 elements.
	Configurations *[]KeyTypedValuePair `json:"configurations,omitempty"`

	// Enables or disables shared device mode.
	EnableSharedDeviceMode *bool `json:"enableSharedDeviceMode,omitempty"`

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosAzureAdSingleSignOnExtension) IosSingleSignOnExtension() BaseIosSingleSignOnExtensionImpl {
	return BaseIosSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s IosAzureAdSingleSignOnExtension) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosAzureAdSingleSignOnExtension{}

func (s IosAzureAdSingleSignOnExtension) MarshalJSON() ([]byte, error) {
	type wrapper IosAzureAdSingleSignOnExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosAzureAdSingleSignOnExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosAzureAdSingleSignOnExtension: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosAzureAdSingleSignOnExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosAzureAdSingleSignOnExtension: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosAzureAdSingleSignOnExtension{}

func (s *IosAzureAdSingleSignOnExtension) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BundleIdAccessControlList *[]string `json:"bundleIdAccessControlList,omitempty"`
		EnableSharedDeviceMode    *bool     `json:"enableSharedDeviceMode,omitempty"`
		ODataId                   *string   `json:"@odata.id,omitempty"`
		ODataType                 *string   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BundleIdAccessControlList = decoded.BundleIdAccessControlList
	s.EnableSharedDeviceMode = decoded.EnableSharedDeviceMode
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosAzureAdSingleSignOnExtension into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Configurations' for 'IosAzureAdSingleSignOnExtension': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Configurations = &output
	}

	return nil
}
