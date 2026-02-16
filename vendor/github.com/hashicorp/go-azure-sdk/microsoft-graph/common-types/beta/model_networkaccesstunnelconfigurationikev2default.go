package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessTunnelConfiguration = NetworkaccessTunnelConfigurationIKEv2Default{}

type NetworkaccessTunnelConfigurationIKEv2Default struct {

	// Fields inherited from NetworkaccessTunnelConfiguration

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

func (s NetworkaccessTunnelConfigurationIKEv2Default) NetworkaccessTunnelConfiguration() BaseNetworkaccessTunnelConfigurationImpl {
	return BaseNetworkaccessTunnelConfigurationImpl{
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
		PreSharedKey:               s.PreSharedKey,
		ZoneRedundancyPreSharedKey: s.ZoneRedundancyPreSharedKey,
	}
}

var _ json.Marshaler = NetworkaccessTunnelConfigurationIKEv2Default{}

func (s NetworkaccessTunnelConfigurationIKEv2Default) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessTunnelConfigurationIKEv2Default
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessTunnelConfigurationIKEv2Default: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessTunnelConfigurationIKEv2Default: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.tunnelConfigurationIKEv2Default"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessTunnelConfigurationIKEv2Default: %+v", err)
	}

	return encoded, nil
}
