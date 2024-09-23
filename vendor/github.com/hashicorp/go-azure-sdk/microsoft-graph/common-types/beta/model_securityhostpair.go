package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityHostPair{}

type SecurityHostPair struct {
	ChildHost *SecurityHost `json:"childHost,omitempty"`

	// The first date and time when Microsoft Defender Threat Intelligence observed the hostPair. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// The last date and time when Microsoft Defender Threat Intelligence observed the hostPair. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The reason that two hosts are identified as hostPair.
	LinkKind nullable.Type[string] `json:"linkKind,omitempty"`

	ParentHost *SecurityHost `json:"parentHost,omitempty"`

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

func (s SecurityHostPair) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHostPair{}

func (s SecurityHostPair) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHostPair
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHostPair: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHostPair: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.hostPair"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHostPair: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityHostPair{}

func (s *SecurityHostPair) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`
		LastSeenDateTime  nullable.Type[string] `json:"lastSeenDateTime,omitempty"`
		LinkKind          nullable.Type[string] `json:"linkKind,omitempty"`
		Id                *string               `json:"id,omitempty"`
		ODataId           *string               `json:"@odata.id,omitempty"`
		ODataType         *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.LinkKind = decoded.LinkKind
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityHostPair into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["childHost"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ChildHost' for 'SecurityHostPair': %+v", err)
		}
		s.ChildHost = &impl
	}

	if v, ok := temp["parentHost"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ParentHost' for 'SecurityHostPair': %+v", err)
		}
		s.ParentHost = &impl
	}

	return nil
}
