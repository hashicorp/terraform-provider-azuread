package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataApiDataConnector interface {
	Entity
	IndustryDataIndustryDataConnector
	IndustryDataApiDataConnector() BaseIndustryDataApiDataConnectorImpl
}

var _ IndustryDataApiDataConnector = BaseIndustryDataApiDataConnectorImpl{}

type BaseIndustryDataApiDataConnectorImpl struct {
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

func (s BaseIndustryDataApiDataConnectorImpl) IndustryDataApiDataConnector() BaseIndustryDataApiDataConnectorImpl {
	return s
}

func (s BaseIndustryDataApiDataConnectorImpl) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return BaseIndustryDataIndustryDataConnectorImpl{
		DisplayName:  s.DisplayName,
		SourceSystem: s.SourceSystem,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s BaseIndustryDataApiDataConnectorImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataApiDataConnector = RawIndustryDataApiDataConnectorImpl{}

// RawIndustryDataApiDataConnectorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataApiDataConnectorImpl struct {
	industryDataApiDataConnector BaseIndustryDataApiDataConnectorImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawIndustryDataApiDataConnectorImpl) IndustryDataApiDataConnector() BaseIndustryDataApiDataConnectorImpl {
	return s.industryDataApiDataConnector
}

func (s RawIndustryDataApiDataConnectorImpl) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return s.industryDataApiDataConnector.IndustryDataIndustryDataConnector()
}

func (s RawIndustryDataApiDataConnectorImpl) Entity() BaseEntityImpl {
	return s.industryDataApiDataConnector.Entity()
}

var _ json.Marshaler = BaseIndustryDataApiDataConnectorImpl{}

func (s BaseIndustryDataApiDataConnectorImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataApiDataConnectorImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataApiDataConnectorImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataApiDataConnectorImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.apiDataConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataApiDataConnectorImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseIndustryDataApiDataConnectorImpl{}

func (s *BaseIndustryDataApiDataConnectorImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApiFormat    *IndustryDataApiFormat              `json:"apiFormat,omitempty"`
		BaseUrl      *string                             `json:"baseUrl,omitempty"`
		DisplayName  *string                             `json:"displayName,omitempty"`
		SourceSystem *IndustryDataSourceSystemDefinition `json:"sourceSystem,omitempty"`
		Id           *string                             `json:"id,omitempty"`
		ODataId      *string                             `json:"@odata.id,omitempty"`
		ODataType    *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApiFormat = decoded.ApiFormat
	s.BaseUrl = decoded.BaseUrl
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SourceSystem = decoded.SourceSystem

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseIndustryDataApiDataConnectorImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["credential"]; ok {
		impl, err := UnmarshalIndustryDataCredentialImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Credential' for 'BaseIndustryDataApiDataConnectorImpl': %+v", err)
		}
		s.Credential = impl
	}

	return nil
}

func UnmarshalIndustryDataApiDataConnectorImplementation(input []byte) (IndustryDataApiDataConnector, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataApiDataConnector into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.oneRosterApiDataConnector") {
		var out IndustryDataOneRosterApiDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOneRosterApiDataConnector: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataApiDataConnectorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataApiDataConnectorImpl: %+v", err)
	}

	return RawIndustryDataApiDataConnectorImpl{
		industryDataApiDataConnector: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
