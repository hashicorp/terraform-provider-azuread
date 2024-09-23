package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EBookInstallSummary{}

type EBookInstallSummary struct {
	// Number of Devices that have failed to install this book.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// Number of Users that have 1 or more device that failed to install this book.
	FailedUserCount *int64 `json:"failedUserCount,omitempty"`

	// Number of Devices that have successfully installed this book.
	InstalledDeviceCount *int64 `json:"installedDeviceCount,omitempty"`

	// Number of Users whose devices have all succeeded to install this book.
	InstalledUserCount *int64 `json:"installedUserCount,omitempty"`

	// Number of Devices that does not have this book installed.
	NotInstalledDeviceCount *int64 `json:"notInstalledDeviceCount,omitempty"`

	// Number of Users that did not install this book.
	NotInstalledUserCount *int64 `json:"notInstalledUserCount,omitempty"`

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

func (s EBookInstallSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EBookInstallSummary{}

func (s EBookInstallSummary) MarshalJSON() ([]byte, error) {
	type wrapper EBookInstallSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EBookInstallSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EBookInstallSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.eBookInstallSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EBookInstallSummary: %+v", err)
	}

	return encoded, nil
}
