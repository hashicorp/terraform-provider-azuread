package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAppSettings{}

type TeamsAppSettings struct {
	// Indicates whether users are allowed to request access to the unavailable Teams apps.
	AllowUserRequestsForAppAccess nullable.Type[bool] `json:"allowUserRequestsForAppAccess,omitempty"`

	CustomAppSettings *CustomAppSettings `json:"customAppSettings,omitempty"`

	// Indicates whether resource-specific consent for chats/meetings has been enabled for the tenant. True indicates that
	// Teams apps that are allowed in the tenant and require resource-specific permissions can be installed inside chats and
	// meetings. False blocks the installation of any Teams app that requires resource-specific permissions in a chat or a
	// meeting.
	IsChatResourceSpecificConsentEnabled nullable.Type[bool] `json:"isChatResourceSpecificConsentEnabled,omitempty"`

	// Indicates whether resource-specific consent for personal scope in Teams apps has been enabled for the tenant. True
	// indicates that Teams apps that are allowed in the tenant and require resource-specific permissions can be installed
	// in the personal scope. False blocks the installation of any Teams app that requires resource-specific permissions in
	// the personal scope.
	IsUserPersonalScopeResourceSpecificConsentEnabled nullable.Type[bool] `json:"isUserPersonalScopeResourceSpecificConsentEnabled,omitempty"`

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

func (s TeamsAppSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAppSettings{}

func (s TeamsAppSettings) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAppSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAppSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAppSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAppSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAppSettings: %+v", err)
	}

	return encoded, nil
}
