package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallNetworkProfile struct {
	// Configures the firewall to merge authorized application rules from group policy with those from local store instead
	// of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and
	// AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes
	// priority.
	AuthorizedApplicationRulesFromGroupPolicyMerged *bool `json:"authorizedApplicationRulesFromGroupPolicyMerged,omitempty"`

	// Configures the firewall to prevent merging authorized application rules from group policy with those from local store
	// instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and
	// AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes
	// priority.
	AuthorizedApplicationRulesFromGroupPolicyNotMerged *bool `json:"authorizedApplicationRulesFromGroupPolicyNotMerged,omitempty"`

	// Configures the firewall to merge connection security rules from group policy with those from local store instead of
	// ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and
	// ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes
	// priority.
	ConnectionSecurityRulesFromGroupPolicyMerged *bool `json:"connectionSecurityRulesFromGroupPolicyMerged,omitempty"`

	// Configures the firewall to prevent merging connection security rules from group policy with those from local store
	// instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and
	// ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes
	// priority.
	ConnectionSecurityRulesFromGroupPolicyNotMerged *bool `json:"connectionSecurityRulesFromGroupPolicyNotMerged,omitempty"`

	// State Management Setting.
	FirewallEnabled *StateManagementSetting `json:"firewallEnabled,omitempty"`

	// Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring
	// the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both
	// true, GlobalPortRulesFromGroupPolicyMerged takes priority.
	GlobalPortRulesFromGroupPolicyMerged *bool `json:"globalPortRulesFromGroupPolicyMerged,omitempty"`

	// Configures the firewall to prevent merging global port rules from group policy with those from local store instead of
	// ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged
	// are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
	GlobalPortRulesFromGroupPolicyNotMerged *bool `json:"globalPortRulesFromGroupPolicyNotMerged,omitempty"`

	// Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and
	// InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
	InboundConnectionsBlocked *bool `json:"inboundConnectionsBlocked,omitempty"`

	// Configures the firewall to allow all incoming connections by default. When InboundConnectionsRequired and
	// InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
	InboundConnectionsRequired *bool `json:"inboundConnectionsRequired,omitempty"`

	// Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When
	// InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes
	// priority.
	InboundNotificationsBlocked *bool `json:"inboundNotificationsBlocked,omitempty"`

	// Allows the firewall to display notifications when an application is blocked from listening on a port. When
	// InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes
	// priority.
	InboundNotificationsRequired *bool `json:"inboundNotificationsRequired,omitempty"`

	// Configures the firewall to block all incoming traffic regardless of other policy settings. When
	// IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
	IncomingTrafficBlocked *bool `json:"incomingTrafficBlocked,omitempty"`

	// Configures the firewall to allow incoming traffic pursuant to other policy settings. When IncomingTrafficRequired and
	// IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
	IncomingTrafficRequired *bool `json:"incomingTrafficRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and
	// OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to
	// Windows releases version 1809 and above.
	OutboundConnectionsBlocked *bool `json:"outboundConnectionsBlocked,omitempty"`

	// Configures the firewall to allow all outgoing connections by default. When OutboundConnectionsRequired and
	// OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to
	// Windows releases version 1809 and above.
	OutboundConnectionsRequired *bool `json:"outboundConnectionsRequired,omitempty"`

	// Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of
	// ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are
	// both true, PolicyRulesFromGroupPolicyMerged takes priority.
	PolicyRulesFromGroupPolicyMerged *bool `json:"policyRulesFromGroupPolicyMerged,omitempty"`

	// Configures the firewall to prevent merging Firewall Rule policies from group policy with those from local store
	// instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and
	// PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
	PolicyRulesFromGroupPolicyNotMerged *bool `json:"policyRulesFromGroupPolicyNotMerged,omitempty"`

	// Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is
	// secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and
	// SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
	SecuredPacketExemptionAllowed *bool `json:"securedPacketExemptionAllowed,omitempty"`

	// Configures the firewall to block the host computer to respond to unsolicited network traffic of that traffic is
	// secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and
	// SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
	SecuredPacketExemptionBlocked *bool `json:"securedPacketExemptionBlocked,omitempty"`

	// Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true,
	// StealthModeBlocked takes priority.
	StealthModeBlocked *bool `json:"stealthModeBlocked,omitempty"`

	// Allow the server to operate in stealth mode. When StealthModeRequired and StealthModeBlocked are both true,
	// StealthModeBlocked takes priority.
	StealthModeRequired *bool `json:"stealthModeRequired,omitempty"`

	// Configures the firewall to block unicast responses to multicast broadcast traffic. When
	// UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true,
	// UnicastResponsesToMulticastBroadcastsBlocked takes priority.
	UnicastResponsesToMulticastBroadcastsBlocked *bool `json:"unicastResponsesToMulticastBroadcastsBlocked,omitempty"`

	// Configures the firewall to allow unicast responses to multicast broadcast traffic. When
	// UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true,
	// UnicastResponsesToMulticastBroadcastsBlocked takes priority.
	UnicastResponsesToMulticastBroadcastsRequired *bool `json:"unicastResponsesToMulticastBroadcastsRequired,omitempty"`
}
