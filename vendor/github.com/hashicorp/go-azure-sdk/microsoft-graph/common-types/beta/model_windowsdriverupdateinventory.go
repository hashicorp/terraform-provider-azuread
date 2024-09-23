package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsDriverUpdateInventory{}

type WindowsDriverUpdateInventory struct {
	// The number of devices for which this driver is applicable.
	ApplicableDeviceCount *int64 `json:"applicableDeviceCount,omitempty"`

	// An enum type to represent approval status of a driver.
	ApprovalStatus *DriverApprovalStatus `json:"approvalStatus,omitempty"`

	// An enum type to represent which category a driver belongs to.
	Category *DriverCategory `json:"category,omitempty"`

	// The date time when a driver should be deployed if approvalStatus is approved.
	DeployDateTime *string `json:"deployDateTime,omitempty"`

	// The class of the driver.
	DriverClass nullable.Type[string] `json:"driverClass,omitempty"`

	// The manufacturer of the driver.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The name of the driver.
	Name *string `json:"name,omitempty"`

	// The release date time of the driver.
	ReleaseDateTime *string `json:"releaseDateTime,omitempty"`

	// The version of the driver.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s WindowsDriverUpdateInventory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsDriverUpdateInventory{}

func (s WindowsDriverUpdateInventory) MarshalJSON() ([]byte, error) {
	type wrapper WindowsDriverUpdateInventory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsDriverUpdateInventory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDriverUpdateInventory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsDriverUpdateInventory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsDriverUpdateInventory: %+v", err)
	}

	return encoded, nil
}
