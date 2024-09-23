package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BulkDriverActionResult struct {
	// List of driver Ids where the action is failed.
	FailedDriverIds *[]string `json:"failedDriverIds,omitempty"`

	// List of driver Ids that are not found.
	NotFoundDriverIds *[]string `json:"notFoundDriverIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// List of driver Ids where the action is successful.
	SuccessfulDriverIds *[]string `json:"successfulDriverIds,omitempty"`
}
