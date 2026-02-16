package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityArtifact = SecurityPassiveDnsRecord{}

type SecurityPassiveDnsRecord struct {
	Artifact *SecurityArtifact `json:"artifact,omitempty"`

	// The date and time that this passiveDnsRecord entry was collected by Microsoft. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CollectedDateTime *string `json:"collectedDateTime,omitempty"`

	// The date and time when this passiveDnsRecord entry was first seen. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`

	// The date and time when this passiveDnsRecord entry was most recently seen. The Timestamp type represents date and
	// time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`

	ParentHost *SecurityHost `json:"parentHost,omitempty"`

	// The DNS record type for this passiveDnsRecord entry.
	RecordType *string `json:"recordType,omitempty"`

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

func (s SecurityPassiveDnsRecord) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityPassiveDnsRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityPassiveDnsRecord{}

func (s SecurityPassiveDnsRecord) MarshalJSON() ([]byte, error) {
	type wrapper SecurityPassiveDnsRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityPassiveDnsRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityPassiveDnsRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.passiveDnsRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityPassiveDnsRecord: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityPassiveDnsRecord{}

func (s *SecurityPassiveDnsRecord) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CollectedDateTime *string `json:"collectedDateTime,omitempty"`
		FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`
		LastSeenDateTime  *string `json:"lastSeenDateTime,omitempty"`
		RecordType        *string `json:"recordType,omitempty"`
		Id                *string `json:"id,omitempty"`
		ODataId           *string `json:"@odata.id,omitempty"`
		ODataType         *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CollectedDateTime = decoded.CollectedDateTime
	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.RecordType = decoded.RecordType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityPassiveDnsRecord into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["artifact"]; ok {
		impl, err := UnmarshalSecurityArtifactImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Artifact' for 'SecurityPassiveDnsRecord': %+v", err)
		}
		s.Artifact = &impl
	}

	if v, ok := temp["parentHost"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ParentHost' for 'SecurityPassiveDnsRecord': %+v", err)
		}
		s.ParentHost = &impl
	}

	return nil
}
