package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkAction interface {
	Entity
	CloudPCBulkAction() BaseCloudPCBulkActionImpl
}

var _ CloudPCBulkAction = BaseCloudPCBulkActionImpl{}

type BaseCloudPCBulkActionImpl struct {
	// Run summary of this bulk action.
	ActionSummary *CloudPCBulkActionSummary `json:"actionSummary,omitempty"`

	CloudPCIds *[]string `json:"cloudPcIds,omitempty"`

	// The date and time when the bulk action was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Name of the bulk action.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates the user principal name (UPN) of the user who initiated this bulk action. Read-only.
	InitiatedByUserPrincipalName nullable.Type[string] `json:"initiatedByUserPrincipalName,omitempty"`

	// Indicates whether the bulk action is scheduled according to the maintenance window. When true, the bulk action uses
	// the maintenance window to schedule the action; false means that the bulk action doesn't use the maintenance window.
	// The default value is false.
	ScheduledDuringMaintenanceWindow nullable.Type[bool] `json:"scheduledDuringMaintenanceWindow,omitempty"`

	// Indicates the status of bulk actions. Possible values are pending, succeeded, failed, unknownFutureValue. The default
	// value is pending. Read-only.
	Status *CloudPCBulkActionStatus `json:"status,omitempty"`

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

func (s BaseCloudPCBulkActionImpl) CloudPCBulkAction() BaseCloudPCBulkActionImpl {
	return s
}

func (s BaseCloudPCBulkActionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CloudPCBulkAction = RawCloudPCBulkActionImpl{}

// RawCloudPCBulkActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCloudPCBulkActionImpl struct {
	cloudPCBulkAction BaseCloudPCBulkActionImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawCloudPCBulkActionImpl) CloudPCBulkAction() BaseCloudPCBulkActionImpl {
	return s.cloudPCBulkAction
}

func (s RawCloudPCBulkActionImpl) Entity() BaseEntityImpl {
	return s.cloudPCBulkAction.Entity()
}

var _ json.Marshaler = BaseCloudPCBulkActionImpl{}

func (s BaseCloudPCBulkActionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCloudPCBulkActionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCloudPCBulkActionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCloudPCBulkActionImpl: %+v", err)
	}

	delete(decoded, "initiatedByUserPrincipalName")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcBulkAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCloudPCBulkActionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalCloudPCBulkActionImplementation(input []byte) (CloudPCBulkAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCBulkAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkDisasterRecoveryFailback") {
		var out CloudPCBulkDisasterRecoveryFailback
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkDisasterRecoveryFailback: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkDisasterRecoveryFailover") {
		var out CloudPCBulkDisasterRecoveryFailover
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkDisasterRecoveryFailover: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkModifyDiskEncryptionType") {
		var out CloudPCBulkModifyDiskEncryptionType
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkModifyDiskEncryptionType: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkPowerOff") {
		var out CloudPCBulkPowerOff
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkPowerOff: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkPowerOn") {
		var out CloudPCBulkPowerOn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkPowerOn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkReprovision") {
		var out CloudPCBulkReprovision
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkReprovision: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkResize") {
		var out CloudPCBulkResize
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkResize: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkRestart") {
		var out CloudPCBulkRestart
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkRestart: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkRestore") {
		var out CloudPCBulkRestore
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkRestore: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudPcBulkTroubleshoot") {
		var out CloudPCBulkTroubleshoot
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudPCBulkTroubleshoot: %+v", err)
		}
		return out, nil
	}

	var parent BaseCloudPCBulkActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCloudPCBulkActionImpl: %+v", err)
	}

	return RawCloudPCBulkActionImpl{
		cloudPCBulkAction: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
