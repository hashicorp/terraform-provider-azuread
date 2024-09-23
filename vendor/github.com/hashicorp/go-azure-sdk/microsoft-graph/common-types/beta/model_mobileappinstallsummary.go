package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppInstallSummary{}

type MobileAppInstallSummary struct {
	// Number of Devices that have failed to install this app.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// Number of Users that have 1 or more device that failed to install this app.
	FailedUserCount *int64 `json:"failedUserCount,omitempty"`

	// Number of Devices that have successfully installed this app.
	InstalledDeviceCount *int64 `json:"installedDeviceCount,omitempty"`

	// Number of Users whose devices have all succeeded to install this app.
	InstalledUserCount *int64 `json:"installedUserCount,omitempty"`

	// Number of Devices that are not applicable for this app.
	NotApplicableDeviceCount *int64 `json:"notApplicableDeviceCount,omitempty"`

	// Number of Users whose devices were all not applicable for this app.
	NotApplicableUserCount *int64 `json:"notApplicableUserCount,omitempty"`

	// Number of Devices that does not have this app installed.
	NotInstalledDeviceCount *int64 `json:"notInstalledDeviceCount,omitempty"`

	// Number of Users that have 1 or more devices that did not install this app.
	NotInstalledUserCount *int64 `json:"notInstalledUserCount,omitempty"`

	// Number of Devices that have been notified to install this app.
	PendingInstallDeviceCount *int64 `json:"pendingInstallDeviceCount,omitempty"`

	// Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with
	// failures.
	PendingInstallUserCount *int64 `json:"pendingInstallUserCount,omitempty"`

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

func (s MobileAppInstallSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppInstallSummary{}

func (s MobileAppInstallSummary) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppInstallSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppInstallSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppInstallSummary: %+v", err)
	}

	delete(decoded, "failedDeviceCount")
	delete(decoded, "failedUserCount")
	delete(decoded, "installedDeviceCount")
	delete(decoded, "installedUserCount")
	delete(decoded, "notApplicableDeviceCount")
	delete(decoded, "notApplicableUserCount")
	delete(decoded, "notInstalledDeviceCount")
	delete(decoded, "notInstalledUserCount")
	delete(decoded, "pendingInstallDeviceCount")
	delete(decoded, "pendingInstallUserCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppInstallSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppInstallSummary: %+v", err)
	}

	return encoded, nil
}
