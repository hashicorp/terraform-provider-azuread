package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsUrlToItemResolverBase interface {
	ExternalConnectorsUrlToItemResolverBase() BaseExternalConnectorsUrlToItemResolverBaseImpl
}

var _ ExternalConnectorsUrlToItemResolverBase = BaseExternalConnectorsUrlToItemResolverBaseImpl{}

type BaseExternalConnectorsUrlToItemResolverBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The priority which defines the sequence in which the urlToItemResolverBase instances are evaluated.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseExternalConnectorsUrlToItemResolverBaseImpl) ExternalConnectorsUrlToItemResolverBase() BaseExternalConnectorsUrlToItemResolverBaseImpl {
	return s
}

var _ ExternalConnectorsUrlToItemResolverBase = RawExternalConnectorsUrlToItemResolverBaseImpl{}

// RawExternalConnectorsUrlToItemResolverBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawExternalConnectorsUrlToItemResolverBaseImpl struct {
	externalConnectorsUrlToItemResolverBase BaseExternalConnectorsUrlToItemResolverBaseImpl
	Type                                    string
	Values                                  map[string]interface{}
}

func (s RawExternalConnectorsUrlToItemResolverBaseImpl) ExternalConnectorsUrlToItemResolverBase() BaseExternalConnectorsUrlToItemResolverBaseImpl {
	return s.externalConnectorsUrlToItemResolverBase
}

func (s RawExternalConnectorsUrlToItemResolverBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalExternalConnectorsUrlToItemResolverBaseImplementation(input []byte) (ExternalConnectorsUrlToItemResolverBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsUrlToItemResolverBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.itemIdResolver") {
		var out ExternalConnectorsItemIdResolver
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsItemIdResolver: %+v", err)
		}
		return out, nil
	}

	var parent BaseExternalConnectorsUrlToItemResolverBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExternalConnectorsUrlToItemResolverBaseImpl: %+v", err)
	}

	return RawExternalConnectorsUrlToItemResolverBaseImpl{
		externalConnectorsUrlToItemResolverBase: parent,
		Type:                                    value,
		Values:                                  temp,
	}, nil

}
