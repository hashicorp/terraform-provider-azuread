package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BulkCatalogItemActionResult struct {
	// List of catalog item Ids where the action is failed.
	FailedIds *[]string `json:"failedIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// List of catalog item Ids where the action is successful.
	SuccessfulIds *[]string `json:"successfulIds,omitempty"`
}
