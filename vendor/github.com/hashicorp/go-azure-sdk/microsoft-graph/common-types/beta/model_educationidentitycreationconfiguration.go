package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationIdentitySynchronizationConfiguration = EducationIdentityCreationConfiguration{}

type EducationIdentityCreationConfiguration struct {
	UserDomains *[]EducationIdentityDomain `json:"userDomains,omitempty"`

	// Fields inherited from EducationIdentitySynchronizationConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationIdentityCreationConfiguration) EducationIdentitySynchronizationConfiguration() BaseEducationIdentitySynchronizationConfigurationImpl {
	return BaseEducationIdentitySynchronizationConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationIdentityCreationConfiguration{}

func (s EducationIdentityCreationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper EducationIdentityCreationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationIdentityCreationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationIdentityCreationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationIdentityCreationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationIdentityCreationConfiguration: %+v", err)
	}

	return encoded, nil
}
