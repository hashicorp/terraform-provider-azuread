package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRule struct {
	// State Management Setting.
	Action *StateManagementSetting `json:"action,omitempty"`

	// The description of the rule.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the rule. Does not need to be unique.
	DisplayName *string `json:"displayName,omitempty"`

	// State Management Setting.
	EdgeTraversal *StateManagementSetting `json:"edgeTraversal,omitempty"`

	// The full file path of an app that's affected by the firewall rule.
	FilePath nullable.Type[string] `json:"filePath,omitempty"`

	// Flags representing firewall rule interface types.
	InterfaceTypes *WindowsFirewallRuleInterfaceTypes `json:"interfaceTypes,omitempty"`

	// List of local addresses covered by the rule. Default is any address. Valid tokens include:'' indicates any local
	// address. If present, this must be the only token included.A subnet can be specified using either the subnet mask or
	// network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to
	// 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no
	// spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
	LocalAddressRanges *[]string `json:"localAddressRanges,omitempty"`

	// List of local port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
	LocalPortRanges *[]string `json:"localPortRanges,omitempty"`

	// Specifies the list of authorized local users for the app container. This is a string in Security Descriptor
	// Definition Language (SDDL) format.
	LocalUserAuthorizations nullable.Type[string] `json:"localUserAuthorizations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The package family name of a Microsoft Store application that's affected by the firewall rule.
	PackageFamilyName nullable.Type[string] `json:"packageFamilyName,omitempty"`

	// Flags representing which network profile types apply to a firewall rule.
	ProfileTypes *WindowsFirewallRuleNetworkProfileTypes `json:"profileTypes,omitempty"`

	// 0-255 number representing the IP protocol (TCP = 6, UDP = 17). If not specified, the default is All. Valid values 0
	// to 255
	Protocol nullable.Type[int64] `json:"protocol,omitempty"`

	// List of tokens specifying the remote addresses covered by the rule. Tokens are case insensitive. Default is any
	// address. Valid tokens include:'' indicates any remote address. If present, this must be the only token
	// included.'Defaultgateway''DHCP''DNS''WINS''Intranet' (supported on Windows versions 1809+)'RmtIntranet' (supported on
	// Windows versions 1809+)'Internet' (supported on Windows versions 1809+)'Ply2Renders' (supported on Windows versions
	// 1809+)'LocalSubnet' indicates any local address on the local subnet.A subnet can be specified using either the subnet
	// mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults
	// to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no
	// spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
	RemoteAddressRanges *[]string `json:"remoteAddressRanges,omitempty"`

	// List of remote port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
	RemotePortRanges *[]string `json:"remotePortRanges,omitempty"`

	// The name used in cases when a service, not an application, is sending or receiving traffic.
	ServiceName nullable.Type[string] `json:"serviceName,omitempty"`

	// Firewall rule traffic directions.
	TrafficDirection *WindowsFirewallRuleTrafficDirectionType `json:"trafficDirection,omitempty"`
}
