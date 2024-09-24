package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthAppImpact{}

type UserExperienceAnalyticsBatteryHealthAppImpact struct {
	// Number of active devices for using that app over a 14-day period. Valid values 0 to 2147483647
	ActiveDevices *int64 `json:"activeDevices,omitempty"`

	// User friendly display name for the app. Eg: Outlook
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// App name. Eg: oltk.exe
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// App publisher. Eg: Microsoft Corporation
	AppPublisher nullable.Type[string] `json:"appPublisher,omitempty"`

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

func (s UserExperienceAnalyticsBatteryHealthAppImpact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthAppImpact{}

func (s UserExperienceAnalyticsBatteryHealthAppImpact) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthAppImpact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthAppImpact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthAppImpact: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthAppImpact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthAppImpact: %+v", err)
	}

	return encoded, nil
}
