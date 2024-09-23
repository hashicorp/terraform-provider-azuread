package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdateState{}

type WindowsUpdateState struct {
	// Device display name.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The id of the device.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The current feature update version of the device.
	FeatureUpdateVersion nullable.Type[string] `json:"featureUpdateVersion,omitempty"`

	// The date time that the Windows Update Agent did a successful scan.
	LastScanDateTime *string `json:"lastScanDateTime,omitempty"`

	// Last date time that the device sync with with Microsoft Intune.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// The Quality Update Version of the device.
	QualityUpdateVersion nullable.Type[string] `json:"qualityUpdateVersion,omitempty"`

	// Windows update for business configuration device states
	Status *WindowsUpdateStatus `json:"status,omitempty"`

	// The id of the user.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// User principal name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s WindowsUpdateState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdateState{}

func (s WindowsUpdateState) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdateState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdateState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdateState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdateState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdateState: %+v", err)
	}

	return encoded, nil
}
