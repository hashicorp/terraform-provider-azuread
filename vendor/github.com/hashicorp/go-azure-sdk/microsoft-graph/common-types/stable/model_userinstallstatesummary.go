package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserInstallStateSummary{}

type UserInstallStateSummary struct {
	// The install state of the eBook.
	DeviceStates *[]DeviceInstallState `json:"deviceStates,omitempty"`

	// Failed Device Count.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// Installed Device Count.
	InstalledDeviceCount *int64 `json:"installedDeviceCount,omitempty"`

	// Not installed device count.
	NotInstalledDeviceCount *int64 `json:"notInstalledDeviceCount,omitempty"`

	// User name.
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s UserInstallStateSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserInstallStateSummary{}

func (s UserInstallStateSummary) MarshalJSON() ([]byte, error) {
	type wrapper UserInstallStateSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserInstallStateSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserInstallStateSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userInstallStateSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserInstallStateSummary: %+v", err)
	}

	return encoded, nil
}
