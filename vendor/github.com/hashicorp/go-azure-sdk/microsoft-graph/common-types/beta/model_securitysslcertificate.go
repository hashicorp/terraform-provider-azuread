package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityArtifact = SecuritySslCertificate{}

type SecuritySslCertificate struct {
	// The date and time when a certificate expires. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// A hash of the certificate calculated on the data and signature.
	Fingerprint nullable.Type[string] `json:"fingerprint,omitempty"`

	// The first date and time when this sslCertificate was observed. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// The date and time when a certificate was issued. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	IssueDateTime nullable.Type[string] `json:"issueDateTime,omitempty"`

	// The entity that grants this certificate.
	Issuer *SecuritySslCertificateEntity `json:"issuer,omitempty"`

	// The most recent date and time when this sslCertificate was observed. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The host resources related with this sslCertificate.
	RelatedHosts *[]SecurityHost `json:"relatedHosts,omitempty"`

	// The serial number associated with an SSL certificate.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// A SHA-1 hash of the certificate. Note: This is not the signature.
	Sha1 nullable.Type[string] `json:"sha1,omitempty"`

	// The person, site, machine, and so on, this certificate is for.
	Subject *SecuritySslCertificateEntity `json:"subject,omitempty"`

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

func (s SecuritySslCertificate) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecuritySslCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecuritySslCertificate{}

func (s SecuritySslCertificate) MarshalJSON() ([]byte, error) {
	type wrapper SecuritySslCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecuritySslCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySslCertificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.sslCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecuritySslCertificate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecuritySslCertificate{}

func (s *SecuritySslCertificate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExpirationDateTime nullable.Type[string]         `json:"expirationDateTime,omitempty"`
		Fingerprint        nullable.Type[string]         `json:"fingerprint,omitempty"`
		FirstSeenDateTime  nullable.Type[string]         `json:"firstSeenDateTime,omitempty"`
		IssueDateTime      nullable.Type[string]         `json:"issueDateTime,omitempty"`
		Issuer             *SecuritySslCertificateEntity `json:"issuer,omitempty"`
		LastSeenDateTime   nullable.Type[string]         `json:"lastSeenDateTime,omitempty"`
		SerialNumber       nullable.Type[string]         `json:"serialNumber,omitempty"`
		Sha1               nullable.Type[string]         `json:"sha1,omitempty"`
		Subject            *SecuritySslCertificateEntity `json:"subject,omitempty"`
		Id                 *string                       `json:"id,omitempty"`
		ODataId            *string                       `json:"@odata.id,omitempty"`
		ODataType          *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.Fingerprint = decoded.Fingerprint
	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.IssueDateTime = decoded.IssueDateTime
	s.Issuer = decoded.Issuer
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.SerialNumber = decoded.SerialNumber
	s.Sha1 = decoded.Sha1
	s.Subject = decoded.Subject
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecuritySslCertificate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["relatedHosts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RelatedHosts into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityHost, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityHostImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RelatedHosts' for 'SecuritySslCertificate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RelatedHosts = &output
	}

	return nil
}
