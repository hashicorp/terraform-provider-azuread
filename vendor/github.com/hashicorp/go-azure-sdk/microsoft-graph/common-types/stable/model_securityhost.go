package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHost interface {
	Entity
	SecurityArtifact
	SecurityHost() BaseSecurityHostImpl
}

var _ SecurityHost = BaseSecurityHostImpl{}

type BaseSecurityHostImpl struct {
	// The hostPairs that are resources associated with a host, where that host is the parentHost and has an outgoing
	// pairing to a childHost.
	ChildHostPairs *[]SecurityHostPair `json:"childHostPairs,omitempty"`

	// The hostComponents that are associated with this host.
	Components *[]SecurityHostComponent `json:"components,omitempty"`

	// The hostCookies that are associated with this host.
	Cookies *[]SecurityHostCookie `json:"cookies,omitempty"`

	// The first date and time when this host was observed. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	// The hostPairs that are associated with this host, where this host is either the parentHost or childHost.
	HostPairs *[]SecurityHostPair `json:"hostPairs,omitempty"`

	// The most recent date and time when this host was observed. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The hostPairs that are associated with a host, where that host is the childHost and has an incoming pairing with a
	// parentHost.
	ParentHostPairs *[]SecurityHostPair `json:"parentHostPairs,omitempty"`

	// Passive DNS retrieval about this host.
	PassiveDns *[]SecurityPassiveDnsRecord `json:"passiveDns,omitempty"`

	// Reverse passive DNS retrieval about this host.
	PassiveDnsReverse *[]SecurityPassiveDnsRecord `json:"passiveDnsReverse,omitempty"`

	// The hostPorts associated with a host.
	Ports *[]SecurityHostPort `json:"ports,omitempty"`

	// Represents a calculated reputation of this host.
	Reputation *SecurityHostReputation `json:"reputation,omitempty"`

	// The hostSslCertificates that are associated with this host.
	SslCertificates *[]SecurityHostSslCertificate `json:"sslCertificates,omitempty"`

	// The subdomains that are associated with this host.
	Subdomains *[]SecuritySubdomain `json:"subdomains,omitempty"`

	// The hostTrackers that are associated with this host.
	Trackers *[]SecurityHostTracker `json:"trackers,omitempty"`

	// The most recent whoisRecord for this host.
	Whois *SecurityWhoisRecord `json:"whois,omitempty"`

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

func (s BaseSecurityHostImpl) SecurityHost() BaseSecurityHostImpl {
	return s
}

func (s BaseSecurityHostImpl) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BaseSecurityHostImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityHost = RawSecurityHostImpl{}

// RawSecurityHostImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityHostImpl struct {
	securityHost BaseSecurityHostImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawSecurityHostImpl) SecurityHost() BaseSecurityHostImpl {
	return s.securityHost
}

func (s RawSecurityHostImpl) SecurityArtifact() BaseSecurityArtifactImpl {
	return s.securityHost.SecurityArtifact()
}

func (s RawSecurityHostImpl) Entity() BaseEntityImpl {
	return s.securityHost.Entity()
}

var _ json.Marshaler = BaseSecurityHostImpl{}

func (s BaseSecurityHostImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityHostImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityHostImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityHostImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.host"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityHostImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityHostImplementation(input []byte) (SecurityHost, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHost into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hostname") {
		var out SecurityHostname
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHostname: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ipAddress") {
		var out SecurityIPAddress
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIPAddress: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityHostImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityHostImpl: %+v", err)
	}

	return RawSecurityHostImpl{
		securityHost: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
