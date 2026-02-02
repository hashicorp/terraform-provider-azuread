package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningErrorInfo struct {
	// Additional details if there's error.
	AdditionalDetails nullable.Type[string] `json:"additionalDetails,omitempty"`

	// Categorizes the error code. Possible values are failure, nonServiceFailure, success, unknownFutureValue
	ErrorCategory *ProvisioningStatusErrorCategory `json:"errorCategory,omitempty"`

	// Unique error code if any occurred. Learn more
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Summarizes the status and describes why the status happened.
	Reason nullable.Type[string] `json:"reason,omitempty"`

	// Provides the resolution for the corresponding error.
	RecommendedAction nullable.Type[string] `json:"recommendedAction,omitempty"`
}
