package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTaskExecution struct {
	// Identifier of the job run.
	ActivityIdentifier nullable.Type[string] `json:"activityIdentifier,omitempty"`

	// Count of processed entries that were assigned for this application.
	CountEntitled *int64 `json:"countEntitled,omitempty"`

	// Count of processed entries that were assigned for provisioning.
	CountEntitledForProvisioning *int64 `json:"countEntitledForProvisioning,omitempty"`

	// Count of entries that were escrowed (errors).
	CountEscrowed *int64 `json:"countEscrowed,omitempty"`

	// Count of entries that were escrowed, including system-generated escrows.
	CountEscrowedRaw *int64 `json:"countEscrowedRaw,omitempty"`

	// Count of exported entries.
	CountExported *int64 `json:"countExported,omitempty"`

	// Count of entries that were expected to be exported.
	CountExports *int64 `json:"countExports,omitempty"`

	// Count of imported entries.
	CountImported *int64 `json:"countImported,omitempty"`

	// Count of imported delta-changes.
	CountImportedDeltas *int64 `json:"countImportedDeltas,omitempty"`

	// Count of imported delta-changes pertaining to reference changes.
	CountImportedReferenceDeltas *int64 `json:"countImportedReferenceDeltas,omitempty"`

	// If an error was encountered, contains a synchronizationError object with details.
	Error *SynchronizationError `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *SynchronizationTaskExecutionResult `json:"state,omitempty"`

	// Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	TimeBegan *string `json:"timeBegan,omitempty"`

	// Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	TimeEnded *string `json:"timeEnded,omitempty"`
}
