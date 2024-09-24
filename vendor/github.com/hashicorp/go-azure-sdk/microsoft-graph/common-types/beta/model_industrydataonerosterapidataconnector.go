package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataApiDataConnector = IndustryDataOneRosterApiDataConnector{}

type IndustryDataOneRosterApiDataConnector struct {
	// The API version of the OneRoster source. Example: 1.1, 1.2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Indicates whether the user specified to import optional contacts data.
	IsContactsEnabled *bool `json:"isContactsEnabled,omitempty"`

	// Indicates whether the user specified to import optional demographics data.
	IsDemographicsEnabled *bool `json:"isDemographicsEnabled,omitempty"`

	// Indicates whether the user specified to import optional flags data.
	IsFlagsEnabled *bool `json:"isFlagsEnabled,omitempty"`

	// Fields inherited from IndustryDataApiDataConnector

	ApiFormat *IndustryDataApiFormat `json:"apiFormat,omitempty"`

	// The base URL, including the scheme, host, and path for the API, with or without a trailing '/'. For example,
	// 'https://example.com/ims/oneRoster/v1p1'
	BaseUrl *string `json:"baseUrl,omitempty"`

	Credential IndustryDataCredential `json:"credential"`

	// Fields inherited from IndustryDataIndustryDataConnector

	// The name of the data connector. Maximum supported length is 100 characters.
	DisplayName *string `json:"displayName,omitempty"`

	SourceSystem *IndustryDataSourceSystemDefinition `json:"sourceSystem,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IndustryDataOneRosterApiDataConnector) IndustryDataApiDataConnector() BaseIndustryDataApiDataConnectorImpl {
	return BaseIndustryDataApiDataConnectorImpl{
		ApiFormat:    s.ApiFormat,
		BaseUrl:      s.BaseUrl,
		Credential:   s.Credential,
		DisplayName:  s.DisplayName,
		SourceSystem: s.SourceSystem,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s IndustryDataOneRosterApiDataConnector) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return BaseIndustryDataIndustryDataConnectorImpl{
		DisplayName:  s.DisplayName,
		SourceSystem: s.SourceSystem,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s IndustryDataOneRosterApiDataConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataOneRosterApiDataConnector{}

func (s IndustryDataOneRosterApiDataConnector) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataOneRosterApiDataConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataOneRosterApiDataConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataOneRosterApiDataConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.oneRosterApiDataConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataOneRosterApiDataConnector: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataOneRosterApiDataConnector{}

func (s *IndustryDataOneRosterApiDataConnector) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApiVersion            *string                             `json:"apiVersion,omitempty"`
		IsContactsEnabled     *bool                               `json:"isContactsEnabled,omitempty"`
		IsDemographicsEnabled *bool                               `json:"isDemographicsEnabled,omitempty"`
		IsFlagsEnabled        *bool                               `json:"isFlagsEnabled,omitempty"`
		ApiFormat             *IndustryDataApiFormat              `json:"apiFormat,omitempty"`
		BaseUrl               *string                             `json:"baseUrl,omitempty"`
		DisplayName           *string                             `json:"displayName,omitempty"`
		SourceSystem          *IndustryDataSourceSystemDefinition `json:"sourceSystem,omitempty"`
		Id                    *string                             `json:"id,omitempty"`
		ODataId               *string                             `json:"@odata.id,omitempty"`
		ODataType             *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApiVersion = decoded.ApiVersion
	s.IsContactsEnabled = decoded.IsContactsEnabled
	s.IsDemographicsEnabled = decoded.IsDemographicsEnabled
	s.IsFlagsEnabled = decoded.IsFlagsEnabled
	s.ApiFormat = decoded.ApiFormat
	s.BaseUrl = decoded.BaseUrl
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SourceSystem = decoded.SourceSystem

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataOneRosterApiDataConnector into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["credential"]; ok {
		impl, err := UnmarshalIndustryDataCredentialImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Credential' for 'IndustryDataOneRosterApiDataConnector': %+v", err)
		}
		s.Credential = impl
	}

	return nil
}
