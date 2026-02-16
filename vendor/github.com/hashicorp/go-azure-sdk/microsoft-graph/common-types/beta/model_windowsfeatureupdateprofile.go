package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsFeatureUpdateProfile{}

type WindowsFeatureUpdateProfile struct {
	// The list of group assignments of the profile.
	Assignments *[]WindowsFeatureUpdateProfileAssignment `json:"assignments,omitempty"`

	// The date time that the profile was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Friendly display name of the quality update profile deployable content
	DeployableContentDisplayName nullable.Type[string] `json:"deployableContentDisplayName,omitempty"`

	// The description of the profile which is specified by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the profile.
	DisplayName *string `json:"displayName,omitempty"`

	// The last supported date for a feature update
	EndOfSupportDate nullable.Type[string] `json:"endOfSupportDate,omitempty"`

	// The feature update version that will be deployed to the devices targeted by this profile. The version could be any
	// supported version for example 1709, 1803 or 1809 and so on.
	FeatureUpdateVersion *string `json:"featureUpdateVersion,omitempty"`

	// If true, the Windows 11 update will become optional
	InstallFeatureUpdatesOptional *bool `json:"installFeatureUpdatesOptional,omitempty"`

	// If true, the latest Microsoft Windows 10 update will be installed on devices ineligible for Microsoft Windows 11
	InstallLatestWindows10OnWindows11IneligibleDevice *bool `json:"installLatestWindows10OnWindows11IneligibleDevice,omitempty"`

	// The date time that the profile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Feature Update entity.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The windows update rollout settings, including offer start date time, offer end date time, and days between each set
	// of offers.
	RolloutSettings *WindowsUpdateRolloutSettings `json:"rolloutSettings,omitempty"`

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

func (s WindowsFeatureUpdateProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsFeatureUpdateProfile{}

func (s WindowsFeatureUpdateProfile) MarshalJSON() ([]byte, error) {
	type wrapper WindowsFeatureUpdateProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsFeatureUpdateProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsFeatureUpdateProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsFeatureUpdateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsFeatureUpdateProfile: %+v", err)
	}

	return encoded, nil
}
