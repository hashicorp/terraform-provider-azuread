package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WritebackConfiguration interface {
	WritebackConfiguration() BaseWritebackConfigurationImpl
}

var _ WritebackConfiguration = BaseWritebackConfigurationImpl{}

type BaseWritebackConfigurationImpl struct {
	// Indicates whether writeback of cloud groups to on-premise Active Directory is enabled. Default value is true for
	// Microsoft 365 groups and false for security groups.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWritebackConfigurationImpl) WritebackConfiguration() BaseWritebackConfigurationImpl {
	return s
}

var _ WritebackConfiguration = RawWritebackConfigurationImpl{}

// RawWritebackConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWritebackConfigurationImpl struct {
	writebackConfiguration BaseWritebackConfigurationImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawWritebackConfigurationImpl) WritebackConfiguration() BaseWritebackConfigurationImpl {
	return s.writebackConfiguration
}

func UnmarshalWritebackConfigurationImplementation(input []byte) (WritebackConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WritebackConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupWritebackConfiguration") {
		var out GroupWritebackConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupWritebackConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseWritebackConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWritebackConfigurationImpl: %+v", err)
	}

	return RawWritebackConfigurationImpl{
		writebackConfiguration: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
