package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementComplianceActionItem{}

type DeviceManagementComplianceActionItem struct {
	// Scheduled Action Type Enum
	ActionType *DeviceManagementComplianceActionType `json:"actionType,omitempty"`

	// Number of hours to wait till the action will be enforced. Valid values 0 to 8760
	GracePeriodHours *int64 `json:"gracePeriodHours,omitempty"`

	// A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100
	// elements.
	NotificationMessageCCList *[]string `json:"notificationMessageCCList,omitempty"`

	// What notification Message template to use
	NotificationTemplateId nullable.Type[string] `json:"notificationTemplateId,omitempty"`

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

func (s DeviceManagementComplianceActionItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementComplianceActionItem{}

func (s DeviceManagementComplianceActionItem) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementComplianceActionItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementComplianceActionItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementComplianceActionItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementComplianceActionItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementComplianceActionItem: %+v", err)
	}

	return encoded, nil
}
