package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResizeValidationResult struct {
	// The cloudPC ID that corresponds to its unique identifier.
	CloudPCId nullable.Type[string] `json:"cloudPcId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Describes a list of the validation result for the Cloud PC resize action. The possible values are: success,
	// cloudPcNotFound, operationCnflict, operationNotSupported, targetLicenseHasAssigned, internalServerError, and
	// unknownFutureValue.
	ValidationResult *CloudPCResizeValidationCode `json:"validationResult,omitempty"`
}
