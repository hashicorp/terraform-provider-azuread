package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPortComponent struct {
	Component *SecurityHostComponent `json:"component,omitempty"`

	// The first date and time when Microsoft Defender Threat Intelligence observed the hostPortComponent. The timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on
	// Jan 1, 2014, is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// Indicates whether this hostPortComponent is recent, which is determined by whether the hostPortComponent was observed
	// either at the same time or after the latest hostPortBanner in the scan history, or within two days of the latest scan
	// of the hostPort when there are no hostPortBanners in the scan history.
	IsRecent nullable.Type[bool] `json:"isRecent,omitempty"`

	// The last date and time when Microsoft Defender Threat Intelligence observed the hostPortComponent. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
