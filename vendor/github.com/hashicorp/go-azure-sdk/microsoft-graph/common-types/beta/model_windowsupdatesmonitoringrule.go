package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringRule struct {
	// The action triggered when the threshold for the given signal is reached. Possible values are: alertError,
	// pauseDeployment, offerFallback, unknownFutureValue. The offerFallback member is only supported on feature update
	// deployments of Windows 11 and must be paired with the ineligible signal. The fallback version offered is the version
	// 22H2 of Windows 10.
	Action *WindowsUpdatesMonitoringAction `json:"action,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The signal to monitor. Possible values are: rollback, ineligible, unknownFutureValue. The ineligible member is only
	// supported on feature update deployments of Windows 11 and must be paired with the offerFallback action.
	Signal *WindowsUpdatesMonitoringSignal `json:"signal,omitempty"`

	// The threshold for a signal at which to trigger the action. An integer from 1 to 100 (inclusive). This value is
	// ignored when the signal is ineligible and the action is offerFallback.
	Threshold nullable.Type[int64] `json:"threshold,omitempty"`
}
