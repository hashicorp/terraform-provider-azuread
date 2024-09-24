package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityHost = SecurityIPAddress{}

type SecurityIPAddress struct {
	// The details about the autonomous system to which this IP address belongs.
	AutonomousSystem *SecurityAutonomousSystem `json:"autonomousSystem,omitempty"`

	// The country/region for this IP address.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// The hosting company listed for this host.
	HostingProvider nullable.Type[string] `json:"hostingProvider,omitempty"`

	// The block of IP addresses this IP address belongs to.
	Netblock nullable.Type[string] `json:"netblock,omitempty"`

	// Fields inherited from SecurityHost

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

func (s SecurityIPAddress) SecurityHost() BaseSecurityHostImpl {
	return BaseSecurityHostImpl{
		ChildHostPairs:    s.ChildHostPairs,
		Components:        s.Components,
		Cookies:           s.Cookies,
		FirstSeenDateTime: s.FirstSeenDateTime,
		HostPairs:         s.HostPairs,
		LastSeenDateTime:  s.LastSeenDateTime,
		ParentHostPairs:   s.ParentHostPairs,
		PassiveDns:        s.PassiveDns,
		PassiveDnsReverse: s.PassiveDnsReverse,
		Ports:             s.Ports,
		Reputation:        s.Reputation,
		SslCertificates:   s.SslCertificates,
		Subdomains:        s.Subdomains,
		Trackers:          s.Trackers,
		Whois:             s.Whois,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s SecurityIPAddress) SecurityArtifact() BaseSecurityArtifactImpl {
	return BaseSecurityArtifactImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityIPAddress) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityIPAddress{}

func (s SecurityIPAddress) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIPAddress
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIPAddress: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIPAddress: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ipAddress"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIPAddress: %+v", err)
	}

	return encoded, nil
}
