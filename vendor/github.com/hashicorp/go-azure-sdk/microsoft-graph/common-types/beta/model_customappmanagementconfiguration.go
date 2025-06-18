package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppManagementConfiguration = CustomAppManagementConfiguration{}

type CustomAppManagementConfiguration struct {
	// Restrictions that are applicable only to application objects to which the policy is attached.
	ApplicationRestrictions *CustomAppManagementApplicationConfiguration `json:"applicationRestrictions,omitempty"`

	// Fields inherited from AppManagementConfiguration

	KeyCredentials *[]KeyCredentialConfiguration `json:"keyCredentials,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PasswordCredentials *[]PasswordCredentialConfiguration `json:"passwordCredentials,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CustomAppManagementConfiguration) AppManagementConfiguration() BaseAppManagementConfigurationImpl {
	return BaseAppManagementConfigurationImpl{
		KeyCredentials:      s.KeyCredentials,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		PasswordCredentials: s.PasswordCredentials,
	}
}

var _ json.Marshaler = CustomAppManagementConfiguration{}

func (s CustomAppManagementConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CustomAppManagementConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomAppManagementConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomAppManagementConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customAppManagementConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomAppManagementConfiguration: %+v", err)
	}

	return encoded, nil
}
