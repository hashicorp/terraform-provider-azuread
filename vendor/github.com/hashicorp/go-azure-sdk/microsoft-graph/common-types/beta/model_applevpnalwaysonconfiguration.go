package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfiguration struct {
	// Determine whether AirPrint service will be exempt from the always-on VPN connection. Possible values are:
	// forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
	AirPrintExceptionAction *VpnServiceExceptionAction `json:"airPrintExceptionAction,omitempty"`

	// Specifies whether traffic from all captive network plugins should be allowed outside the vpn
	AllowAllCaptiveNetworkPlugins nullable.Type[bool] `json:"allowAllCaptiveNetworkPlugins,omitempty"`

	// Determines whether traffic from the Websheet app is allowed outside of the VPN
	AllowCaptiveWebSheet nullable.Type[bool] `json:"allowCaptiveWebSheet,omitempty"`

	// Determines whether all, some, or no non-native captive networking apps are allowed
	AllowedCaptiveNetworkPlugins *SpecifiedCaptiveNetworkPlugins `json:"allowedCaptiveNetworkPlugins,omitempty"`

	// Determine whether Cellular service will be exempt from the always-on VPN connection. Possible values are:
	// forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
	CellularExceptionAction *VpnServiceExceptionAction `json:"cellularExceptionAction,omitempty"`

	// Specifies how often in seconds to send a network address translation keepalive package through the VPN
	NatKeepAliveIntervalInSeconds nullable.Type[int64] `json:"natKeepAliveIntervalInSeconds,omitempty"`

	// Enable hardware offloading of NAT keepalive signals when the device is asleep
	NatKeepAliveOffloadEnable nullable.Type[bool] `json:"natKeepAliveOffloadEnable,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of tunnels that will be present to the VPN client for configuration
	TunnelConfiguration *VpnTunnelConfigurationType `json:"tunnelConfiguration,omitempty"`

	// Allow the user to toggle the VPN configuration using the UI
	UserToggleEnabled nullable.Type[bool] `json:"userToggleEnabled,omitempty"`

	// Determine whether voicemail service will be exempt from the always-on VPN connection. Possible values are:
	// forceTrafficViaVPN, allowTrafficOutside, dropTraffic.
	VoicemailExceptionAction *VpnServiceExceptionAction `json:"voicemailExceptionAction,omitempty"`
}
