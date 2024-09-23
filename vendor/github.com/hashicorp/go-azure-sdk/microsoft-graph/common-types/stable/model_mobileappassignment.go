package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppAssignment{}

type MobileAppAssignment struct {
	// Possible values for the install intent chosen by the admin.
	Intent *InstallIntent `json:"intent,omitempty"`

	// The settings for target assignment defined by the admin.
	Settings MobileAppAssignmentSettings `json:"settings"`

	// The target group assignment defined by the admin.
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

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

func (s MobileAppAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppAssignment{}

func (s MobileAppAssignment) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MobileAppAssignment{}

func (s *MobileAppAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Intent    *InstallIntent `json:"intent,omitempty"`
		Id        *string        `json:"id,omitempty"`
		ODataId   *string        `json:"@odata.id,omitempty"`
		ODataType *string        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Intent = decoded.Intent
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MobileAppAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settings"]; ok {
		impl, err := UnmarshalMobileAppAssignmentSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Settings' for 'MobileAppAssignment': %+v", err)
		}
		s.Settings = impl
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'MobileAppAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
