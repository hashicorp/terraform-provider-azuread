package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetails struct {
	// Error code of the most recent run that encountered an error. The possible values are: queryExecutionFailed,
	// queryExecutionThrottling, queryExceededResultSize, queryLimitsExceeded, queryTimeout, alertCreationFailed,
	// alertReportNotFound, partialRowsFailed, unknownFutureValue.
	ErrorCode *SecurityHuntingRuleErrorCode `json:"errorCode,omitempty"`

	// Reason for failure when the custom detection last ran and failed. See the table below.
	FailureReason nullable.Type[string] `json:"failureReason,omitempty"`

	// Timestamp when the custom detection was last run.
	LastRunDateTime nullable.Type[string] `json:"lastRunDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Status of custom detection when it was last run. The possible values are: running, completed, failed,
	// partiallyFailed, unknownFutureValue.
	Status *SecurityHuntingRuleRunStatus `json:"status,omitempty"`
}
