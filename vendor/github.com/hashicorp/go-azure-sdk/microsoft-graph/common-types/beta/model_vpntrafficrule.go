package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTrafficRule struct {
	// App identifier, if this traffic rule is triggered by an app.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Indicates the type of app that a VPN traffic rule is associated with.
	AppType *VpnTrafficRuleAppType `json:"appType,omitempty"`

	// Claims associated with this traffic rule.
	Claims nullable.Type[string] `json:"claims,omitempty"`

	// Local address range. This collection can contain a maximum of 500 elements.
	LocalAddressRanges *[]IPv4Range `json:"localAddressRanges,omitempty"`

	// Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum
	// of 500 elements.
	LocalPortRanges *[]NumberRange `json:"localPortRanges,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Protocols (0-255). Valid values 0 to 255
	Protocols nullable.Type[int64] `json:"protocols,omitempty"`

	// Remote address range. This collection can contain a maximum of 500 elements.
	RemoteAddressRanges *[]IPv4Range `json:"remoteAddressRanges,omitempty"`

	// Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum
	// of 500 elements.
	RemotePortRanges *[]NumberRange `json:"remotePortRanges,omitempty"`

	// Specifies the routing policy for a VPN traffic rule.
	RoutingPolicyType *VpnTrafficRuleRoutingPolicyType `json:"routingPolicyType,omitempty"`

	// Specify whether the rule applies to inbound traffic or outbound traffic.
	VpnTrafficDirection *VpnTrafficDirection `json:"vpnTrafficDirection,omitempty"`
}
