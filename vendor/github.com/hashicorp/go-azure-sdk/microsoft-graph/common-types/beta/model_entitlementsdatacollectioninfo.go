package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementsDataCollectionInfo interface {
	EntitlementsDataCollectionInfo() BaseEntitlementsDataCollectionInfoImpl
}

var _ EntitlementsDataCollectionInfo = BaseEntitlementsDataCollectionInfoImpl{}

type BaseEntitlementsDataCollectionInfoImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEntitlementsDataCollectionInfoImpl) EntitlementsDataCollectionInfo() BaseEntitlementsDataCollectionInfoImpl {
	return s
}

var _ EntitlementsDataCollectionInfo = RawEntitlementsDataCollectionInfoImpl{}

// RawEntitlementsDataCollectionInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntitlementsDataCollectionInfoImpl struct {
	entitlementsDataCollectionInfo BaseEntitlementsDataCollectionInfoImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawEntitlementsDataCollectionInfoImpl) EntitlementsDataCollectionInfo() BaseEntitlementsDataCollectionInfoImpl {
	return s.entitlementsDataCollectionInfo
}

func UnmarshalEntitlementsDataCollectionInfoImplementation(input []byte) (EntitlementsDataCollectionInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EntitlementsDataCollectionInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.entitlementsDataCollection") {
		var out EntitlementsDataCollection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EntitlementsDataCollection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noEntitlementsDataCollection") {
		var out NoEntitlementsDataCollection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoEntitlementsDataCollection: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntitlementsDataCollectionInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntitlementsDataCollectionInfoImpl: %+v", err)
	}

	return RawEntitlementsDataCollectionInfoImpl{
		entitlementsDataCollectionInfo: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
