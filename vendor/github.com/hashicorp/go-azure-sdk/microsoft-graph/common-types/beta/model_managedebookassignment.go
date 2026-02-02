package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEBookAssignment interface {
	Entity
	ManagedEBookAssignment() BaseManagedEBookAssignmentImpl
}

var _ ManagedEBookAssignment = BaseManagedEBookAssignmentImpl{}

type BaseManagedEBookAssignmentImpl struct {
	// Possible values for the install intent chosen by the admin.
	InstallIntent *InstallIntent `json:"installIntent,omitempty"`

	// The assignment target for eBook.
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

func (s BaseManagedEBookAssignmentImpl) ManagedEBookAssignment() BaseManagedEBookAssignmentImpl {
	return s
}

func (s BaseManagedEBookAssignmentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedEBookAssignment = RawManagedEBookAssignmentImpl{}

// RawManagedEBookAssignmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedEBookAssignmentImpl struct {
	managedEBookAssignment BaseManagedEBookAssignmentImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawManagedEBookAssignmentImpl) ManagedEBookAssignment() BaseManagedEBookAssignmentImpl {
	return s.managedEBookAssignment
}

func (s RawManagedEBookAssignmentImpl) Entity() BaseEntityImpl {
	return s.managedEBookAssignment.Entity()
}

var _ json.Marshaler = BaseManagedEBookAssignmentImpl{}

func (s BaseManagedEBookAssignmentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedEBookAssignmentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedEBookAssignmentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedEBookAssignmentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedEBookAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedEBookAssignmentImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseManagedEBookAssignmentImpl{}

func (s *BaseManagedEBookAssignmentImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		InstallIntent *InstallIntent `json:"installIntent,omitempty"`
		Id            *string        `json:"id,omitempty"`
		ODataId       *string        `json:"@odata.id,omitempty"`
		ODataType     *string        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.InstallIntent = decoded.InstallIntent
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseManagedEBookAssignmentImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'BaseManagedEBookAssignmentImpl': %+v", err)
		}
		s.Target = impl
	}

	return nil
}

func UnmarshalManagedEBookAssignmentImplementation(input []byte) (ManagedEBookAssignment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedEBookAssignment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppEBookAssignment") {
		var out IosVppEBookAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppEBookAssignment: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedEBookAssignmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedEBookAssignmentImpl: %+v", err)
	}

	return RawManagedEBookAssignmentImpl{
		managedEBookAssignment: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
