package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementPartner{}

type DeviceManagementPartner struct {
	// Partner display name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// User groups that specifies whether enrollment is through partner.
	GroupsRequiringPartnerEnrollment *[]DeviceManagementPartnerAssignment `json:"groupsRequiringPartnerEnrollment,omitempty"`

	// Whether device management partner is configured or not
	IsConfigured *bool `json:"isConfigured,omitempty"`

	// Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
	LastHeartbeatDateTime *string `json:"lastHeartbeatDateTime,omitempty"`

	// Partner App Type.
	PartnerAppType *DeviceManagementPartnerAppType `json:"partnerAppType,omitempty"`

	// Partner state of this tenant.
	PartnerState *DeviceManagementPartnerTenantState `json:"partnerState,omitempty"`

	// Partner Single tenant App id
	SingleTenantAppId nullable.Type[string] `json:"singleTenantAppId,omitempty"`

	// DateTime in UTC when PartnerDevices will be marked as NonCompliant
	WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime nullable.Type[string] `json:"whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime,omitempty"`

	// DateTime in UTC when PartnerDevices will be removed
	WhenPartnerDevicesWillBeRemovedDateTime nullable.Type[string] `json:"whenPartnerDevicesWillBeRemovedDateTime,omitempty"`

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

func (s DeviceManagementPartner) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementPartner{}

func (s DeviceManagementPartner) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementPartner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementPartner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementPartner: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementPartner"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementPartner: %+v", err)
	}

	return encoded, nil
}
