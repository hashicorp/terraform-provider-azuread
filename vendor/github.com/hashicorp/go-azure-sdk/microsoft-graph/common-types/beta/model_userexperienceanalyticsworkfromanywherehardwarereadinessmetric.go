package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{}

type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric struct {
	// The count of total devices in an organization. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only.
	// Valid values -2147483648 to 2147483647
	TotalDeviceCount *int64 `json:"totalDeviceCount,omitempty"`

	// The count of devices in an organization eligible for windows upgrade. Valid values 0 to 2147483647. Supports:
	// $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
	UpgradeEligibleDeviceCount *int64 `json:"upgradeEligibleDeviceCount,omitempty"`

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

func (s UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{}

func (s UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric: %+v", err)
	}

	delete(decoded, "totalDeviceCount")
	delete(decoded, "upgradeEligibleDeviceCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric: %+v", err)
	}

	return encoded, nil
}
