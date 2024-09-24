package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsRemoteConnection{}

type UserExperienceAnalyticsRemoteConnection struct {
	// The count of remote connection. Valid values 0 to 2147483647
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The id of the device.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The name of the device.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The user experience analytics manufacturer.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The user experience analytics device model.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The user experience analytics userPrincipalName.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The user experience analytics virtual network.
	VirtualNetwork nullable.Type[string] `json:"virtualNetwork,omitempty"`

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

func (s UserExperienceAnalyticsRemoteConnection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsRemoteConnection{}

func (s UserExperienceAnalyticsRemoteConnection) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsRemoteConnection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsRemoteConnection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsRemoteConnection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsRemoteConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsRemoteConnection: %+v", err)
	}

	return encoded, nil
}
