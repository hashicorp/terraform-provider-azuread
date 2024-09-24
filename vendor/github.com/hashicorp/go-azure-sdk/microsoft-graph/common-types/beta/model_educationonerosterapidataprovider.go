package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationSynchronizationDataProvider = EducationOneRosterApiDataProvider{}

type EducationOneRosterApiDataProvider struct {
	ConnectionSettings EducationSynchronizationConnectionSettings `json:"connectionSettings"`

	// The connection URL to the OneRoster instance.
	ConnectionUrl *string `json:"connectionUrl,omitempty"`

	// Optional customization to be applied to the synchronization profile.
	Customizations *EducationSynchronizationCustomizations `json:"customizations,omitempty"`

	// The OneRoster Service Provider name as defined by the OneRoster specification.
	ProviderName nullable.Type[string] `json:"providerName,omitempty"`

	// The list of School/Org sourcedId to sync.
	SchoolsIds *[]string `json:"schoolsIds,omitempty"`

	// The list of academic sessions to sync.
	TermIds *[]string `json:"termIds,omitempty"`

	// Fields inherited from EducationSynchronizationDataProvider

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationOneRosterApiDataProvider) EducationSynchronizationDataProvider() BaseEducationSynchronizationDataProviderImpl {
	return BaseEducationSynchronizationDataProviderImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationOneRosterApiDataProvider{}

func (s EducationOneRosterApiDataProvider) MarshalJSON() ([]byte, error) {
	type wrapper EducationOneRosterApiDataProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationOneRosterApiDataProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationOneRosterApiDataProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationOneRosterApiDataProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationOneRosterApiDataProvider: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationOneRosterApiDataProvider{}

func (s *EducationOneRosterApiDataProvider) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ConnectionUrl  *string                                 `json:"connectionUrl,omitempty"`
		Customizations *EducationSynchronizationCustomizations `json:"customizations,omitempty"`
		ProviderName   nullable.Type[string]                   `json:"providerName,omitempty"`
		SchoolsIds     *[]string                               `json:"schoolsIds,omitempty"`
		TermIds        *[]string                               `json:"termIds,omitempty"`
		ODataId        *string                                 `json:"@odata.id,omitempty"`
		ODataType      *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ConnectionUrl = decoded.ConnectionUrl
	s.Customizations = decoded.Customizations
	s.ProviderName = decoded.ProviderName
	s.SchoolsIds = decoded.SchoolsIds
	s.TermIds = decoded.TermIds
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationOneRosterApiDataProvider into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["connectionSettings"]; ok {
		impl, err := UnmarshalEducationSynchronizationConnectionSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ConnectionSettings' for 'EducationOneRosterApiDataProvider': %+v", err)
		}
		s.ConnectionSettings = impl
	}

	return nil
}
