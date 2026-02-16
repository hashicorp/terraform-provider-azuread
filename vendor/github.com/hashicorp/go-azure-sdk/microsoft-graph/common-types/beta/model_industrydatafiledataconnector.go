package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFileDataConnector interface {
	Entity
	IndustryDataIndustryDataConnector
	IndustryDataFileDataConnector() BaseIndustryDataFileDataConnectorImpl
}

var _ IndustryDataFileDataConnector = BaseIndustryDataFileDataConnectorImpl{}

type BaseIndustryDataFileDataConnectorImpl struct {

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

func (s BaseIndustryDataFileDataConnectorImpl) IndustryDataFileDataConnector() BaseIndustryDataFileDataConnectorImpl {
	return s
}

func (s BaseIndustryDataFileDataConnectorImpl) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return BaseIndustryDataIndustryDataConnectorImpl{
		DisplayName:  s.DisplayName,
		SourceSystem: s.SourceSystem,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s BaseIndustryDataFileDataConnectorImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataFileDataConnector = RawIndustryDataFileDataConnectorImpl{}

// RawIndustryDataFileDataConnectorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataFileDataConnectorImpl struct {
	industryDataFileDataConnector BaseIndustryDataFileDataConnectorImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawIndustryDataFileDataConnectorImpl) IndustryDataFileDataConnector() BaseIndustryDataFileDataConnectorImpl {
	return s.industryDataFileDataConnector
}

func (s RawIndustryDataFileDataConnectorImpl) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return s.industryDataFileDataConnector.IndustryDataIndustryDataConnector()
}

func (s RawIndustryDataFileDataConnectorImpl) Entity() BaseEntityImpl {
	return s.industryDataFileDataConnector.Entity()
}

var _ json.Marshaler = BaseIndustryDataFileDataConnectorImpl{}

func (s BaseIndustryDataFileDataConnectorImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataFileDataConnectorImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataFileDataConnectorImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataFileDataConnectorImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.fileDataConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataFileDataConnectorImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataFileDataConnectorImplementation(input []byte) (IndustryDataFileDataConnector, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataFileDataConnector into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.azureDataLakeConnector") {
		var out IndustryDataAzureDataLakeConnector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataAzureDataLakeConnector: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataFileDataConnectorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataFileDataConnectorImpl: %+v", err)
	}

	return RawIndustryDataFileDataConnectorImpl{
		industryDataFileDataConnector: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
