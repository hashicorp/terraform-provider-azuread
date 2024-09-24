package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GovernanceRoleSetting{}

type GovernanceRoleSetting struct {
	// The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
	AdminEligibleSettings *[]GovernanceRuleSetting `json:"adminEligibleSettings,omitempty"`

	// The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
	AdminMemberSettings *[]GovernanceRuleSetting `json:"adminMemberSettings,omitempty"`

	// Read-only. Indicate if the roleSetting is a default roleSetting
	IsDefault nullable.Type[bool] `json:"isDefault,omitempty"`

	// Read-only. The display name of the administrator who last updated the roleSetting.
	LastUpdatedBy nullable.Type[string] `json:"lastUpdatedBy,omitempty"`

	// Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// Read-only. The associated resource for this role setting.
	Resource *GovernanceResource `json:"resource,omitempty"`

	// Required. The id of the resource that the role setting is associated with.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Read-only. The role definition that is enforced with this role setting.
	RoleDefinition *GovernanceRoleDefinition `json:"roleDefinition,omitempty"`

	// Required. The id of the role definition that the role setting is associated with.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not
	// supported for now.
	UserEligibleSettings *[]GovernanceRuleSetting `json:"userEligibleSettings,omitempty"`

	// The rule settings that are evaluated when a user tries to activate his role assignment.
	UserMemberSettings *[]GovernanceRuleSetting `json:"userMemberSettings,omitempty"`

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

func (s GovernanceRoleSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GovernanceRoleSetting{}

func (s GovernanceRoleSetting) MarshalJSON() ([]byte, error) {
	type wrapper GovernanceRoleSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GovernanceRoleSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceRoleSetting: %+v", err)
	}

	delete(decoded, "isDefault")
	delete(decoded, "lastUpdatedBy")
	delete(decoded, "lastUpdatedDateTime")
	delete(decoded, "resource")
	delete(decoded, "roleDefinition")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceRoleSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GovernanceRoleSetting: %+v", err)
	}

	return encoded, nil
}
