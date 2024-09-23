package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataFileDataConnector = IndustryDataAzureDataLakeConnector{}

type IndustryDataAzureDataLakeConnector struct {
	// The file format that external systems can upload using this connector.
	FileFormat *IndustryDataFileFormatReferenceValue `json:"fileFormat,omitempty"`

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

func (s IndustryDataAzureDataLakeConnector) IndustryDataFileDataConnector() BaseIndustryDataFileDataConnectorImpl {
	return BaseIndustryDataFileDataConnectorImpl{
		DisplayName:  s.DisplayName,
		SourceSystem: s.SourceSystem,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s IndustryDataAzureDataLakeConnector) IndustryDataIndustryDataConnector() BaseIndustryDataIndustryDataConnectorImpl {
	return BaseIndustryDataIndustryDataConnectorImpl{
		DisplayName:  s.DisplayName,
		SourceSystem: s.SourceSystem,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s IndustryDataAzureDataLakeConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataAzureDataLakeConnector{}

func (s IndustryDataAzureDataLakeConnector) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataAzureDataLakeConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataAzureDataLakeConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataAzureDataLakeConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.azureDataLakeConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataAzureDataLakeConnector: %+v", err)
	}

	return encoded, nil
}
