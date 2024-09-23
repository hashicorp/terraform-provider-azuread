package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppManagementConfiguration = AppManagementApplicationConfiguration{}

type AppManagementApplicationConfiguration struct {

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

func (s AppManagementApplicationConfiguration) AppManagementConfiguration() BaseAppManagementConfigurationImpl {
	return BaseAppManagementConfigurationImpl{
		KeyCredentials:      s.KeyCredentials,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		PasswordCredentials: s.PasswordCredentials,
	}
}

var _ json.Marshaler = AppManagementApplicationConfiguration{}

func (s AppManagementApplicationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AppManagementApplicationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppManagementApplicationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppManagementApplicationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appManagementApplicationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppManagementApplicationConfiguration: %+v", err)
	}

	return encoded, nil
}
