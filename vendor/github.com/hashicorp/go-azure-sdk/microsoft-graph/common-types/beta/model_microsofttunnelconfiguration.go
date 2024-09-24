package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MicrosoftTunnelConfiguration{}

type MicrosoftTunnelConfiguration struct {
	// Additional settings that may be applied to the server
	AdvancedSettings *[]KeyValuePair `json:"advancedSettings,omitempty"`

	// The Default Domain appendix that will be used by the clients
	DefaultDomainSuffix nullable.Type[string] `json:"defaultDomainSuffix,omitempty"`

	// The configuration's description (optional)
	Description nullable.Type[string] `json:"description,omitempty"`

	// When DisableUdpConnections is set, the clients and VPN server will not use DTLS connections to transfer data.
	DisableUdpConnections nullable.Type[bool] `json:"disableUdpConnections,omitempty"`

	// The display name for the server configuration. This property is required when a server is created.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The DNS servers that will be used by the clients
	DnsServers *[]string `json:"dnsServers,omitempty"`

	// The IPv6 subnet that will be used to allocate virtual address for the clients
	IPv6Network nullable.Type[string] `json:"ipv6Network,omitempty"`

	// When the configuration was last updated
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`

	// The port that both TCP and UPD will listen over on the server
	ListenPort nullable.Type[int64] `json:"listenPort,omitempty"`

	// The IPv4 subnet that will be used to allocate virtual address for the clients
	Network nullable.Type[string] `json:"network,omitempty"`

	// List of Scope Tags for this Entity instance
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Subsets of the routes that will not be routed by the server
	RouteExcludes *[]string `json:"routeExcludes,omitempty"`

	// The routes that will be routed by the server
	RouteIncludes *[]string `json:"routeIncludes,omitempty"`

	// Subsets of the routes that will not be routed by the server. This property is going to be deprecated with the option
	// of using the new property, 'RouteExcludes'.
	RoutesExclude *[]string `json:"routesExclude,omitempty"`

	// The routes that will be routed by the server. This property is going to be deprecated with the option of using the
	// new property, 'RouteIncludes'.
	RoutesInclude *[]string `json:"routesInclude,omitempty"`

	// The domains that will be resolved using the provided dns servers
	SplitDNS *[]string `json:"splitDNS,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MicrosoftTunnelConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftTunnelConfiguration{}

func (s MicrosoftTunnelConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftTunnelConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftTunnelConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftTunnelConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftTunnelConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftTunnelConfiguration: %+v", err)
	}

	return encoded, nil
}
