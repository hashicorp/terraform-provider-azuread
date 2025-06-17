package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OAuth2PermissionGrant{}

type OAuth2PermissionGrant struct {
	// The object id (not appId) of the client service principal for the application that is authorized to act on behalf of
	// a signed-in user when accessing an API. Required. Supports $filter (eq only).
	ClientId *string `json:"clientId,omitempty"`

	// Indicates whether authorization is granted for the client application to impersonate all users or only a specific
	// user. AllPrincipals indicates authorization to impersonate all users. Principal indicates authorization to
	// impersonate a specific user. Consent on behalf of all users can be granted by an administrator. Nonadmin users may be
	// authorized to consent on behalf of themselves in some cases, for some delegated permissions. Required. Supports
	// $filter (eq only).
	ConsentType nullable.Type[string] `json:"consentType,omitempty"`

	// Currently, the end time value is ignored, but a value is required when creating an oAuth2PermissionGrant. Required.
	ExpiryTime nullable.Type[string] `json:"expiryTime,omitempty"`

	// The id of the user on behalf of whom the client is authorized to access the resource, when consentType is Principal.
	// If consentType is AllPrincipals this value is null. Required when consentType is Principal. Supports $filter (eq
	// only).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// The id of the resource service principal to which access is authorized. This identifies the API that the client is
	// authorized to attempt to call on behalf of a signed-in user. Supports $filter (eq only).
	ResourceId *string `json:"resourceId,omitempty"`

	// A space-separated list of the claim values for delegated permissions that should be included in access tokens for the
	// resource application (the API). For example, openid User.Read GroupMember.Read.All. Each claim value should match the
	// value field of one of the delegated permissions defined by the API, listed in the publishedPermissionScopes property
	// of the resource service principal. Must not exceed 3850 characters in length.
	Scope nullable.Type[string] `json:"scope,omitempty"`

	// Currently, the start time value is ignored, but a value is required when creating an oAuth2PermissionGrant. Required.
	StartTime nullable.Type[string] `json:"startTime,omitempty"`

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

func (s OAuth2PermissionGrant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OAuth2PermissionGrant{}

func (s OAuth2PermissionGrant) MarshalJSON() ([]byte, error) {
	type wrapper OAuth2PermissionGrant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OAuth2PermissionGrant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OAuth2PermissionGrant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.oAuth2PermissionGrant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OAuth2PermissionGrant: %+v", err)
	}

	return encoded, nil
}
