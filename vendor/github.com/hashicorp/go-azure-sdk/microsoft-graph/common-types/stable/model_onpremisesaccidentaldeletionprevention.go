package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAccidentalDeletionPrevention struct {
	// Threshold value which triggers accidental deletion prevention. The threshold is either an absolute number of objects
	// or a percentage number of objects.
	AlertThreshold nullable.Type[int64] `json:"alertThreshold,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The status of the accidental deletion prevention feature. The possible values are: disabled, enabledForCount,
	// enabledForPercentage, unknownFutureValue.
	SynchronizationPreventionType *OnPremisesDirectorySynchronizationDeletionPreventionType `json:"synchronizationPreventionType,omitempty"`
}
