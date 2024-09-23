package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = AndroidManagedStoreAppAssignmentSettings{}

type AndroidManagedStoreAppAssignmentSettings struct {
	// The track IDs to enable for this app assignment.
	AndroidManagedStoreAppTrackIds *[]string `json:"androidManagedStoreAppTrackIds,omitempty"`

	// Prioritization for automatic updates of Android Managed Store apps set on assignment.
	AutoUpdateMode *AndroidManagedStoreAutoUpdateMode `json:"autoUpdateMode,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidManagedStoreAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidManagedStoreAppAssignmentSettings{}

func (s AndroidManagedStoreAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper AndroidManagedStoreAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidManagedStoreAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidManagedStoreAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidManagedStoreAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidManagedStoreAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
