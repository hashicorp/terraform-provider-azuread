package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppInstallation interface {
	Entity
	TeamsAppInstallation() BaseTeamsAppInstallationImpl
}

var _ TeamsAppInstallation = BaseTeamsAppInstallationImpl{}

type BaseTeamsAppInstallationImpl struct {
	// The set of resource-specific permissions consented to while installing or upgrading the teamsApp.
	ConsentedPermissionSet *TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`

	// The app that is installed.
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`

	// The details of this version of the app.
	TeamsAppDefinition *TeamsAppDefinition `json:"teamsAppDefinition,omitempty"`

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

func (s BaseTeamsAppInstallationImpl) TeamsAppInstallation() BaseTeamsAppInstallationImpl {
	return s
}

func (s BaseTeamsAppInstallationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ TeamsAppInstallation = RawTeamsAppInstallationImpl{}

// RawTeamsAppInstallationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTeamsAppInstallationImpl struct {
	teamsAppInstallation BaseTeamsAppInstallationImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawTeamsAppInstallationImpl) TeamsAppInstallation() BaseTeamsAppInstallationImpl {
	return s.teamsAppInstallation
}

func (s RawTeamsAppInstallationImpl) Entity() BaseEntityImpl {
	return s.teamsAppInstallation.Entity()
}

var _ json.Marshaler = BaseTeamsAppInstallationImpl{}

func (s BaseTeamsAppInstallationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseTeamsAppInstallationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseTeamsAppInstallationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseTeamsAppInstallationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAppInstallation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseTeamsAppInstallationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalTeamsAppInstallationImplementation(input []byte) (TeamsAppInstallation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppInstallation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.userScopeTeamsAppInstallation") {
		var out UserScopeTeamsAppInstallation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserScopeTeamsAppInstallation: %+v", err)
		}
		return out, nil
	}

	var parent BaseTeamsAppInstallationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTeamsAppInstallationImpl: %+v", err)
	}

	return RawTeamsAppInstallationImpl{
		teamsAppInstallation: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
