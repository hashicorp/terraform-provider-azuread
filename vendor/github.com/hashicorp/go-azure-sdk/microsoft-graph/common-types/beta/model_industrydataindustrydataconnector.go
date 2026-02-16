package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataConnector interface {
	Entity
	IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl
}

var _ IndustryDataIndustryDataConnector = BaseIndustryDataIndustryDataConnectorImpl{}

type BaseIndustryDataIndustryDataConnectorImpl struct {
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

func (s BaseIndustryDataIndustryDataConnectorImpl) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return s
}

func (s BaseIndustryDataIndustryDataConnectorImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataIndustryDataConnector = RawIndustryDataIndustryDataConnectorImpl{}

// RawIndustryDataIndustryDataConnectorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataIndustryDataConnectorImpl struct {
	industryDataIndustryDataConnector BaseIndustryDataIndustryDataConnectorImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawIndustryDataIndustryDataConnectorImpl) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return s.industryDataIndustryDataConnector
}

func (s RawIndustryDataIndustryDataConnectorImpl) Entity() BaseEntityImpl {
	return s.industryDataIndustryDataConnector.Entity()
}

var _ json.Marshaler = BaseIndustryDataIndustryDataConnectorImpl{}

func (s BaseIndustryDataIndustryDataConnectorImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataIndustryDataConnectorImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataIndustryDataConnectorImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataIndustryDataConnectorImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.industryDataConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataIndustryDataConnectorImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataIndustryDataConnectorImplementation(input []byte) (IndustryDataIndustryDataConnector, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataConnector into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.apiDataConnector") {
		var out IndustryDataApiDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataApiDataConnector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.fileDataConnector") {
		var out IndustryDataFileDataConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataFileDataConnector: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataIndustryDataConnectorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataIndustryDataConnectorImpl: %+v", err)
	}

	return RawIndustryDataIndustryDataConnectorImpl{
		industryDataIndustryDataConnector: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
