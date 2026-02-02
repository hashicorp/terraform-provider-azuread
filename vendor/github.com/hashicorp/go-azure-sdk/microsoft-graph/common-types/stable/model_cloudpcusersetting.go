package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCUserSetting{}

type CloudPCUserSetting struct {
	// Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting
	// assigned. Returned only on $expand. For an example, see Get cloudPcUserSetting.
	Assignments *[]CloudPCUserSettingAssignment `json:"assignments,omitempty"`

	// The date and time when the setting was created. The timestamp type represents the date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The setting name displayed in the user interface.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the setting was last modified. The timestamp type represents the date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Indicates whether the local admin option is enabled. The default value is false. To enable the local admin option,
	// change the setting to true. If the local admin option is enabled, the end user can be an admin of the Cloud PC
	// device.
	LocalAdminEnabled nullable.Type[bool] `json:"localAdminEnabled,omitempty"`

	// Indicates whether an end user is allowed to reset their Cloud PC. When true, the user is allowed to reset their Cloud
	// PC. When false, end-user initiated reset is not allowed. The default value is false.
	ResetEnabled nullable.Type[bool] `json:"resetEnabled,omitempty"`

	// Defines how frequently a restore point is created that is, a snapshot is taken) for users' provisioned Cloud PCs
	// (default is 12 hours), and whether the user is allowed to restore their own Cloud PCs to a backup made at a specific
	// point in time.
	RestorePointSetting *CloudPCRestorePointSetting `json:"restorePointSetting,omitempty"`

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

func (s CloudPCUserSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCUserSetting{}

func (s CloudPCUserSetting) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCUserSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCUserSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCUserSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcUserSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCUserSetting: %+v", err)
	}

	return encoded, nil
}
