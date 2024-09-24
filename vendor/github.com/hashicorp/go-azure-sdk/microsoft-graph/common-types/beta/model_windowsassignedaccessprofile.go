package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsAssignedAccessProfile{}

type WindowsAssignedAccessProfile struct {
	// These are the only Windows Store Apps that will be available to launch from the Start menu.
	AppUserModelIds *[]string `json:"appUserModelIds,omitempty"`

	// These are the paths of the Desktop Apps that will be available on the Start menu and the only apps the user will be
	// able to launch.
	DesktopAppPaths *[]string `json:"desktopAppPaths,omitempty"`

	// This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the
	// users to whom this kiosk configuration is assigned.
	ProfileName *string `json:"profileName,omitempty"`

	// This setting allows the admin to specify whether the Task Bar is shown or not.
	ShowTaskBar *bool `json:"showTaskBar,omitempty"`

	// Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by
	// specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
	StartMenuLayoutXml *string `json:"startMenuLayoutXml,omitempty"`

	// The user accounts that will be locked to this kiosk configuration.
	UserAccounts *[]string `json:"userAccounts,omitempty"`

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

func (s WindowsAssignedAccessProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsAssignedAccessProfile{}

func (s WindowsAssignedAccessProfile) MarshalJSON() ([]byte, error) {
	type wrapper WindowsAssignedAccessProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsAssignedAccessProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsAssignedAccessProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsAssignedAccessProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsAssignedAccessProfile: %+v", err)
	}

	return encoded, nil
}
