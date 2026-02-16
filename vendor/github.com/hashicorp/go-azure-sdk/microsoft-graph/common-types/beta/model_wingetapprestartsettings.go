package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppRestartSettings struct {
	// The number of minutes before the restart time to display the countdown dialog for pending restarts.
	CountdownDisplayBeforeRestartInMinutes *int64 `json:"countdownDisplayBeforeRestartInMinutes,omitempty"`

	// The number of minutes to wait before restarting the device after an app installation.
	GracePeriodInMinutes *int64 `json:"gracePeriodInMinutes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of minutes to snooze the restart notification dialog when the snooze button is selected.
	RestartNotificationSnoozeDurationInMinutes nullable.Type[int64] `json:"restartNotificationSnoozeDurationInMinutes,omitempty"`
}
