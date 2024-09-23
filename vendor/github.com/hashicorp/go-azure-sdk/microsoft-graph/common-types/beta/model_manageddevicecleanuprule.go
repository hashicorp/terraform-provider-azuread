package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceCleanupRule{}

type ManagedDeviceCleanupRule struct {
	// Indicates the description for the device clean up rule.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Define the platform type for which the admin wants to create the device clean up rule
	DeviceCleanupRulePlatformType *DeviceCleanupRulePlatformType `json:"deviceCleanupRulePlatformType,omitempty"`

	// Indicates the number of days when the device has not contacted Intune. Valid values 0 to 2147483647
	DeviceInactivityBeforeRetirementInDays *int64 `json:"deviceInactivityBeforeRetirementInDays,omitempty"`

	// Indicates the display name of the device cleanup rule.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates the date and time when the device cleanup rule was last modified. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s ManagedDeviceCleanupRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceCleanupRule{}

func (s ManagedDeviceCleanupRule) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceCleanupRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceCleanupRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceCleanupRule: %+v", err)
	}

	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceCleanupRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceCleanupRule: %+v", err)
	}

	return encoded, nil
}
