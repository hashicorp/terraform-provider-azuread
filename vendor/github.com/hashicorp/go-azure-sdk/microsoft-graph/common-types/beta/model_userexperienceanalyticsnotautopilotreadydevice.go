package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsNotAutopilotReadyDevice{}

type UserExperienceAnalyticsNotAutopilotReadyDevice struct {
	// The intune device's autopilotProfileAssigned.
	AutoPilotProfileAssigned *bool `json:"autoPilotProfileAssigned,omitempty"`

	// The intune device's autopilotRegistered.
	AutoPilotRegistered *bool `json:"autoPilotRegistered,omitempty"`

	// The intune device's azure Ad joinType.
	AzureAdJoinType nullable.Type[string] `json:"azureAdJoinType,omitempty"`

	// The intune device's azureAdRegistered.
	AzureAdRegistered nullable.Type[bool] `json:"azureAdRegistered,omitempty"`

	// The intune device's name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The intune device's managed by.
	ManagedBy nullable.Type[string] `json:"managedBy,omitempty"`

	// The intune device's manufacturer.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The intune device's model.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The intune device's serial number.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

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

func (s UserExperienceAnalyticsNotAutopilotReadyDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsNotAutopilotReadyDevice{}

func (s UserExperienceAnalyticsNotAutopilotReadyDevice) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsNotAutopilotReadyDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsNotAutopilotReadyDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsNotAutopilotReadyDevice: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsNotAutopilotReadyDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsNotAutopilotReadyDevice: %+v", err)
	}

	return encoded, nil
}
