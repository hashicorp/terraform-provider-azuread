package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IndustryDataSourceSystemDefinition{}

type IndustryDataSourceSystemDefinition struct {
	// The name of the source system. Maximum supported length is 100 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// A collection of user matching settings by roleGroup.
	UserMatchingSettings *[]IndustryDataUserMatchingSetting `json:"userMatchingSettings,omitempty"`

	// The name of the vendor who supplies the source system. Maximum supported length is 100 characters.
	Vendor nullable.Type[string] `json:"vendor,omitempty"`

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

func (s IndustryDataSourceSystemDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataSourceSystemDefinition{}

func (s IndustryDataSourceSystemDefinition) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataSourceSystemDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataSourceSystemDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataSourceSystemDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.sourceSystemDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataSourceSystemDefinition: %+v", err)
	}

	return encoded, nil
}
