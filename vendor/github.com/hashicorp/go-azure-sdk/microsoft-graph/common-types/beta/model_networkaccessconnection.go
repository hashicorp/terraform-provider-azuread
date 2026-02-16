package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessConnection{}

type NetworkaccessConnection struct {
	// The version of the client that initiated the connection.
	AgentVersion nullable.Type[string] `json:"agentVersion,omitempty"`

	// appId (or client ID) of the destination Microsoft Entra application.
	ApplicationSnapshot *NetworkaccessApplicationSnapshot `json:"applicationSnapshot,omitempty"`

	// The time the connection was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The destination FQDN of the connection.
	DestinationFqdn nullable.Type[string] `json:"destinationFqdn,omitempty"`

	// The destination IP of the connection.
	DestinationIp nullable.Type[string] `json:"destinationIp,omitempty"`

	// The destination port of the connection.
	DestinationPort nullable.Type[int64] `json:"destinationPort,omitempty"`

	// The category of the device. The possible values are: client, branch, unknownFutureValue, remoteNetwork. Use the
	// Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum:
	// remoteNetwork.
	DeviceCategory *NetworkaccessDeviceCategory `json:"deviceCategory,omitempty"`

	// The DeviceID.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The device operating system type.
	DeviceOperatingSystem nullable.Type[string] `json:"deviceOperatingSystem,omitempty"`

	// The device operating system version.
	DeviceOperatingSystemVersion nullable.Type[string] `json:"deviceOperatingSystemVersion,omitempty"`

	// The time the connection was terminated.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The process initiating the traffic connection.
	InitiatingProcessName nullable.Type[string] `json:"initiatingProcessName,omitempty"`

	// When the connection was last updated.
	LastUpdateDateTime nullable.Type[string] `json:"lastUpdateDateTime,omitempty"`

	// The network protocol of the connection. The possible values are: ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6,
	// ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6,
	// ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII, unknownFutureValue.
	NetworkProtocol *NetworkaccessNetworkingProtocol `json:"networkProtocol,omitempty"`

	// The Point-of-Presence processing region of the traffic.
	PopProcessingRegion nullable.Type[string] `json:"popProcessingRegion,omitempty"`

	// Private access details.
	PrivateAccessDetails *NetworkaccessPrivateAccessDetails `json:"privateAccessDetails,omitempty"`

	// Accumulative bytes received.
	ReceivedBytes nullable.Type[int64] `json:"receivedBytes,omitempty"`

	// Accumulative bytes sent.
	SentBytes nullable.Type[int64] `json:"sentBytes,omitempty"`

	// The source IP of the connection.
	SourceIp nullable.Type[string] `json:"sourceIp,omitempty"`

	// The source port of the connection.
	SourcePort nullable.Type[int64] `json:"sourcePort,omitempty"`

	// Status of the connection. The possible values are: open, active, closed, unknownFutureValue.
	Status *NetworkaccessConnectionStatus `json:"status,omitempty"`

	// The ID of the tenant where the connection was initiated.
	TenantId *string `json:"tenantId,omitempty"`

	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`

	// The number of blocked transactions belonging to the connection.
	TransactionBlockCount nullable.Type[int64] `json:"transactionBlockCount,omitempty"`

	// The number of transactions belonging to the connection.
	TransactionCount nullable.Type[int64] `json:"transactionCount,omitempty"`

	// The transport protocol of the connection. The possible values are: ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp,
	// ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6,
	// ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII, unknownFutureValue.
	TransportProtocol *NetworkaccessNetworkingProtocol `json:"transportProtocol,omitempty"`

	// The user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The principal name of the user.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s NetworkaccessConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessConnection{}

func (s NetworkaccessConnection) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessConnection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.connection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessConnection: %+v", err)
	}

	return encoded, nil
}
