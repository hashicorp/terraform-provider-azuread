package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainState struct {
	// Timestamp for when the last activity occurred. The value is updated when an operation is scheduled, the asynchronous
	// task starts, and when the operation completes.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of asynchronous operation. The values can be ForceDelete or Verification.
	Operation nullable.Type[string] `json:"operation,omitempty"`

	// Current status of the operation. Scheduled - Operation is scheduled but hasn't started. InProgress - Task is in
	// progress. Failed - The operation failed.
	Status nullable.Type[string] `json:"status,omitempty"`
}
