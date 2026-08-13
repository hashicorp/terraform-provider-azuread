package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
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

// RawEntitlementsDataCollectionInfoImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEntitlementsDataCollectionInfoImpl struct {
	entitlementsDataCollectionInfo BaseEntitlementsDataCollectionInfoImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawEntitlementsDataCollectionInfoImpl) EntitlementsDataCollectionInfo() BaseEntitlementsDataCollectionInfoImpl {
	return s.entitlementsDataCollectionInfo
}

func (s RawEntitlementsDataCollectionInfoImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
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
