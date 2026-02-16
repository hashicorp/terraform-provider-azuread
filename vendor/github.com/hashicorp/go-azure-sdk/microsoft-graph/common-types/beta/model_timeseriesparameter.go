package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeSeriesParameter struct {
	// End time of the series being requested. Optional; if not specified, current time is used.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The name of the metric for which a time series is requested.
	MetricName nullable.Type[string] `json:"metricName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Start time of the series being requested.
	StartDateTime *string `json:"startDateTime,omitempty"`
}
