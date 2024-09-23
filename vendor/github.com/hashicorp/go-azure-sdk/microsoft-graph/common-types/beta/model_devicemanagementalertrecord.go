package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementAlertRecord{}

type DeviceManagementAlertRecord struct {
	// The impact of the alert event. Consists of a list of key-value pair and a number followed by the aggregation type.
	// For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. 12 affectedCloudPcPercentage means 12% of
	// Cloud PCs are affected. The list of key-value pair indicates the details of the alert impact.
	AlertImpact *DeviceManagementAlertImpact `json:"alertImpact,omitempty"`

	// The corresponding ID of the alert rule.
	AlertRuleId nullable.Type[string] `json:"alertRuleId,omitempty"`

	// The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario,
	// cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue, cloudPcInGracePeriodScenario. Note that you must
	// use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum:
	// cloudPcInGracePeriodScenario.
	AlertRuleTemplate *DeviceManagementAlertRuleTemplate `json:"alertRuleTemplate,omitempty"`

	// The date and time when the alert event was detected. The Timestamp type represents date and time information using
	// ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	DetectedDateTime nullable.Type[string] `json:"detectedDateTime,omitempty"`

	// The display name of the alert record.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the alert record was last updated. The Timestamp type represents date and time information
	// using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The date and time when the alert event was resolved. The Timestamp type represents date and time information using
	// ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ResolvedDateTime nullable.Type[string] `json:"resolvedDateTime,omitempty"`

	// The severity of the alert event. The possible values are: unknown, informational, warning, critical,
	// unknownFutureValue.
	Severity *DeviceManagementRuleSeverityType `json:"severity,omitempty"`

	// The status of the alert record. The possible values are: active, resolved, unknownFutureValue.
	Status *DeviceManagementAlertStatusType `json:"status,omitempty"`

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

func (s DeviceManagementAlertRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementAlertRecord{}

func (s DeviceManagementAlertRecord) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementAlertRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementAlertRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementAlertRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagement.alertRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementAlertRecord: %+v", err)
	}

	return encoded, nil
}
