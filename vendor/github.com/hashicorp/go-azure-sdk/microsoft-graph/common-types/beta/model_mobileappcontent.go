package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppContent{}

type MobileAppContent struct {
	// The collection of contained apps in a MobileLobApp acting as a package.
	ContainedApps *[]MobileContainedApp `json:"containedApps,omitempty"`

	// The list of files for this app content version.
	Files *[]MobileAppContentFile `json:"files,omitempty"`

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

func (s MobileAppContent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppContent{}

func (s MobileAppContent) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppContent: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MobileAppContent{}

func (s *MobileAppContent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Files     *[]MobileAppContentFile `json:"files,omitempty"`
		Id        *string                 `json:"id,omitempty"`
		ODataId   *string                 `json:"@odata.id,omitempty"`
		ODataType *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Files = decoded.Files
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MobileAppContent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["containedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ContainedApps into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileContainedApp, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileContainedAppImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ContainedApps' for 'MobileAppContent': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ContainedApps = &output
	}

	return nil
}
