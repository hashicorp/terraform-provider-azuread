package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAssociation interface {
	NetworkaccessAssociation() BaseNetworkaccessAssociationImpl
}

var _ NetworkaccessAssociation = BaseNetworkaccessAssociationImpl{}

type BaseNetworkaccessAssociationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseNetworkaccessAssociationImpl) NetworkaccessAssociation() BaseNetworkaccessAssociationImpl {
	return s
}

var _ NetworkaccessAssociation = RawNetworkaccessAssociationImpl{}

// RawNetworkaccessAssociationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawNetworkaccessAssociationImpl struct {
	networkaccessAssociation BaseNetworkaccessAssociationImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawNetworkaccessAssociationImpl) NetworkaccessAssociation() BaseNetworkaccessAssociationImpl {
	return s.networkaccessAssociation
}

func (s RawNetworkaccessAssociationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalNetworkaccessAssociationImplementation(input []byte) (NetworkaccessAssociation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessAssociation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.associatedBranch") {
		var out NetworkaccessAssociatedBranch
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessAssociatedBranch: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessAssociationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessAssociationImpl: %+v", err)
	}

	return RawNetworkaccessAssociationImpl{
		networkaccessAssociation: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
