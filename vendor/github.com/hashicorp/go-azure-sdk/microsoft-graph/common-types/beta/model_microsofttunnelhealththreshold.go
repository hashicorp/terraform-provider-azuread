package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MicrosoftTunnelHealthThreshold{}

type MicrosoftTunnelHealthThreshold struct {
	// The threshold for being healthy based on default health status metrics: CPU usage healthy < 50%, Memory usage healthy
	// < 50%, Disk space healthy > 5GB, Latency healthy < 10ms, health metrics can be customized. Read-only.
	DefaultHealthyThreshold *int64 `json:"defaultHealthyThreshold,omitempty"`

	// The threshold for being unhealthy based on default health status metrics: CPU usage unhealthy > 75%, Memory usage
	// unhealthy > 75%, Disk space < 3GB, Latency unhealthy > 20ms, health metrics can be customized. Read-only.
	DefaultUnhealthyThreshold *int64 `json:"defaultUnhealthyThreshold,omitempty"`

	// The threshold for being healthy based on default health status metrics: CPU usage healthy < 50%, Memory usage healthy
	// < 50%, Disk space healthy > 5GB, Latency healthy < 10ms, health metrics can be customized.
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty"`

	// The threshold for being unhealthy based on default health status metrics: CPU usage unhealthy > 75%, Memory usage
	// unhealthy > 75%, Disk space < 3GB, Latency Unhealthy > 20ms, health metrics can be customized.
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty"`

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

func (s MicrosoftTunnelHealthThreshold) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftTunnelHealthThreshold{}

func (s MicrosoftTunnelHealthThreshold) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftTunnelHealthThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftTunnelHealthThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftTunnelHealthThreshold: %+v", err)
	}

	delete(decoded, "defaultHealthyThreshold")
	delete(decoded, "defaultUnhealthyThreshold")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftTunnelHealthThreshold"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftTunnelHealthThreshold: %+v", err)
	}

	return encoded, nil
}
