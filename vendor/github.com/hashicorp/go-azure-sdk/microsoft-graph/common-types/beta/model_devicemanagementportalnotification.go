package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPortalNotification struct {
	// The associated alert impact.
	AlertImpact *DeviceManagementAlertImpact `json:"alertImpact,omitempty"`

	// The associated alert record ID.
	AlertRecordId nullable.Type[string] `json:"alertRecordId,omitempty"`

	// The associated alert rule ID.
	AlertRuleId nullable.Type[string] `json:"alertRuleId,omitempty"`

	// The associated alert rule name.
	AlertRuleName nullable.Type[string] `json:"alertRuleName,omitempty"`

	// The associated alert rule template. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario,
	// cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue, cloudPcInGracePeriodScenario. Use the Prefer:
	// include-unknown-enum-members request header to get the following values from this evolvable enum:
	// cloudPcInGracePeriodScenario.
	AlertRuleTemplate *DeviceManagementAlertRuleTemplate `json:"alertRuleTemplate,omitempty"`

	// The unique identifier for the portal notification.
	Id nullable.Type[string] `json:"id,omitempty"`

	// true if the portal notification has already been sent to the user; false otherwise.
	IsPortalNotificationSent nullable.Type[bool] `json:"isPortalNotificationSent,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The associated alert rule severity. The possible values are: unknown, informational, warning, critical,
	// unknownFutureValue.
	Severity *DeviceManagementRuleSeverityType `json:"severity,omitempty"`
}
