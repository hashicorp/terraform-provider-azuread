package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsTraceRouteHop struct {
	// The network path count of this hop that was used to compute the round-trip time.
	HopCount nullable.Type[int64] `json:"hopCount,omitempty"`

	// IP address used for this hop in the network trace.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time from when the trace route packet was sent from the client to this hop and back to the client, denoted in ISO
	// 8601 format. For example, 1 second is denoted as PT1S, where P is the duration designator, T is the time designator,
	// and S is the second designator.
	RoundTripTime nullable.Type[string] `json:"roundTripTime,omitempty"`
}
