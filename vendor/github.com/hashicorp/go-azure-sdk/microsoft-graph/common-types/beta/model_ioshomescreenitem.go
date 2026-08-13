package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosHomeScreenItem interface {
	IosHomeScreenItem() BaseIosHomeScreenItemImpl
}

var _ IosHomeScreenItem = BaseIosHomeScreenItemImpl{}

type BaseIosHomeScreenItemImpl struct {
	// Name of the app
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIosHomeScreenItemImpl) IosHomeScreenItem() BaseIosHomeScreenItemImpl {
	return s
}

var _ IosHomeScreenItem = RawIosHomeScreenItemImpl{}

// RawIosHomeScreenItemImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIosHomeScreenItemImpl struct {
	iosHomeScreenItem BaseIosHomeScreenItemImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawIosHomeScreenItemImpl) IosHomeScreenItem() BaseIosHomeScreenItemImpl {
	return s.iosHomeScreenItem
}

func (s RawIosHomeScreenItemImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIosHomeScreenItemImplementation(input []byte) (IosHomeScreenItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IosHomeScreenItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosHomeScreenApp") {
		var out IosHomeScreenApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosHomeScreenApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosHomeScreenFolder") {
		var out IosHomeScreenFolder
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosHomeScreenFolder: %+v", err)
		}
		return out, nil
	}

	var parent BaseIosHomeScreenItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIosHomeScreenItemImpl: %+v", err)
	}

	return RawIosHomeScreenItemImpl{
		iosHomeScreenItem: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
