package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsQualityUpdatePolicy{}

type WindowsQualityUpdatePolicy struct {
	// List of the groups this profile is assgined to.
	Assignments *[]WindowsQualityUpdatePolicyAssignment `json:"assignments,omitempty"`

	// Timestamp of when the profile was created. The value cannot be modified and is automatically populated when the
	// profile is created. The Timestamp type represents date and time information using ISO 8601 format and is always in
	// UTC time. Read-only
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the policy which is specified by the user. Max allowed length is 1500 chars.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the policy. Max allowed length is 200 chars.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates if hotpatch is enabled for the tenants. When 'true', tenant can apply quality updates without rebooting
	// their devices. When 'false', tenant devices will receive cold patch associated with Windows quality updates.
	HotpatchEnabled *bool `json:"hotpatchEnabled,omitempty"`

	// Timestamp of when the profile was modified. The value cannot be modified and is automatically populated when the
	// profile is modified. The Timestamp type represents date and time information using ISO 8601 format and is always in
	// UTC time. Read-only
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of the scope tag ids for this profile.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s WindowsQualityUpdatePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsQualityUpdatePolicy{}

func (s WindowsQualityUpdatePolicy) MarshalJSON() ([]byte, error) {
	type wrapper WindowsQualityUpdatePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsQualityUpdatePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsQualityUpdatePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsQualityUpdatePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsQualityUpdatePolicy: %+v", err)
	}

	return encoded, nil
}
