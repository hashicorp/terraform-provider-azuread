package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataPasswordSettings = IndustryDataSimplePasswordSettings{}

type IndustryDataSimplePasswordSettings struct {
	// The password for the user.
	Password *string `json:"password,omitempty"`

	// Fields inherited from IndustryDataPasswordSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IndustryDataSimplePasswordSettings) IndustryDataPasswordSettings() BaseIndustryDataPasswordSettingsImpl {
	return BaseIndustryDataPasswordSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataSimplePasswordSettings{}

func (s IndustryDataSimplePasswordSettings) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataSimplePasswordSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataSimplePasswordSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataSimplePasswordSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.simplePasswordSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataSimplePasswordSettings: %+v", err)
	}

	return encoded, nil
}
