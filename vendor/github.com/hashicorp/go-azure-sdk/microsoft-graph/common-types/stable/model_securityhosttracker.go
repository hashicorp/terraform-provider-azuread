package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityArtifact = SecurityHostTracker{}

type SecurityHostTracker struct {
	// The first date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on
	// Jan 1, 2014, is 2014-01-01T00:00:00Z.
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`

	Host *SecurityHost `json:"host,omitempty"`

	// The kind of hostTracker that was detected. For example, GoogleAnalyticsID or JarmHash.
	Kind *string `json:"kind,omitempty"`

	// The most recent date and time when this hostTracker was observed by Microsoft Defender Threat Intelligence. The
	// timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight
	// UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`

	// The identification value for the hostTracker.
	Value *string `json:"value,omitempty"`

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

func (s SecurityHostTracker) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityHostTracker) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHostTracker{}

func (s SecurityHostTracker) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHostTracker
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHostTracker: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHostTracker: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.hostTracker"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHostTracker: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityHostTracker{}

func (s *SecurityHostTracker) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`
		Kind              *string `json:"kind,omitempty"`
		LastSeenDateTime  *string `json:"lastSeenDateTime,omitempty"`
		Value             *string `json:"value,omitempty"`
		Id                *string `json:"id,omitempty"`
		ODataId           *string `json:"@odata.id,omitempty"`
		ODataType         *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.Kind = decoded.Kind
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.Value = decoded.Value
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityHostTracker into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'SecurityHostTracker': %+v", err)
		}
		s.Host = &impl
	}

	return nil
}
