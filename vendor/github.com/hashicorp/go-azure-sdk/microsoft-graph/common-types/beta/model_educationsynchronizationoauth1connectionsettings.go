package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationSynchronizationConnectionSettings = EducationSynchronizationOAuth1ConnectionSettings{}

type EducationSynchronizationOAuth1ConnectionSettings struct {

	// Fields inherited from EducationSynchronizationConnectionSettings

	// Client ID used to connect to the provider.
	ClientId *string `json:"clientId,omitempty"`

	// Client secret to authenticate the connection to the provider.
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationSynchronizationOAuth1ConnectionSettings) EducationSynchronizationConnectionSettings() BaseEducationSynchronizationConnectionSettingsImpl {
	return BaseEducationSynchronizationConnectionSettingsImpl{
		ClientId:     s.ClientId,
		ClientSecret: s.ClientSecret,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

var _ json.Marshaler = EducationSynchronizationOAuth1ConnectionSettings{}

func (s EducationSynchronizationOAuth1ConnectionSettings) MarshalJSON() ([]byte, error) {
	type wrapper EducationSynchronizationOAuth1ConnectionSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSynchronizationOAuth1ConnectionSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationOAuth1ConnectionSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSynchronizationOAuth1ConnectionSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSynchronizationOAuth1ConnectionSettings: %+v", err)
	}

	return encoded, nil
}
