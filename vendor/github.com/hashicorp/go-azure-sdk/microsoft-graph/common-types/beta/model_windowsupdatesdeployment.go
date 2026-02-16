package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesDeployment{}

type WindowsUpdatesDeployment struct {
	// Specifies the audience to which content is deployed.
	Audience *WindowsUpdatesDeploymentAudience `json:"audience,omitempty"`

	// Specifies what content to deploy. Cannot be changed. Returned by default.
	Content WindowsUpdatesDeployableContent `json:"content"`

	// The date and time the deployment was created. Returned by default. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time the deployment was last modified. Returned by default. Read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Settings specified on the specific deployment governing how to deploy content. Returned by default.
	Settings *WindowsUpdatesDeploymentSettings `json:"settings,omitempty"`

	// Execution status of the deployment. Returned by default.
	State *WindowsUpdatesDeploymentState `json:"state,omitempty"`

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

func (s WindowsUpdatesDeployment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesDeployment{}

func (s WindowsUpdatesDeployment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesDeployment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesDeployment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesDeployment: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.deployment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesDeployment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesDeployment{}

func (s *WindowsUpdatesDeployment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Audience             *WindowsUpdatesDeploymentAudience `json:"audience,omitempty"`
		CreatedDateTime      *string                           `json:"createdDateTime,omitempty"`
		LastModifiedDateTime *string                           `json:"lastModifiedDateTime,omitempty"`
		Settings             *WindowsUpdatesDeploymentSettings `json:"settings,omitempty"`
		State                *WindowsUpdatesDeploymentState    `json:"state,omitempty"`
		Id                   *string                           `json:"id,omitempty"`
		ODataId              *string                           `json:"@odata.id,omitempty"`
		ODataType            *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Audience = decoded.Audience
	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Settings = decoded.Settings
	s.State = decoded.State
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesDeployment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["content"]; ok {
		impl, err := UnmarshalWindowsUpdatesDeployableContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Content' for 'WindowsUpdatesDeployment': %+v", err)
		}
		s.Content = impl
	}

	return nil
}
