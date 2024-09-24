package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicyMigrationReport{}

type GroupPolicyMigrationReport struct {
	// The date and time at which the GroupPolicyMigrationReport was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The name of Group Policy Object from the GPO Xml Content
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time at which the GroupPolicyMigrationReport was created.
	GroupPolicyCreatedDateTime *string `json:"groupPolicyCreatedDateTime,omitempty"`

	// The date and time at which the GroupPolicyMigrationReport was last modified.
	GroupPolicyLastModifiedDateTime *string `json:"groupPolicyLastModifiedDateTime,omitempty"`

	// The Group Policy Object GUID from GPO Xml content
	GroupPolicyObjectId *string `json:"groupPolicyObjectId,omitempty"`

	// A list of group policy settings to MDM/Intune mappings.
	GroupPolicySettingMappings *[]GroupPolicySettingMapping `json:"groupPolicySettingMappings,omitempty"`

	// The date and time at which the GroupPolicyMigrationReport was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Indicates if the Group Policy Object file is covered and ready for Intune migration.
	MigrationReadiness *GroupPolicyMigrationReadiness `json:"migrationReadiness,omitempty"`

	// The distinguished name of the OU.
	OuDistinguishedName nullable.Type[string] `json:"ouDistinguishedName,omitempty"`

	// The list of scope tags for the configuration.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The number of Group Policy Settings supported by Intune.
	SupportedSettingsCount *int64 `json:"supportedSettingsCount,omitempty"`

	// The Percentage of Group Policy Settings supported by Intune.
	SupportedSettingsPercent *int64 `json:"supportedSettingsPercent,omitempty"`

	// The Targeted in AD property from GPO Xml Content
	TargetedInActiveDirectory *bool `json:"targetedInActiveDirectory,omitempty"`

	// The total number of Group Policy Settings from GPO file.
	TotalSettingsCount *int64 `json:"totalSettingsCount,omitempty"`

	// A list of unsupported group policy extensions inside the Group Policy Object.
	UnsupportedGroupPolicyExtensions *[]UnsupportedGroupPolicyExtension `json:"unsupportedGroupPolicyExtensions,omitempty"`

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

func (s GroupPolicyMigrationReport) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyMigrationReport{}

func (s GroupPolicyMigrationReport) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyMigrationReport
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyMigrationReport: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyMigrationReport: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyMigrationReport"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyMigrationReport: %+v", err)
	}

	return encoded, nil
}
