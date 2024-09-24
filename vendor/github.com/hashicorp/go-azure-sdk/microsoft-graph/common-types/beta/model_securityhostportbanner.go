package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPortBanner struct {
	// The text response received from a web component when scanning a hostPort.
	Banner *string `json:"banner,omitempty"`

	// The first date and time when Microsoft Defender Threat Intelligence observed the hostPortBanner. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// The last date and time when Microsoft Defender Threat Intelligence observed the hostPortBanner. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The specific protocol used to scan the hostPort.
	ScanProtocol nullable.Type[string] `json:"scanProtocol,omitempty"`

	// The total amount of times that Microsoft Defender Threat Intelligence has observed the hostPortBanner in all its
	// scans.
	TimesObserved nullable.Type[int64] `json:"timesObserved,omitempty"`
}
