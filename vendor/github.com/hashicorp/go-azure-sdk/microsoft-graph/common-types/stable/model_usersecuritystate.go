package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityState struct {
	// AAD User object identifier (GUID) - represents the physical/multi-account user entity.
	AadUserId nullable.Type[string] `json:"aadUserId,omitempty"`

	// Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
	AccountName nullable.Type[string] `json:"accountName,omitempty"`

	// NetBIOS/Active Directory domain of user account (that is, domain/account format).
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
	EmailRole *EmailRole `json:"emailRole,omitempty"`

	// Indicates whether the user logged on through a VPN.
	IsVpn nullable.Type[bool] `json:"isVpn,omitempty"`

	// Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LogonDateTime nullable.Type[string] `json:"logonDateTime,omitempty"`

	// User sign-in ID.
	LogonId nullable.Type[string] `json:"logonId,omitempty"`

	// IP Address the sign-in request originated from.
	LogonIp nullable.Type[string] `json:"logonIp,omitempty"`

	// Location (by IP address mapping) associated with a user sign-in event by this user.
	LogonLocation nullable.Type[string] `json:"logonLocation,omitempty"`

	// Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
	LogonType *LogonType `json:"logonType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Active Directory (on-premises) Security Identifier (SID) of the user.
	OnPremisesSecurityIdentifier nullable.Type[string] `json:"onPremisesSecurityIdentifier,omitempty"`

	// Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a
	// percentage.
	RiskScore nullable.Type[string] `json:"riskScore,omitempty"`

	// User account type (group membership), per Windows definition. Possible values are: unknown, standard, power,
	// administrator.
	UserAccountType *UserAccountSecurityType `json:"userAccountType,omitempty"`

	// User sign-in name - internet format: (user account name)@(user account DNS domain name).
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
