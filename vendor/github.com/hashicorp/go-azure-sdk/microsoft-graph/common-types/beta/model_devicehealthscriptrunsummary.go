package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceHealthScriptRunSummary{}

type DeviceHealthScriptRunSummary struct {
	// Number of devices on which the detection script execution encountered an error and did not complete
	DetectionScriptErrorDeviceCount *int64 `json:"detectionScriptErrorDeviceCount,omitempty"`

	// Number of devices for which the detection script was not applicable
	DetectionScriptNotApplicableDeviceCount *int64 `json:"detectionScriptNotApplicableDeviceCount,omitempty"`

	// Number of devices which have not yet run the latest version of the device health script
	DetectionScriptPendingDeviceCount *int64 `json:"detectionScriptPendingDeviceCount,omitempty"`

	// Number of devices for which the detection script found an issue
	IssueDetectedDeviceCount *int64 `json:"issueDetectedDeviceCount,omitempty"`

	// Number of devices that were remediated over the last 30 days
	IssueRemediatedCumulativeDeviceCount *int64 `json:"issueRemediatedCumulativeDeviceCount,omitempty"`

	// Number of devices for which the remediation script was able to resolve the detected issue
	IssueRemediatedDeviceCount *int64 `json:"issueRemediatedDeviceCount,omitempty"`

	// Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
	IssueReoccurredDeviceCount *int64 `json:"issueReoccurredDeviceCount,omitempty"`

	// Last run time for the script across all devices
	LastScriptRunDateTime nullable.Type[string] `json:"lastScriptRunDateTime,omitempty"`

	// Number of devices for which the detection script did not find an issue and the device is healthy
	NoIssueDetectedDeviceCount *int64 `json:"noIssueDetectedDeviceCount,omitempty"`

	// Number of devices for which the remediation script execution encountered an error and did not complete
	RemediationScriptErrorDeviceCount *int64 `json:"remediationScriptErrorDeviceCount,omitempty"`

	// Number of devices for which remediation was skipped
	RemediationSkippedDeviceCount *int64 `json:"remediationSkippedDeviceCount,omitempty"`

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

func (s DeviceHealthScriptRunSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptRunSummary{}

func (s DeviceHealthScriptRunSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptRunSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptRunSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptRunSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptRunSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptRunSummary: %+v", err)
	}

	return encoded, nil
}
