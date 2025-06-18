package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = IosDdmLobAppAssignmentSettings{}

type IosDdmLobAppAssignmentSettings struct {

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosDdmLobAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosDdmLobAppAssignmentSettings{}

func (s IosDdmLobAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper IosDdmLobAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosDdmLobAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosDdmLobAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosDdmLobAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosDdmLobAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
