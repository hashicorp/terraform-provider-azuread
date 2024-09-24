package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceComplianceScriptRunSummary{}

type DeviceComplianceScriptRunSummary struct {
	// Number of devices on which the detection script execution encountered an error and did not complete. Valid values
	// -2147483648 to 2147483647
	DetectionScriptErrorDeviceCount *int64 `json:"detectionScriptErrorDeviceCount,omitempty"`

	// Number of devices which have not yet run the latest version of the device compliance script. Valid values -2147483648
	// to 2147483647
	DetectionScriptPendingDeviceCount *int64 `json:"detectionScriptPendingDeviceCount,omitempty"`

	// Number of devices for which the detection script found an issue. Valid values -2147483648 to 2147483647
	IssueDetectedDeviceCount *int64 `json:"issueDetectedDeviceCount,omitempty"`

	// Last run time for the script across all devices
	LastScriptRunDateTime nullable.Type[string] `json:"lastScriptRunDateTime,omitempty"`

	// Number of devices for which the detection script did not find an issue and the device is healthy. Valid values
	// -2147483648 to 2147483647
	NoIssueDetectedDeviceCount *int64 `json:"noIssueDetectedDeviceCount,omitempty"`

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

func (s DeviceComplianceScriptRunSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceComplianceScriptRunSummary{}

func (s DeviceComplianceScriptRunSummary) MarshalJSON() ([]byte, error) {
	type wrapper DeviceComplianceScriptRunSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceComplianceScriptRunSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceComplianceScriptRunSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceComplianceScriptRunSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceComplianceScriptRunSummary: %+v", err)
	}

	return encoded, nil
}
