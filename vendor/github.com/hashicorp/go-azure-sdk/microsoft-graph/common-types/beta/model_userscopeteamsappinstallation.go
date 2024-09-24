package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamsAppInstallation = UserScopeTeamsAppInstallation{}

type UserScopeTeamsAppInstallation struct {
	// The chat between the user and Teams app.
	Chat *Chat `json:"chat,omitempty"`

	// Fields inherited from TeamsAppInstallation

	// The set of resource-specific permissions consented to while installing or upgrading the teamsApp.
	ConsentedPermissionSet *TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`

	ScopeInfo TeamsAppInstallationScopeInfo `json:"scopeInfo"`

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

func (s UserScopeTeamsAppInstallation) TeamsAppInstallation() BaseTeamsAppInstallationImpl {
	return BaseTeamsAppInstallationImpl{
		ConsentedPermissionSet: s.ConsentedPermissionSet,
		ScopeInfo:              s.ScopeInfo,
		TeamsApp:               s.TeamsApp,
		TeamsAppDefinition:     s.TeamsAppDefinition,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s UserScopeTeamsAppInstallation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserScopeTeamsAppInstallation{}

func (s UserScopeTeamsAppInstallation) MarshalJSON() ([]byte, error) {
	type wrapper UserScopeTeamsAppInstallation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserScopeTeamsAppInstallation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserScopeTeamsAppInstallation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userScopeTeamsAppInstallation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserScopeTeamsAppInstallation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UserScopeTeamsAppInstallation{}

func (s *UserScopeTeamsAppInstallation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Chat                   *Chat                  `json:"chat,omitempty"`
		ConsentedPermissionSet *TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`
		TeamsApp               *TeamsApp              `json:"teamsApp,omitempty"`
		TeamsAppDefinition     *TeamsAppDefinition    `json:"teamsAppDefinition,omitempty"`
		Id                     *string                `json:"id,omitempty"`
		ODataId                *string                `json:"@odata.id,omitempty"`
		ODataType              *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Chat = decoded.Chat
	s.ConsentedPermissionSet = decoded.ConsentedPermissionSet
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.TeamsApp = decoded.TeamsApp
	s.TeamsAppDefinition = decoded.TeamsAppDefinition

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UserScopeTeamsAppInstallation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["scopeInfo"]; ok {
		impl, err := UnmarshalTeamsAppInstallationScopeInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ScopeInfo' for 'UserScopeTeamsAppInstallation': %+v", err)
		}
		s.ScopeInfo = impl
	}

	return nil
}
