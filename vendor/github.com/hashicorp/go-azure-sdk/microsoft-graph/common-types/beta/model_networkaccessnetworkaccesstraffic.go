package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkAccessTraffic struct {
	// Indicates the action taken based on filtering policies. The possible values are: block, allow, unknownFutureValue,
	// bypass, alert. Use the Prefer: include-unknown-enum-members request header to get the following values from this
	// {evolvable enum}(/graph/best-practices-concept#handling-future-members-in-evolvable-enumerations): bypass , alert.
	Action *NetworkaccessFilteringPolicyAction `json:"action,omitempty"`

	// Represents the version of the Global Secure Access (GSA) client agent software. Supports $filter (eq) and $orderby.
	AgentVersion nullable.Type[string] `json:"agentVersion,omitempty"`

	// Destination Application ID accessed in Azure AD during the transaction. Supports $filter (eq) and $orderby.
	ApplicationSnapshot *NetworkaccessApplicationSnapshot `json:"applicationSnapshot,omitempty"`

	// Represents a unique identifier assigned to a connection. Supports $filter (eq) and $orderby.
	ConnectionId nullable.Type[string] `json:"connectionId,omitempty"`

	// Represents the date and time when a network access traffic log entry was created. Supports $filter (eq) and $orderby.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Informational error message. For example: 'Threat intelligence detected a transaction and triggered an alert.' or
	// 'The Global Secure Access (GSA) policy blocked the destination and triggered an alert.' Supports $filter (eq) and
	// $orderby.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Represents the Fully Qualified Domain Name (FQDN) of the destination host or server in a network communication.
	// Supports $filter (eq) and $orderby.
	DestinationFQDN nullable.Type[string] `json:"destinationFQDN,omitempty"`

	// Represents the IP address of the destination host or server in a network communication. Supports $filter (eq) and
	// $orderby.
	DestinationIp nullable.Type[string] `json:"destinationIp,omitempty"`

	// Represents the network port number on the destination host or server in a network communication. Supports $filter
	// (eq) and $orderby.
	DestinationPort nullable.Type[int64] `json:"destinationPort,omitempty"`

	// Represents the URL of the destination in a network communication. Supports $filter (eq) and $orderby.
	DestinationUrl nullable.Type[string] `json:"destinationUrl,omitempty"`

	// The destination FQDN's Web Category (e.g., Gambling). Supports $filter (eq) and $orderby.
	DestinationWebCategory *NetworkaccessWebCategory `json:"destinationWebCategory,omitempty"`

	// Represents the device associated with the network traffic, providing details about the hardware or virtual machine
	// involved in the transaction.
	Device *Device `json:"device,omitempty"`

	// Represents the category classification of a device within a network infrastructure. The possible values are: client,
	// branch, unknownFutureValue. Supports $filter (eq) and $orderby.
	DeviceCategory *NetworkaccessDeviceCategory `json:"deviceCategory,omitempty"`

	// Represents a unique identifier assigned to a device within a network infrastructure. Supports $filter (eq) and
	// $orderby.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Represents the operating system installed on a device within a network infrastructure. Supports $filter (eq) and
	// $orderby.
	DeviceOperatingSystem nullable.Type[string] `json:"deviceOperatingSystem,omitempty"`

	// Represents the version or release number of the operating system installed on a device within a network
	// infrastructure. Supports $filter (eq) and $orderby.
	DeviceOperatingSystemVersion nullable.Type[string] `json:"deviceOperatingSystemVersion,omitempty"`

	// The ID of the Filtering Profile associated with the action performed on traffic. Supports $filter (eq) and $orderby.
	FilteringProfileId nullable.Type[string] `json:"filteringProfileId,omitempty"`

	// The name of the Filtering Profile associated with the action performed on traffic. Supports $filter (eq) and
	// $orderby.
	FilteringProfileName nullable.Type[string] `json:"filteringProfileName,omitempty"`

	// Represents the headers included in a network request or response. Supports $filter (eq) and $orderby.
	Headers *NetworkaccessHeaders `json:"headers,omitempty"`

	// The HTTP method inspected in the intercepted HTTP traffic. Supports $filter (eq) and $orderby.
	HttpMethod *NetworkaccessHttpMethod `json:"httpMethod,omitempty"`

	// The process initiating the traffic transaction. Supports $filter (eq) and $orderby.
	InitiatingProcessName nullable.Type[string] `json:"initiatingProcessName,omitempty"`

	// Represents the networking protocol used for communication. The possible values are: ip, icmp, igmp, ggp, ipv4, tcp,
	// pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload,
	// ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII,
	// unknownFutureValue. Supports $filter (eq) and $orderby.
	NetworkProtocol *NetworkaccessNetworkingProtocol `json:"networkProtocol,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indication if traffic was successfully processed. The possible values are: success, failure, unknownFutureValue.
	// Supports $filter (eq) and $orderby.
	OperationStatus *NetworkaccessNetworkTrafficOperationStatus `json:"operationStatus,omitempty"`

	// Represents a unique identifier assigned to a policy. Supports $filter (eq) and $orderby.
	PolicyId nullable.Type[string] `json:"policyId,omitempty"`

	// The name of the filtering policy associated with the action performed on traffic. Supports $filter (eq) and $orderby.
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`

	// Represents a unique identifier assigned to a policy rule. Supports $filter (eq) and $orderby.
	PolicyRuleId nullable.Type[string] `json:"policyRuleId,omitempty"`

	// The name of the rule associated with the action performed on traffic. Supports $filter (eq) and $orderby.
	PolicyRuleName nullable.Type[string] `json:"policyRuleName,omitempty"`

	// The Point-of-Presence processing region of the traffic. Supports $filter (eq) and $orderby.
	PopProcessingRegion nullable.Type[string] `json:"popProcessingRegion,omitempty"`

	// Details about private access traffic. Supports $filter (eq) and $orderby.
	PrivateAccessDetails *NetworkaccessPrivateAccessDetails `json:"privateAccessDetails,omitempty"`

	// Represents the total number of bytes received in a network communication or data transfer. Supports $filter (eq) and
	// $orderby.
	ReceivedBytes nullable.Type[int64] `json:"receivedBytes,omitempty"`

	// The ID from which traffic was sent or received, providing visibility into the origin of the traffic. Supports $filter
	// (eq) and $orderby.
	RemoteNetworkId nullable.Type[string] `json:"remoteNetworkId,omitempty"`

	// Tenant ID that owns the resource. Supports $filter (eq) and $orderby.
	ResourceTenantId nullable.Type[string] `json:"resourceTenantId,omitempty"`

	// The HTTP response code inspected in the intercepted HTTP traffic. Supports $filter (eq) and $orderby.
	ResponseCode nullable.Type[int64] `json:"responseCode,omitempty"`

	// Represents the total number of bytes sent in a network communication or data transfer. Supports $filter (eq) and
	// $orderby.
	SentBytes nullable.Type[int64] `json:"sentBytes,omitempty"`

	// Represents a unique identifier assigned to a session or connection within a network infrastructure. Supports $filter
	// (eq) and $orderby.
	SessionId nullable.Type[string] `json:"sessionId,omitempty"`

	// Represents the source IP address in a network communication. Supports $filter (eq) and $orderby.
	SourceIp nullable.Type[string] `json:"sourceIp,omitempty"`

	// Represents the network port number on the source host or device in a network communication. Supports $filter (eq) and
	// $orderby.
	SourcePort nullable.Type[int64] `json:"sourcePort,omitempty"`

	// Represents a unique identifier assigned to a tenant within a network infrastructure. Supports $filter (eq) and
	// $orderby.
	TenantId *string `json:"tenantId,omitempty"`

	// The type of threat detected in the traffic. Supports $filter (eq) and $orderby.
	ThreatType nullable.Type[string] `json:"threatType,omitempty"`

	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`

	// Represents a unique identifier assigned to a specific transaction or operation. Key. Supports $filter (eq) and
	// $orderby.
	TransactionId *string `json:"transactionId,omitempty"`

	// Represents the transport protocol used for communication. The possible values are: ip, icmp, igmp, ggp, ipv4, tcp,
	// pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload,
	// ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII,
	// unknownFutureValue. Supports $filter (eq) and $orderby.
	TransportProtocol *NetworkaccessNetworkingProtocol `json:"transportProtocol,omitempty"`

	// Represents the user associated with the network traffic, providing details about the individual or account initiating
	// the transaction.
	User *User `json:"user,omitempty"`

	// Represents a unique identifier assigned to a user. Supports $filter (eq) and $orderby.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Represents the user principal name (UPN) associated with a user. Supports $filter (eq) and $orderby.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The name of the vendors who detected the threat. Supports $filter (eq) and $orderby.
	VendorNames *[]string `json:"vendorNames,omitempty"`
}
