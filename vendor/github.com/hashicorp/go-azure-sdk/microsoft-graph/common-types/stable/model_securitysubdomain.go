package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecuritySubdomain{}

type SecuritySubdomain struct {
	// The date and time when Microsoft Defender Threat Intelligence first observed the subdomain. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014 is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	Host *SecurityHost `json:"host,omitempty"`

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

func (s SecuritySubdomain) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecuritySubdomain{}

func (s SecuritySubdomain) MarshalJSON() ([]byte, error) {
	type wrapper SecuritySubdomain
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecuritySubdomain: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySubdomain: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.subdomain"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecuritySubdomain: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecuritySubdomain{}

func (s *SecuritySubdomain) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`
		Id                *string               `json:"id,omitempty"`
		ODataId           *string               `json:"@odata.id,omitempty"`
		ODataType         *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecuritySubdomain into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'SecuritySubdomain': %+v", err)
		}
		s.Host = &impl
	}

	return nil
}
