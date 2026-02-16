package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceAndAppManagementAssignmentFilter = PayloadCompatibleAssignmentFilter{}

type PayloadCompatibleAssignmentFilter struct {
	// Represents the payload type AssignmentFilter is being assigned to.
	PayloadType *AssignmentFilterPayloadType `json:"payloadType,omitempty"`

	// Fields inherited from DeviceAndAppManagementAssignmentFilter

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

func (s PayloadCompatibleAssignmentFilter) DeviceAndAppManagementAssignmentFilter() BaseDeviceAndAppManagementAssignmentFilterImpl {
	return BaseDeviceAndAppManagementAssignmentFilterImpl{
		AssignmentFilterManagementType: s.AssignmentFilterManagementType,
		CreatedDateTime:                s.CreatedDateTime,
		Description:                    s.Description,
		DisplayName:                    s.DisplayName,
		LastModifiedDateTime:           s.LastModifiedDateTime,
		Payloads:                       s.Payloads,
		Platform:                       s.Platform,
		RoleScopeTags:                  s.RoleScopeTags,
		Rule:                           s.Rule,
		Id:                             s.Id,
		ODataId:                        s.ODataId,
		ODataType:                      s.ODataType,
	}
}

func (s PayloadCompatibleAssignmentFilter) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PayloadCompatibleAssignmentFilter{}

func (s PayloadCompatibleAssignmentFilter) MarshalJSON() ([]byte, error) {
	type wrapper PayloadCompatibleAssignmentFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PayloadCompatibleAssignmentFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PayloadCompatibleAssignmentFilter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.payloadCompatibleAssignmentFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PayloadCompatibleAssignmentFilter: %+v", err)
	}

	return encoded, nil
}
