package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilter interface {
	Entity
	DeviceAndAppManagementAssignmentFilter() BaseDeviceAndAppManagementAssignmentFilterImpl
}

var _ DeviceAndAppManagementAssignmentFilter = BaseDeviceAndAppManagementAssignmentFilterImpl{}

type BaseDeviceAndAppManagementAssignmentFilterImpl struct {
	// Supported filter management types whether its devices or apps.
	AssignmentFilterManagementType *AssignmentFilterManagementType `json:"assignmentFilterManagementType,omitempty"`

	// The creation time of the assignment filter. The value cannot be modified and is automatically populated during new
	// assignment filter process. The timestamp type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Optional description of the Assignment Filter.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the Assignment Filter.
	DisplayName *string `json:"displayName,omitempty"`

	// Last modified time of the Assignment Filter. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Indicates associated assignments for a specific filter.
	Payloads *[]PayloadByFilter `json:"payloads,omitempty"`

	// Supported platform types.
	Platform *DevicePlatformType `json:"platform,omitempty"`

	// Indicates role scope tags assigned for the assignment filter.
	RoleScopeTags *[]string `json:"roleScopeTags,omitempty"`

	// Rule definition of the assignment filter.
	Rule *string `json:"rule,omitempty"`

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

func (s BaseDeviceAndAppManagementAssignmentFilterImpl) DeviceAndAppManagementAssignmentFilter() BaseDeviceAndAppManagementAssignmentFilterImpl {
	return s
}

func (s BaseDeviceAndAppManagementAssignmentFilterImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceAndAppManagementAssignmentFilter = RawDeviceAndAppManagementAssignmentFilterImpl{}

// RawDeviceAndAppManagementAssignmentFilterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceAndAppManagementAssignmentFilterImpl struct {
	deviceAndAppManagementAssignmentFilter BaseDeviceAndAppManagementAssignmentFilterImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawDeviceAndAppManagementAssignmentFilterImpl) DeviceAndAppManagementAssignmentFilter() BaseDeviceAndAppManagementAssignmentFilterImpl {
	return s.deviceAndAppManagementAssignmentFilter
}

func (s RawDeviceAndAppManagementAssignmentFilterImpl) Entity() BaseEntityImpl {
	return s.deviceAndAppManagementAssignmentFilter.Entity()
}

var _ json.Marshaler = BaseDeviceAndAppManagementAssignmentFilterImpl{}

func (s BaseDeviceAndAppManagementAssignmentFilterImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceAndAppManagementAssignmentFilterImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceAndAppManagementAssignmentFilterImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceAndAppManagementAssignmentFilterImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceAndAppManagementAssignmentFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceAndAppManagementAssignmentFilterImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceAndAppManagementAssignmentFilterImplementation(input []byte) (DeviceAndAppManagementAssignmentFilter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementAssignmentFilter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.payloadCompatibleAssignmentFilter") {
		var out PayloadCompatibleAssignmentFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PayloadCompatibleAssignmentFilter: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceAndAppManagementAssignmentFilterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceAndAppManagementAssignmentFilterImpl: %+v", err)
	}

	return RawDeviceAndAppManagementAssignmentFilterImpl{
		deviceAndAppManagementAssignmentFilter: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
