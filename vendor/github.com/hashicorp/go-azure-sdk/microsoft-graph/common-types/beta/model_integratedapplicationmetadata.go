package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegratedApplicationMetadata interface {
	IntegratedApplicationMetadata() BaseIntegratedApplicationMetadataImpl
}

var _ IntegratedApplicationMetadata = BaseIntegratedApplicationMetadataImpl{}

type BaseIntegratedApplicationMetadataImpl struct {
	// The name of the integrated application.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The version number of the integrated application.
	Version nullable.Type[string] `json:"version,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIntegratedApplicationMetadataImpl) IntegratedApplicationMetadata() BaseIntegratedApplicationMetadataImpl {
	return s
}

var _ IntegratedApplicationMetadata = RawIntegratedApplicationMetadataImpl{}

// RawIntegratedApplicationMetadataImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIntegratedApplicationMetadataImpl struct {
	integratedApplicationMetadata BaseIntegratedApplicationMetadataImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawIntegratedApplicationMetadataImpl) IntegratedApplicationMetadata() BaseIntegratedApplicationMetadataImpl {
	return s.integratedApplicationMetadata
}

func (s RawIntegratedApplicationMetadataImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIntegratedApplicationMetadataImplementation(input []byte) (IntegratedApplicationMetadata, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IntegratedApplicationMetadata into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.protectedApplicationMetadata") {
		var out ProtectedApplicationMetadata
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectedApplicationMetadata: %+v", err)
		}
		return out, nil
	}

	var parent BaseIntegratedApplicationMetadataImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIntegratedApplicationMetadataImpl: %+v", err)
	}

	return RawIntegratedApplicationMetadataImpl{
		integratedApplicationMetadata: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
