package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Shared struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identity of the owner of the shared item. Read-only.
	Owner *IdentitySet `json:"owner,omitempty"`

	// Indicates the scope of how the item is shared. The possible values are: anonymous, organization, or users. Read-only.
	Scope nullable.Type[string] `json:"scope,omitempty"`

	// The identity of the user who shared the item. Read-only.
	SharedBy *IdentitySet `json:"sharedBy,omitempty"`

	// The UTC date and time when the item was shared. Read-only.
	SharedDateTime nullable.Type[string] `json:"sharedDateTime,omitempty"`
}

var _ json.Marshaler = Shared{}

func (s Shared) MarshalJSON() ([]byte, error) {
	type wrapper Shared
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Shared: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Shared: %+v", err)
	}

	delete(decoded, "owner")
	delete(decoded, "scope")
	delete(decoded, "sharedBy")
	delete(decoded, "sharedDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Shared: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Shared{}

func (s *Shared) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId        *string               `json:"@odata.id,omitempty"`
		ODataType      *string               `json:"@odata.type,omitempty"`
		Scope          nullable.Type[string] `json:"scope,omitempty"`
		SharedDateTime nullable.Type[string] `json:"sharedDateTime,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Scope = decoded.Scope
	s.SharedDateTime = decoded.SharedDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Shared into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["owner"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Owner' for 'Shared': %+v", err)
		}
		s.Owner = &impl
	}

	if v, ok := temp["sharedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SharedBy' for 'Shared': %+v", err)
		}
		s.SharedBy = &impl
	}

	return nil
}
