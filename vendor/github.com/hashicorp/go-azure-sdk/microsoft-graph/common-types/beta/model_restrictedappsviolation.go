package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RestrictedAppsViolation{}

type RestrictedAppsViolation struct {
	// Device configuration profile unique identifier, must be Guid
	DeviceConfigurationId *string `json:"deviceConfigurationId,omitempty"`

	// Device configuration profile name
	DeviceConfigurationName nullable.Type[string] `json:"deviceConfigurationName,omitempty"`

	// Device name
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Managed device unique identifier, must be Guid
	ManagedDeviceId *string `json:"managedDeviceId,omitempty"`

	// Supported platform types for policies.
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`

	// List of violated restricted apps
	RestrictedApps *[]ManagedDeviceReportedApp `json:"restrictedApps,omitempty"`

	// Restricted apps state
	RestrictedAppsState *RestrictedAppsState `json:"restrictedAppsState,omitempty"`

	// User unique identifier, must be Guid
	UserId *string `json:"userId,omitempty"`

	// User name
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s RestrictedAppsViolation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RestrictedAppsViolation{}

func (s RestrictedAppsViolation) MarshalJSON() ([]byte, error) {
	type wrapper RestrictedAppsViolation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RestrictedAppsViolation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RestrictedAppsViolation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.restrictedAppsViolation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RestrictedAppsViolation: %+v", err)
	}

	return encoded, nil
}
