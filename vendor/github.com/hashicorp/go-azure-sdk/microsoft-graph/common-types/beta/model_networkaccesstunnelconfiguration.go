package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfiguration interface {
	NetworkaccessTunnelConfiguration() BaseNetworkaccessTunnelConfigurationImpl
}

var _ NetworkaccessTunnelConfiguration = BaseNetworkaccessTunnelConfigurationImpl{}

type BaseNetworkaccessTunnelConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A key to establish secure connection between the link and VPN tunnel on the edge.
	PreSharedKey nullable.Type[string] `json:"preSharedKey,omitempty"`

	// Another key for zone redundant tunnel. Required only when you select zoneRedundancy redindancyTier when creating a
	// deviceLink.
	ZoneRedundancyPreSharedKey nullable.Type[string] `json:"zoneRedundancyPreSharedKey,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseNetworkaccessTunnelConfigurationImpl) NetworkaccessTunnelConfiguration() BaseNetworkaccessTunnelConfigurationImpl {
	return s
}

var _ NetworkaccessTunnelConfiguration = RawNetworkaccessTunnelConfigurationImpl{}

// RawNetworkaccessTunnelConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessTunnelConfigurationImpl struct {
	networkaccessTunnelConfiguration BaseNetworkaccessTunnelConfigurationImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawNetworkaccessTunnelConfigurationImpl) NetworkaccessTunnelConfiguration() BaseNetworkaccessTunnelConfigurationImpl {
	return s.networkaccessTunnelConfiguration
}

func UnmarshalNetworkaccessTunnelConfigurationImplementation(input []byte) (NetworkaccessTunnelConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessTunnelConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.tunnelConfigurationIKEv2Custom") {
		var out NetworkaccessTunnelConfigurationIKEv2Custom
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessTunnelConfigurationIKEv2Custom: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.tunnelConfigurationIKEv2Default") {
		var out NetworkaccessTunnelConfigurationIKEv2Default
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessTunnelConfigurationIKEv2Default: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessTunnelConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessTunnelConfigurationImpl: %+v", err)
	}

	return RawNetworkaccessTunnelConfigurationImpl{
		networkaccessTunnelConfiguration: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
