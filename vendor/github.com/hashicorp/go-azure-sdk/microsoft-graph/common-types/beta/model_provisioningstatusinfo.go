package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStatusInfo struct {
	// If status isn't success/ skipped details for the error are contained in this.
	ErrorInformation *ProvisioningErrorInfo `json:"errorInformation,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Possible values are: success, failure, skipped, warning, unknownFutureValue. Supports $filter (eq, contains).
	Status *ProvisioningResult `json:"status,omitempty"`
}
