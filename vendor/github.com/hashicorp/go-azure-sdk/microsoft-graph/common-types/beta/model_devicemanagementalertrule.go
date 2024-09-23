package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementAlertRule{}

type DeviceManagementAlertRule struct {
	// The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario,
	// cloudPcOnPremiseNetworkConnectionCheckScenario, cloudPcInGracePeriodScenario,
	// cloudPcFrontlineInsufficientLicensesScenario, cloudPcInaccessibleScenario. Note that you must use the Prefer:
	// include-unknown-enum-members request header to get the following values from this evolvable enum:
	// cloudPcInGracePeriodScenario.
	AlertRuleTemplate *DeviceManagementAlertRuleTemplate `json:"alertRuleTemplate,omitempty"`

	// The conditions that determine when to send alerts. For example, you can configure a condition to send an alert when
	// provisioning fails for six or more Cloud PCs.
	Conditions *[]DeviceManagementRuleCondition `json:"conditions,omitempty"`

	// The rule description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the rule.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The status of the rule that indicates whether the rule is enabled or disabled. If true, the rule is enabled;
	// otherwise, the rule is disabled.
	Enabled nullable.Type[bool] `json:"enabled,omitempty"`

	// Indicates whether the rule is a system rule. If true, the rule is a system rule; otherwise, the rule is a
	// custom-defined rule and can be edited. System rules are built in and only a few properties can be edited.
	IsSystemRule nullable.Type[bool] `json:"isSystemRule,omitempty"`

	// The notification channels of the rule selected by the user.
	NotificationChannels *[]DeviceManagementNotificationChannel `json:"notificationChannels,omitempty"`

	// The severity of the rule. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
	Severity *DeviceManagementRuleSeverityType `json:"severity,omitempty"`

	// The conditions that determine when to send alerts. For example, you can configure a condition to send an alert when
	// provisioning fails for six or more Cloud PCs. This property is deprecated. Use conditions instead.
	Threshold *DeviceManagementRuleThreshold `json:"threshold,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementAlertRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementAlertRule{}

func (s DeviceManagementAlertRule) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementAlertRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementAlertRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementAlertRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagement.alertRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementAlertRule: %+v", err)
	}

	return encoded, nil
}
