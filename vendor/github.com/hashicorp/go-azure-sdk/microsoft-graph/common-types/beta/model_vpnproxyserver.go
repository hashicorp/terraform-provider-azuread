package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnProxyServer interface {
	VpnProxyServer() BaseVpnProxyServerImpl
}

var _ VpnProxyServer = BaseVpnProxyServerImpl{}

type BaseVpnProxyServerImpl struct {
	// Address.
	Address nullable.Type[string] `json:"address,omitempty"`

	// Proxy's automatic configuration script url.
	AutomaticConfigurationScriptUrl nullable.Type[string] `json:"automaticConfigurationScriptUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Port. Valid values 0 to 65535
	Port nullable.Type[int64] `json:"port,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseVpnProxyServerImpl) VpnProxyServer() BaseVpnProxyServerImpl {
	return s
}

var _ VpnProxyServer = RawVpnProxyServerImpl{}

// RawVpnProxyServerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawVpnProxyServerImpl struct {
	vpnProxyServer BaseVpnProxyServerImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawVpnProxyServerImpl) VpnProxyServer() BaseVpnProxyServerImpl {
	return s.vpnProxyServer
}

func UnmarshalVpnProxyServerImplementation(input []byte) (VpnProxyServer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling VpnProxyServer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10VpnProxyServer") {
		var out Windows10VpnProxyServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10VpnProxyServer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81VpnProxyServer") {
		var out Windows81VpnProxyServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81VpnProxyServer: %+v", err)
		}
		return out, nil
	}

	var parent BaseVpnProxyServerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseVpnProxyServerImpl: %+v", err)
	}

	return RawVpnProxyServerImpl{
		vpnProxyServer: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
