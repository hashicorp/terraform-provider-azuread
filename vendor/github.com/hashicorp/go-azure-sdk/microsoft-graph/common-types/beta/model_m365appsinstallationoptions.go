package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = M365AppsInstallationOptions{}

type M365AppsInstallationOptions struct {
	AppsForMac     *AppsInstallationOptionsForMac     `json:"appsForMac,omitempty"`
	AppsForWindows *AppsInstallationOptionsForWindows `json:"appsForWindows,omitempty"`
	UpdateChannel  *AppsUpdateChannelType             `json:"updateChannel,omitempty"`

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

func (s M365AppsInstallationOptions) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = M365AppsInstallationOptions{}

func (s M365AppsInstallationOptions) MarshalJSON() ([]byte, error) {
	type wrapper M365AppsInstallationOptions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling M365AppsInstallationOptions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling M365AppsInstallationOptions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.m365AppsInstallationOptions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling M365AppsInstallationOptions: %+v", err)
	}

	return encoded, nil
}
