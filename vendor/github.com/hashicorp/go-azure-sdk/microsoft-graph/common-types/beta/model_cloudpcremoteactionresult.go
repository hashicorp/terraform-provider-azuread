package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionResult struct {
	// The specified action. Supported values in the Microsoft Endpoint Manager portal are: Reprovision, Resize, Restore.
	// Supported values in enterprise Cloud PC devices are: Reboot, Rename, Reprovision, Troubleshoot.
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	// State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
	ActionState *ActionState `json:"actionState,omitempty"`

	// The ID of the Cloud PC device on which the remote action is performed. Read-only.
	CloudPCId nullable.Type[string] `json:"cloudPcId,omitempty"`

	// Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For
	// example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The ID of the Intune managed device on which the remote action is performed. Read-only.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For
	// example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The extended details of the action status, including error code, error message, and additional information. For
	// example, 'statusDetail': {'code': 'internalServerError','message': 'There was an internal server error. Please
	// contact support xxx.','additionalInformation': [ { '@odata.type':'microsoft.graph.keyValuePair','name':
	// 'correlationId','value': '52367774-cfb7-4e9c-ab51-1b864c31f2d1'} ]}
	StatusDetail *CloudPCStatusDetail `json:"statusDetail,omitempty"`

	// The details of the Cloud PC status. This property is deprecated and will no longer be supported effective August 31,
	// 2024. Use statusDetail instead.
	StatusDetails *CloudPCStatusDetails `json:"statusDetails,omitempty"`
}

var _ json.Marshaler = CloudPCRemoteActionResult{}

func (s CloudPCRemoteActionResult) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCRemoteActionResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCRemoteActionResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCRemoteActionResult: %+v", err)
	}

	delete(decoded, "actionState")
	delete(decoded, "cloudPcId")
	delete(decoded, "managedDeviceId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCRemoteActionResult: %+v", err)
	}

	return encoded, nil
}
