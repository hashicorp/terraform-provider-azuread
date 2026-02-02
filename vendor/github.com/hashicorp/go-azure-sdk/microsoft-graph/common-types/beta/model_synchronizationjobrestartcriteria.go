package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobRestartCriteria struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Comma-separated combination of the following values: None, ConnectorDataStore, Escrows, Watermark, QuarantineState,
	// Full, ForceDeletes. The property can also be empty. None: Starts a paused or quarantined provisioning job. DO NOT
	// USE. Use the Start synchronizationJob API instead.ConnectorDataStore - Clears the underlying cache for all users. DO
	// NOT USE. Contact Microsoft Support for guidance.Escrows - Provisioning failures are marked as escrows and retried.
	// Clearing escrows will stop the service from retrying failures.Watermark - Removing the watermark causes the service
	// to re-evaluate all the users again, rather than just processing changes.QuarantineState - Temporarily lifts the
	// quarantine.Use Full if you want all of the options.ForceDeletes - Forces the system to delete the pending deleted
	// users when using the accidental deletions prevention feature and the deletion threshold is exceeded. Leaving this
	// property empty emulates the Restart provisioning option in the Microsoft Entra admin center. It is similar to setting
	// the resetScope to include QuarantineState, Watermark, and Escrows. This option meets most customer needs.
	ResetScope *SynchronizationJobRestartScope `json:"resetScope,omitempty"`
}
