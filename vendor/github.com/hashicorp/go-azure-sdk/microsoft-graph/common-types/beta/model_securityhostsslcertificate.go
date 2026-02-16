package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityArtifact = SecurityHostSslCertificate{}

type SecurityHostSslCertificate struct {
	// The first date and time that this hostSslCertificate was observed. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// The host for this hostSslCertificate.
	Host *SecurityHost `json:"host,omitempty"`

	// The most recent date and time that this hostSslCertificate was observed. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The ports related with this hostSslCertificate.
	Ports *[]SecurityHostSslCertificatePort `json:"ports,omitempty"`

	// The sslCertificate for this hostSslCertificate.
	SslCertificate *SecuritySslCertificate `json:"sslCertificate,omitempty"`

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

func (s SecurityHostSslCertificate) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityHostSslCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHostSslCertificate{}

func (s SecurityHostSslCertificate) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHostSslCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHostSslCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHostSslCertificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.hostSslCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHostSslCertificate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityHostSslCertificate{}

func (s *SecurityHostSslCertificate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FirstSeenDateTime nullable.Type[string]             `json:"firstSeenDateTime,omitempty"`
		LastSeenDateTime  nullable.Type[string]             `json:"lastSeenDateTime,omitempty"`
		Ports             *[]SecurityHostSslCertificatePort `json:"ports,omitempty"`
		SslCertificate    *SecuritySslCertificate           `json:"sslCertificate,omitempty"`
		Id                *string                           `json:"id,omitempty"`
		ODataId           *string                           `json:"@odata.id,omitempty"`
		ODataType         *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.Ports = decoded.Ports
	s.SslCertificate = decoded.SslCertificate
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityHostSslCertificate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'SecurityHostSslCertificate': %+v", err)
		}
		s.Host = &impl
	}

	return nil
}
