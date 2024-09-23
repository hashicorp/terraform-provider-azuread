package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfo struct {
	// The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
	BasicServiceSetIdentifier nullable.Type[string] `json:"basicServiceSetIdentifier,omitempty"`

	ConnectionType *CallRecordsNetworkConnectionType `json:"connectionType,omitempty"`

	// DNS suffix associated with the network adapter of the media endpoint.
	DnsSuffix nullable.Type[string] `json:"dnsSuffix,omitempty"`

	// IP address of the media endpoint.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Link speed in bits per second reported by the network adapter used by the media endpoint.
	LinkSpeed nullable.Type[int64] `json:"linkSpeed,omitempty"`

	// The media access control (MAC) address of the media endpoint's network device. This value may be missing or shown as
	// 02:00:00:00:00:00 due to operating system privacy policies.
	MacAddress nullable.Type[string] `json:"macAddress,omitempty"`

	NetworkTransportProtocol *CallRecordsNetworkTransportProtocol `json:"networkTransportProtocol,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Network port number used by media endpoint.
	Port nullable.Type[int64] `json:"port,omitempty"`

	// IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address
	// associated to the endpoint.
	ReflexiveIPAddress nullable.Type[string] `json:"reflexiveIPAddress,omitempty"`

	// IP address of the media relay server allocated by the media endpoint.
	RelayIPAddress nullable.Type[string] `json:"relayIPAddress,omitempty"`

	// Network port number allocated on the media relay server by the media endpoint.
	RelayPort nullable.Type[int64] `json:"relayPort,omitempty"`

	// Subnet used for media stream by the media endpoint.
	Subnet nullable.Type[string] `json:"subnet,omitempty"`

	// List of network trace route hops collected for this media stream.*
	TraceRouteHops *[]CallRecordsTraceRouteHop `json:"traceRouteHops,omitempty"`

	WifiBand *CallRecordsWifiBand `json:"wifiBand,omitempty"`

	// Estimated remaining battery charge in percentage reported by the media endpoint.
	WifiBatteryCharge nullable.Type[int64] `json:"wifiBatteryCharge,omitempty"`

	// WiFi channel used by the media endpoint.
	WifiChannel nullable.Type[int64] `json:"wifiChannel,omitempty"`

	// Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by
	// endpoint.
	WifiMicrosoftDriver nullable.Type[string] `json:"wifiMicrosoftDriver,omitempty"`

	// Version of the Microsoft WiFi driver used by the media endpoint.
	WifiMicrosoftDriverVersion nullable.Type[string] `json:"wifiMicrosoftDriverVersion,omitempty"`

	WifiRadioType *CallRecordsWifiRadioType `json:"wifiRadioType,omitempty"`

	// WiFi signal strength in percentage reported by the media endpoint.
	WifiSignalStrength nullable.Type[int64] `json:"wifiSignalStrength,omitempty"`

	// Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
	WifiVendorDriver nullable.Type[string] `json:"wifiVendorDriver,omitempty"`

	// Version of the WiFi driver used by the media endpoint.
	WifiVendorDriverVersion nullable.Type[string] `json:"wifiVendorDriverVersion,omitempty"`
}
