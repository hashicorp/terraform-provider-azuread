package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationStatus struct {
	Code *SynchronizationStatusCode `json:"code,omitempty"`

	// Number of consecutive times this job failed.
	CountSuccessiveCompleteFailures *int64 `json:"countSuccessiveCompleteFailures,omitempty"`

	// true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if
	// during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine.
	// Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial
	// synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the
	// customer to clean up the errors.
	EscrowsPruned *bool `json:"escrowsPruned,omitempty"`

	// Details of the last execution of the job.
	LastExecution *SynchronizationTaskExecution `json:"lastExecution,omitempty"`

	// Details of the last execution of this job, which didn't have any errors.
	LastSuccessfulExecution *SynchronizationTaskExecution `json:"lastSuccessfulExecution,omitempty"`

	// Details of the last execution of the job, which exported objects into the target directory.
	LastSuccessfulExecutionWithExports *SynchronizationTaskExecution `json:"lastSuccessfulExecutionWithExports,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Details of the progress of a job toward completion.
	Progress *[]SynchronizationProgress `json:"progress,omitempty"`

	// If job is in quarantine, quarantine details.
	Quarantine *SynchronizationQuarantine `json:"quarantine,omitempty"`

	// The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date
	// and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	SteadyStateFirstAchievedTime *string `json:"steadyStateFirstAchievedTime,omitempty"`

	// The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	SteadyStateLastAchievedTime *string `json:"steadyStateLastAchievedTime,omitempty"`

	// Count of synchronized objects, listed by object type.
	SynchronizedEntryCountByType *[]StringKeyLongValuePair `json:"synchronizedEntryCountByType,omitempty"`

	// In the event of an error, the URL with the troubleshooting steps for the issue.
	TroubleshootingUrl nullable.Type[string] `json:"troubleshootingUrl,omitempty"`
}
