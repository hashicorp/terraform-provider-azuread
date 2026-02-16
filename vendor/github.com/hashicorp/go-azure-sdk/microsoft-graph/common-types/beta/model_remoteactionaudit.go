package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RemoteActionAudit{}

type RemoteActionAudit struct {
	// Remote actions Intune supports.
	Action *RemoteAction `json:"action,omitempty"`

	ActionState *ActionState `json:"actionState,omitempty"`

	// BulkAction ID
	BulkDeviceActionId nullable.Type[string] `json:"bulkDeviceActionId,omitempty"`

	// Enum type used for DeviceActionCategory
	DeviceActionCategory *DeviceActionCategory `json:"deviceActionCategory,omitempty"`

	// Intune device name.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// IMEI of the device.
	DeviceIMEI nullable.Type[string] `json:"deviceIMEI,omitempty"`

	// Upn of the device owner.
	DeviceOwnerUserPrincipalName nullable.Type[string] `json:"deviceOwnerUserPrincipalName,omitempty"`

	// User who initiated the device action, format is UPN.
	InitiatedByUserPrincipalName nullable.Type[string] `json:"initiatedByUserPrincipalName,omitempty"`

	// Action target.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// Time when the action was issued, given in UTC.
	RequestDateTime *string `json:"requestDateTime,omitempty"`

	// [deprecated] Please use InitiatedByUserPrincipalName instead.
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s RemoteActionAudit) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RemoteActionAudit{}

func (s RemoteActionAudit) MarshalJSON() ([]byte, error) {
	type wrapper RemoteActionAudit
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RemoteActionAudit: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RemoteActionAudit: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.remoteActionAudit"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RemoteActionAudit: %+v", err)
	}

	return encoded, nil
}
