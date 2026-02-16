package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingInvitation struct {
	// The email address provided for the recipient of the sharing invitation. Read-only.
	Email nullable.Type[string] `json:"email,omitempty"`

	// Provides information about who sent the invitation that created this permission, if that information is available.
	// Read-only.
	InvitedBy *IdentitySet `json:"invitedBy,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RedeemedBy nullable.Type[string] `json:"redeemedBy,omitempty"`

	// If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
	SignInRequired nullable.Type[bool] `json:"signInRequired,omitempty"`
}

var _ json.Marshaler = SharingInvitation{}

func (s SharingInvitation) MarshalJSON() ([]byte, error) {
	type wrapper SharingInvitation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharingInvitation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharingInvitation: %+v", err)
	}

	delete(decoded, "email")
	delete(decoded, "invitedBy")
	delete(decoded, "signInRequired")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharingInvitation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SharingInvitation{}

func (s *SharingInvitation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Email          nullable.Type[string] `json:"email,omitempty"`
		ODataId        *string               `json:"@odata.id,omitempty"`
		ODataType      *string               `json:"@odata.type,omitempty"`
		RedeemedBy     nullable.Type[string] `json:"redeemedBy,omitempty"`
		SignInRequired nullable.Type[bool]   `json:"signInRequired,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Email = decoded.Email
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RedeemedBy = decoded.RedeemedBy
	s.SignInRequired = decoded.SignInRequired

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharingInvitation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["invitedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InvitedBy' for 'SharingInvitation': %+v", err)
		}
		s.InvitedBy = &impl
	}

	return nil
}
