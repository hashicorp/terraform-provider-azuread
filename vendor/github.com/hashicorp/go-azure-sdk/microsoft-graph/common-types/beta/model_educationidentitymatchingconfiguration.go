package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationIdentitySynchronizationConfiguration = EducationIdentityMatchingConfiguration{}

type EducationIdentityMatchingConfiguration struct {
	// Mapping between the user account and the options to use to uniquely identify the user to update.
	MatchingOptions *[]EducationIdentityMatchingOptions `json:"matchingOptions,omitempty"`

	// Fields inherited from EducationIdentitySynchronizationConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationIdentityMatchingConfiguration) EducationIdentitySynchronizationConfiguration() BaseEducationIdentitySynchronizationConfigurationImpl {
	return BaseEducationIdentitySynchronizationConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationIdentityMatchingConfiguration{}

func (s EducationIdentityMatchingConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper EducationIdentityMatchingConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationIdentityMatchingConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationIdentityMatchingConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationIdentityMatchingConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationIdentityMatchingConfiguration: %+v", err)
	}

	return encoded, nil
}
