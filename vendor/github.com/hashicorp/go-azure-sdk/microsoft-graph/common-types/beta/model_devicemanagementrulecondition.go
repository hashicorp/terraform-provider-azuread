package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleCondition struct {
	// The built-in aggregation method for the rule condition. The possible values are: count, percentage,
	// affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
	Aggregation *DeviceManagementAggregationType `json:"aggregation,omitempty"`

	// The property that the rule condition monitors. Possible values are: provisionFailures, imageUploadFailures,
	// azureNetworkConnectionCheckFailures, cloudPcInGracePeriod, frontlineInsufficientLicenses, cloudPcConnectionErrors,
	// cloudPcHostHealthCheckFailures, cloudPcZoneOutage, unknownFutureValue.
	ConditionCategory *DeviceManagementConditionCategory `json:"conditionCategory,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The built-in operator for the rule condition. The possible values are: greaterOrEqual, equal, greater, less,
	// lessOrEqual, notEqual, unknownFutureValue.
	Operator *DeviceManagementOperatorType `json:"operator,omitempty"`

	// The relationship type. Possible values are: and, or.
	RelationshipType *DeviceManagementRelationshipType `json:"relationshipType,omitempty"`

	// The threshold value of the alert condition. The threshold value can be a number in string form or string like
	// 'WestUS'.
	ThresholdValue nullable.Type[string] `json:"thresholdValue,omitempty"`
}
