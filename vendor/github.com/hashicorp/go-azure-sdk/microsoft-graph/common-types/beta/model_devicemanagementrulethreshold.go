package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleThreshold struct {
	// Indicates the built-in aggregation methods. The possible values are: count, percentage, affectedCloudPcCount,
	// affectedCloudPcPercentage, unknownFutureValue.
	Aggregation *DeviceManagementAggregationType `json:"aggregation,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the built-in operator. The possible values are: greaterOrEqual, equal, greater, less, lessOrEqual,
	// notEqual, unknownFutureValue.
	Operator *DeviceManagementOperatorType `json:"operator,omitempty"`

	// The target threshold value.
	Target nullable.Type[int64] `json:"target,omitempty"`
}
