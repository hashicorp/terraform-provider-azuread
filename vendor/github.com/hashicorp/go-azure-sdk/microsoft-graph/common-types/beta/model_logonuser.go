package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUser struct {
	// Domain of user account used to logon.
	AccountDomain nullable.Type[string] `json:"accountDomain,omitempty"`

	// Account name of user account used to logon.
	AccountName nullable.Type[string] `json:"accountName,omitempty"`

	// User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
	AccountType *UserAccountSecurityType `json:"accountType,omitempty"`

	// DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// User logon ID.
	LogonId nullable.Type[string] `json:"logonId,omitempty"`

	// Collection of the logon types observed for the logged on user from when first to last seen. Possible values are:
	// unknown, interactive, remoteInteractive, network, batch, service.
	LogonTypes *[]LogonType `json:"logonTypes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
