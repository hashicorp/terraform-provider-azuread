package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppRole struct {
	// Specifies whether this app role can be assigned to users and groups (by setting to ['User']), to other application's
	// (by setting to ['Application'], or both (by setting to ['User', 'Application']). App roles supporting assignment to
	// other applications' service principals are also known as application permissions. The 'Application' value is only
	// supported for app roles defined on application entities.
	AllowedMemberTypes *[]string `json:"allowedMemberTypes,omitempty"`

	// The description for the app role. This is displayed when the app role is being assigned and, if the app role
	// functions as an application permission, during consent experiences.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the permission that appears in the app role assignment and consent experiences.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique role identifier inside the appRoles collection. When creating a new app role, a new GUID identifier must be
	// provided.
	Id *string `json:"id,omitempty"`

	// When creating or updating an app role, this must be set to true (which is the default). To delete a role, this must
	// first be set to false. At that point, in a subsequent call, this role may be removed.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies if the app role is defined on the application object or on the servicePrincipal entity. Must not be
	// included in any POST or PATCH requests. Read-only.
	Origin nullable.Type[string] `json:"origin,omitempty"`

	// Specifies the value to include in the roles claim in ID tokens and access tokens authenticating an assigned user or
	// service principal. Must not exceed 120 characters in length. Allowed characters are : ! # $ % & ' ( ) * + , - . / : ;
	// = ? @ [ ] ^ + _ { } ~, and characters in the ranges 0-9, A-Z and a-z. Any other character, including the space
	// character, aren't allowed. May not begin with ..
	Value nullable.Type[string] `json:"value,omitempty"`
}

var _ json.Marshaler = AppRole{}

func (s AppRole) MarshalJSON() ([]byte, error) {
	type wrapper AppRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppRole: %+v", err)
	}

	delete(decoded, "origin")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppRole: %+v", err)
	}

	return encoded, nil
}
