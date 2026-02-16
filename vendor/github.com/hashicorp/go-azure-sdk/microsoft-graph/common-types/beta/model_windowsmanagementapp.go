package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsManagementApp{}

type WindowsManagementApp struct {
	// Windows management app available version.
	AvailableVersion nullable.Type[string] `json:"availableVersion,omitempty"`

	// The list of health states for installed Windows management app.
	HealthStates *[]WindowsManagementAppHealthState `json:"healthStates,omitempty"`

	// ManagedInstallerStatus
	ManagedInstaller *ManagedInstallerStatus `json:"managedInstaller,omitempty"`

	// Managed Installer Configured Date Time
	ManagedInstallerConfiguredDateTime nullable.Type[string] `json:"managedInstallerConfiguredDateTime,omitempty"`

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

func (s WindowsManagementApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsManagementApp{}

func (s WindowsManagementApp) MarshalJSON() ([]byte, error) {
	type wrapper WindowsManagementApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsManagementApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsManagementApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsManagementApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsManagementApp: %+v", err)
	}

	return encoded, nil
}
