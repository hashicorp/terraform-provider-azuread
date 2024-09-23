package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsResourcePerformance{}

type UserExperienceAnalyticsResourcePerformance struct {
	// AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
	AverageSpikeTimeScore *int64 `json:"averageSpikeTimeScore,omitempty"`

	// The name of the processor on the device, For example, 11th Gen Intel(R) Core(TM) i7.
	CpuDisplayName nullable.Type[string] `json:"cpuDisplayName,omitempty"`

	// The user experience analytics device CPU spike time score. Valid values 0 to 100
	CpuSpikeTimeScore *int64 `json:"cpuSpikeTimeScore,omitempty"`

	// User experience analytics summarized device count.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The id of the device.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The name of the device.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Resource performance score of a specific device. Valid values 0 to 100
	DeviceResourcePerformanceScore *int64 `json:"deviceResourcePerformanceScore,omitempty"`

	DiskType     *DiskType                           `json:"diskType,omitempty"`
	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// Indicates if machine is physical or virtual. Possible values are: physical or virtual
	MachineType *UserExperienceAnalyticsMachineType `json:"machineType,omitempty"`

	// The user experience analytics device manufacturer.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The user experience analytics device model.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The user experience analytics device RAM spike time score. Valid values 0 to 100
	RamSpikeTimeScore *int64 `json:"ramSpikeTimeScore,omitempty"`

	// The count of cores of the processor of device. Valid values 0 to 512
	TotalProcessorCoreCount *int64 `json:"totalProcessorCoreCount,omitempty"`

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

func (s UserExperienceAnalyticsResourcePerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsResourcePerformance{}

func (s UserExperienceAnalyticsResourcePerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsResourcePerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsResourcePerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsResourcePerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsResourcePerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsResourcePerformance: %+v", err)
	}

	return encoded, nil
}
