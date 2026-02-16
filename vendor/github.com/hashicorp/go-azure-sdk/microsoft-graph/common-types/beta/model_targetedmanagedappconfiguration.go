package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedAppConfiguration = TargetedManagedAppConfiguration{}

type TargetedManagedAppConfiguration struct {
	// Indicates a collection of apps to target which can be one of several pre-defined lists of apps or a manually selected
	// list of apps
	AppGroupType *TargetedManagedAppGroupType `json:"appGroupType,omitempty"`

	// List of apps to which the policy is deployed.
	Apps *[]ManagedMobileApp `json:"apps,omitempty"`

	// Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
	Assignments *[]TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`

	// Count of apps to which the current policy is deployed.
	DeployedAppCount *int64 `json:"deployedAppCount,omitempty"`

	// Navigation property to deployment summary of the configuration.
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`

	// Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// Management levels for apps
	TargetedAppManagementLevels *AppManagementLevel `json:"targetedAppManagementLevels,omitempty"`

	// Fields inherited from ManagedAppConfiguration

	// A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped,
	// unalterned by this service
	CustomSettings *[]KeyValuePair `json:"customSettings,omitempty"`

	// List of settings contained in this App Configuration policy
	Settings *[]DeviceManagementConfigurationSetting `json:"settings,omitempty"`

	// Fields inherited from ManagedAppPolicy

	// The date and time the policy was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The policy's description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Last time the policy was modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s TargetedManagedAppConfiguration) ManagedAppConfiguration() BaseManagedAppConfigurationImpl {
	return BaseManagedAppConfigurationImpl{
		CustomSettings:       s.CustomSettings,
		Settings:             s.Settings,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s TargetedManagedAppConfiguration) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return BaseManagedAppPolicyImpl{
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s TargetedManagedAppConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TargetedManagedAppConfiguration{}

func (s TargetedManagedAppConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper TargetedManagedAppConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TargetedManagedAppConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TargetedManagedAppConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.targetedManagedAppConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TargetedManagedAppConfiguration: %+v", err)
	}

	return encoded, nil
}
