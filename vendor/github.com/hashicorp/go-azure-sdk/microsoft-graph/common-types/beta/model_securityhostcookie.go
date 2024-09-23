package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityArtifact = SecurityHostCookie{}

type SecurityHostCookie struct {
	// The URI for which the cookie is valid.
	Domain *string `json:"domain,omitempty"`

	// The first date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The Timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC
	// on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`

	Host *SecurityHost `json:"host,omitempty"`

	// The most recent date and time when this hostCookie was observed by Microsoft Defender Threat Intelligence. The
	// Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`

	// The name of the cookie, for example, JSESSIONID or SEARCH_NAMESITE.
	Name *string `json:"name,omitempty"`

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

func (s SecurityHostCookie) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityHostCookie) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHostCookie{}

func (s SecurityHostCookie) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHostCookie
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHostCookie: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHostCookie: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.hostCookie"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHostCookie: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityHostCookie{}

func (s *SecurityHostCookie) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Domain            *string `json:"domain,omitempty"`
		FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`
		LastSeenDateTime  *string `json:"lastSeenDateTime,omitempty"`
		Name              *string `json:"name,omitempty"`
		Id                *string `json:"id,omitempty"`
		ODataId           *string `json:"@odata.id,omitempty"`
		ODataType         *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Domain = decoded.Domain
	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.Name = decoded.Name
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityHostCookie into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'SecurityHostCookie': %+v", err)
		}
		s.Host = &impl
	}

	return nil
}
