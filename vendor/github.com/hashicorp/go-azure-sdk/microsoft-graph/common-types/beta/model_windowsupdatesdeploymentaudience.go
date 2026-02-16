package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesDeploymentAudience{}

type WindowsUpdatesDeploymentAudience struct {
	// Content eligible to deploy to devices in the audience. Not nullable. Read-only.
	ApplicableContent *[]WindowsUpdatesApplicableContent `json:"applicableContent,omitempty"`

	// Specifies the assets to exclude from the audience.
	Exclusions *[]WindowsUpdatesUpdatableAsset `json:"exclusions,omitempty"`

	// Specifies the assets to include in the audience.
	Members *[]WindowsUpdatesUpdatableAsset `json:"members,omitempty"`

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

func (s WindowsUpdatesDeploymentAudience) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesDeploymentAudience{}

func (s WindowsUpdatesDeploymentAudience) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDeploymentAudience
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDeploymentAudience: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDeploymentAudience: %+v", err)
	}

	delete(decoded, "applicableContent")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.deploymentAudience"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDeploymentAudience: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesDeploymentAudience{}

func (s *WindowsUpdatesDeploymentAudience) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicableContent *[]WindowsUpdatesApplicableContent `json:"applicableContent,omitempty"`
		Id                *string                            `json:"id,omitempty"`
		ODataId           *string                            `json:"@odata.id,omitempty"`
		ODataType         *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicableContent = decoded.ApplicableContent
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesDeploymentAudience into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["exclusions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Exclusions into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesUpdatableAsset, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesUpdatableAssetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Exclusions' for 'WindowsUpdatesDeploymentAudience': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Exclusions = &output
	}

	if v, ok := temp["members"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Members into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdatesUpdatableAsset, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdatesUpdatableAssetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'WindowsUpdatesDeploymentAudience': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	return nil
}
