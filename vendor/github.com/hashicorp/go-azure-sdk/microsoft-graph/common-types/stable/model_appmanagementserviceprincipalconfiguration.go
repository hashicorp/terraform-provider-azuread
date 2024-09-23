package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppManagementConfiguration = AppManagementServicePrincipalConfiguration{}

type AppManagementServicePrincipalConfiguration struct {

	// Fields inherited from AppManagementConfiguration

	// Collection of keyCredential restrictions settings to be applied to an application or service principal.
	KeyCredentials *[]KeyCredentialConfiguration `json:"keyCredentials,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of password restrictions settings to be applied to an application or service principal.
	PasswordCredentials *[]PasswordCredentialConfiguration `json:"passwordCredentials,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AppManagementServicePrincipalConfiguration) AppManagementConfiguration() BaseAppManagementConfigurationImpl {
	return BaseAppManagementConfigurationImpl{
		KeyCredentials:      s.KeyCredentials,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		PasswordCredentials: s.PasswordCredentials,
	}
}

var _ json.Marshaler = AppManagementServicePrincipalConfiguration{}

func (s AppManagementServicePrincipalConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AppManagementServicePrincipalConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppManagementServicePrincipalConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppManagementServicePrincipalConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appManagementServicePrincipalConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppManagementServicePrincipalConfiguration: %+v", err)
	}

	return encoded, nil
}
