package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsQualityUpdateProfile{}

type WindowsQualityUpdateProfile struct {
	// The list of group assignments of the profile.
	Assignments *[]WindowsQualityUpdateProfileAssignment `json:"assignments,omitempty"`

	// The date time that the profile was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Friendly display name of the quality update profile deployable content
	DeployableContentDisplayName nullable.Type[string] `json:"deployableContentDisplayName,omitempty"`

	// The description of the profile which is specified by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the profile.
	DisplayName *string `json:"displayName,omitempty"`

	// Expedited update settings.
	ExpeditedUpdateSettings *ExpeditedWindowsQualityUpdateSettings `json:"expeditedUpdateSettings,omitempty"`

	// The date time that the profile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Friendly release date to display for a Quality Update release
	ReleaseDateDisplayName nullable.Type[string] `json:"releaseDateDisplayName,omitempty"`

	// List of Scope Tags for this Quality Update entity.
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

func (s WindowsQualityUpdateProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsQualityUpdateProfile{}

func (s WindowsQualityUpdateProfile) MarshalJSON() ([]byte, error) {
	type wrapper WindowsQualityUpdateProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsQualityUpdateProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsQualityUpdateProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsQualityUpdateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsQualityUpdateProfile: %+v", err)
	}

	return encoded, nil
}
