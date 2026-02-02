package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkFeaturesConfiguration struct {
	// Email address to send logs and feedback.
	EmailToSendLogsAndFeedback nullable.Type[string] `json:"emailToSendLogsAndFeedback,omitempty"`

	// True if auto screen shared is enabled.
	IsAutoScreenShareEnabled nullable.Type[bool] `json:"isAutoScreenShareEnabled,omitempty"`

	// True if Bluetooth beaconing is enabled.
	IsBluetoothBeaconingEnabled nullable.Type[bool] `json:"isBluetoothBeaconingEnabled,omitempty"`

	// True if hiding meeting names is enabled.
	IsHideMeetingNamesEnabled nullable.Type[bool] `json:"isHideMeetingNamesEnabled,omitempty"`

	// True if sending logs and feedback is enabled.
	IsSendLogsAndFeedbackEnabled nullable.Type[bool] `json:"isSendLogsAndFeedbackEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
