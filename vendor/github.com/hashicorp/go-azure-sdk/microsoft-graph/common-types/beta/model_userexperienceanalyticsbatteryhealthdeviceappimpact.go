package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthDeviceAppImpact{}

type UserExperienceAnalyticsBatteryHealthDeviceAppImpact struct {
	// User friendly display name for the app. Eg: Outlook
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// App name. Eg: oltk.exe
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// App publisher. Eg: Microsoft Corporation
	AppPublisher nullable.Type[string] `json:"appPublisher,omitempty"`

	// The unique identifier of the device, Intune DeviceID or SCCM device id.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// true if the user had active interaction with the app.
	IsForegroundApp *bool `json:"isForegroundApp,omitempty"`

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

func (s UserExperienceAnalyticsBatteryHealthDeviceAppImpact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthDeviceAppImpact{}

func (s UserExperienceAnalyticsBatteryHealthDeviceAppImpact) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthDeviceAppImpact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthDeviceAppImpact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthDeviceAppImpact: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthDeviceAppImpact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthDeviceAppImpact: %+v", err)
	}

	return encoded, nil
}
