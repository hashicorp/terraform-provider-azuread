package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VpnProxyServer = Windows81VpnProxyServer{}

type Windows81VpnProxyServer struct {
	// Automatically detect proxy settings.
	AutomaticallyDetectProxySettings *bool `json:"automaticallyDetectProxySettings,omitempty"`

	// Bypass proxy server for local address.
	BypassProxyServerForLocalAddress *bool `json:"bypassProxyServerForLocalAddress,omitempty"`

	// Fields inherited from VpnProxyServer

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

func (s Windows81VpnProxyServer) VpnProxyServer() BaseVpnProxyServerImpl {
	return BaseVpnProxyServerImpl{
		Address:                         s.Address,
		AutomaticConfigurationScriptUrl: s.AutomaticConfigurationScriptUrl,
		ODataId:                         s.ODataId,
		ODataType:                       s.ODataType,
		Port:                            s.Port,
	}
}

var _ json.Marshaler = Windows81VpnProxyServer{}

func (s Windows81VpnProxyServer) MarshalJSON() ([]byte, error) {
	type wrapper Windows81VpnProxyServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows81VpnProxyServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows81VpnProxyServer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows81VpnProxyServer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows81VpnProxyServer: %+v", err)
	}

	return encoded, nil
}
