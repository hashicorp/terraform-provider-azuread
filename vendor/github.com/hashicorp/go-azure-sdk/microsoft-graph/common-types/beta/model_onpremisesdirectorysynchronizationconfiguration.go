package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronizationConfiguration struct {
	// Contains the accidental deletion prevention configuration for a tenant.
	AccidentalDeletionPrevention *OnPremisesAccidentalDeletionPrevention `json:"accidentalDeletionPrevention,omitempty"`

	// The anchor attribute allows customers to customize the property used to create source anchors for synchronization
	// enabled objects.
	AnchorAttribute nullable.Type[string] `json:"anchorAttribute,omitempty"`

	// The identifier of the on-premises directory synchronization client application that is configured for the tenant.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// Data for the current export run.
	CurrentExportData *OnPremisesCurrentExportData `json:"currentExportData,omitempty"`

	// Interval of time that the customer requested the sync client waits between sync cycles.
	CustomerRequestedSynchronizationInterval nullable.Type[string] `json:"customerRequestedSynchronizationInterval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the version of the on-premises directory synchronization application.
	SynchronizationClientVersion nullable.Type[string] `json:"synchronizationClientVersion,omitempty"`

	// Interval of time the sync client should honor between sync cycles
	SynchronizationInterval nullable.Type[string] `json:"synchronizationInterval,omitempty"`

	// Configuration to control how cloud created or owned objects are synchronized back to the on-premises directory.
	WritebackConfiguration *OnPremisesWritebackConfiguration `json:"writebackConfiguration,omitempty"`
}
