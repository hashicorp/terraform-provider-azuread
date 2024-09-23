package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationSynchronizationDataProvider = EducationPowerSchoolDataProvider{}

type EducationPowerSchoolDataProvider struct {
	// Indicates whether the source has multiple identifiers for a single student or teacher.
	AllowTeachersInMultipleSchools nullable.Type[bool] `json:"allowTeachersInMultipleSchools,omitempty"`

	// The client ID used to connect to PowerSchool.
	ClientId *string `json:"clientId,omitempty"`

	// The client secret to authenticate the connection to the PowerSchool instance.
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// The connection URL to the PowerSchool instance.
	ConnectionUrl *string `json:"connectionUrl,omitempty"`

	// Optional customization to be applied to the synchronization profile.
	Customizations *EducationSynchronizationCustomizations `json:"customizations,omitempty"`

	// The school year to sync.
	SchoolYear nullable.Type[string] `json:"schoolYear,omitempty"`

	// The list of schools to sync.
	SchoolsIds *[]string `json:"schoolsIds,omitempty"`

	// Fields inherited from EducationSynchronizationDataProvider

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationPowerSchoolDataProvider) EducationSynchronizationDataProvider() BaseEducationSynchronizationDataProviderImpl {
	return BaseEducationSynchronizationDataProviderImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationPowerSchoolDataProvider{}

func (s EducationPowerSchoolDataProvider) MarshalJSON() ([]byte, error) {
	type wrapper EducationPowerSchoolDataProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationPowerSchoolDataProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationPowerSchoolDataProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationPowerSchoolDataProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationPowerSchoolDataProvider: %+v", err)
	}

	return encoded, nil
}
