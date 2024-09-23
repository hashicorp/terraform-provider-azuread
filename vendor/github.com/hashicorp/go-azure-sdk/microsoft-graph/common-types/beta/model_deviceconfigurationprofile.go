package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceConfigurationProfile{}

type DeviceConfigurationProfile struct {
	// Account Id.
	AccountId nullable.Type[string] `json:"accountId,omitempty"`

	// Configuration Technologies for Settins Catalog Policies
	ConfigurationTechnologies *int64 `json:"configurationTechnologies,omitempty"`

	// The date and time the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Platform Type
	PlatformType *PlatformType `json:"platformType,omitempty"`

	// Profile name
	ProfileName nullable.Type[string] `json:"profileName,omitempty"`

	// Profile Type
	ProfileType *ProfileType `json:"profileType,omitempty"`

	// The list of scope tags for the configuration.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// TemplateId for Settings Catalog Policies
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

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

func (s DeviceConfigurationProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceConfigurationProfile{}

func (s DeviceConfigurationProfile) MarshalJSON() ([]byte, error) {
	type wrapper DeviceConfigurationProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceConfigurationProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfigurationProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfigurationProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceConfigurationProfile: %+v", err)
	}

	return encoded, nil
}
