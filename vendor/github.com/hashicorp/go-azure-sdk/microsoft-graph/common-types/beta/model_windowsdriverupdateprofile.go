package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDriverUpdateProfile{}

type WindowsDriverUpdateProfile struct {
	// An enum type to represent approval type of a driver update profile.
	ApprovalType *DriverUpdateProfileApprovalType `json:"approvalType,omitempty"`

	// The list of group assignments of the profile.
	Assignments *[]WindowsDriverUpdateProfileAssignment `json:"assignments,omitempty"`

	// The date time that the profile was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
	DeploymentDeferralInDays nullable.Type[int64] `json:"deploymentDeferralInDays,omitempty"`

	// The description of the profile which is specified by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Number of devices reporting for this profile
	DeviceReporting nullable.Type[int64] `json:"deviceReporting,omitempty"`

	// The display name for the profile.
	DisplayName *string `json:"displayName,omitempty"`

	// Driver inventories for this profile.
	DriverInventories *[]WindowsDriverUpdateInventory `json:"driverInventories,omitempty"`

	// Driver inventory sync status for this profile.
	InventorySyncStatus *WindowsDriverUpdateProfileInventorySyncStatus `json:"inventorySyncStatus,omitempty"`

	// The date time that the profile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Number of new driver updates available for this profile.
	NewUpdates *int64 `json:"newUpdates,omitempty"`

	// List of Scope Tags for this Driver Update entity.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s WindowsDriverUpdateProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDriverUpdateProfile{}

func (s WindowsDriverUpdateProfile) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDriverUpdateProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDriverUpdateProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDriverUpdateProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDriverUpdateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDriverUpdateProfile: %+v", err)
	}

	return encoded, nil
}
