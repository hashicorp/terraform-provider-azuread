package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Permission{}

type Permission struct {
	// A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue
	// indicates there's no expiration set for this permission. Optional.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// For user type permissions, the details of the users and applications for this permission. Read-only.
	GrantedTo *IdentitySet `json:"grantedTo,omitempty"`

	// For type permissions, the details of the users to whom permission was granted. Read-only.
	GrantedToIdentities *[]IdentitySet `json:"grantedToIdentities,omitempty"`

	// For link type permissions, the details of the users to whom permission was granted. Read-only.
	GrantedToIdentitiesV2 *[]SharePointIdentitySet `json:"grantedToIdentitiesV2,omitempty"`

	// For user type permissions, the details of the users and applications for this permission. Read-only.
	GrantedToV2 *SharePointIdentitySet `json:"grantedToV2,omitempty"`

	// Indicates whether the password is set for this permission. This property only appears in the response. Optional.
	// Read-only. For OneDrive Personal only.
	HasPassword nullable.Type[bool] `json:"hasPassword,omitempty"`

	// Provides a reference to the ancestor of the current permission, if inherited from an ancestor. Read-only.
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`

	// Details of any associated sharing invitation for this permission. Read-only.
	Invitation *SharingInvitation `json:"invitation,omitempty"`

	// Provides the link details of the current permission, if it's a link type permission. Read-only.
	Link *SharingLink `json:"link,omitempty"`

	// The type of permission, for example, read. See the Roles property values section for the full list of roles.
	// Read-only.
	Roles *[]string `json:"roles,omitempty"`

	// A unique token that can be used to access this shared item via the shares API. Read-only.
	ShareId nullable.Type[string] `json:"shareId,omitempty"`

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

func (s Permission) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Permission{}

func (s Permission) MarshalJSON() ([]byte, error) {
	type wrapper Permission
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Permission: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Permission: %+v", err)
	}

	delete(decoded, "grantedTo")
	delete(decoded, "grantedToIdentities")
	delete(decoded, "grantedToIdentitiesV2")
	delete(decoded, "grantedToV2")
	delete(decoded, "hasPassword")
	delete(decoded, "inheritedFrom")
	delete(decoded, "invitation")
	delete(decoded, "link")
	delete(decoded, "roles")
	delete(decoded, "shareId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.permission"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Permission: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Permission{}

func (s *Permission) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExpirationDateTime    nullable.Type[string]    `json:"expirationDateTime,omitempty"`
		GrantedToIdentitiesV2 *[]SharePointIdentitySet `json:"grantedToIdentitiesV2,omitempty"`
		GrantedToV2           *SharePointIdentitySet   `json:"grantedToV2,omitempty"`
		HasPassword           nullable.Type[bool]      `json:"hasPassword,omitempty"`
		InheritedFrom         *ItemReference           `json:"inheritedFrom,omitempty"`
		Invitation            *SharingInvitation       `json:"invitation,omitempty"`
		Link                  *SharingLink             `json:"link,omitempty"`
		Roles                 *[]string                `json:"roles,omitempty"`
		ShareId               nullable.Type[string]    `json:"shareId,omitempty"`
		Id                    *string                  `json:"id,omitempty"`
		ODataId               *string                  `json:"@odata.id,omitempty"`
		ODataType             *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.GrantedToIdentitiesV2 = decoded.GrantedToIdentitiesV2
	s.GrantedToV2 = decoded.GrantedToV2
	s.HasPassword = decoded.HasPassword
	s.InheritedFrom = decoded.InheritedFrom
	s.Invitation = decoded.Invitation
	s.Link = decoded.Link
	s.Roles = decoded.Roles
	s.ShareId = decoded.ShareId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Permission into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["grantedTo"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GrantedTo' for 'Permission': %+v", err)
		}
		s.GrantedTo = &impl
	}

	if v, ok := temp["grantedToIdentities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling GrantedToIdentities into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentitySet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentitySetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'GrantedToIdentities' for 'Permission': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.GrantedToIdentities = &output
	}

	return nil
}
