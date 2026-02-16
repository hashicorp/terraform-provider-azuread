package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityHostPort{}

type SecurityHostPort struct {
	// The hostPortBanners retrieved from scanning the port.
	Banners *[]SecurityHostPortBanner `json:"banners,omitempty"`

	// The first date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	Host *SecurityHost `json:"host,omitempty"`

	// The last date and time when Microsoft Defender Threat Intelligence scanned the hostPort. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	LastScanDateTime nullable.Type[string] `json:"lastScanDateTime,omitempty"`

	// The last date and time when Microsoft Defender Threat Intelligence observed the hostPort. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014, is 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The most recent sslCertificate used to communicate on the port.
	MostRecentSslCertificate *SecuritySslCertificate `json:"mostRecentSslCertificate,omitempty"`

	// The numerical identifier of the port which is standardized across the internet.
	Port *int64 `json:"port,omitempty"`

	// The general protocol used to scan the port. The possible values are: tcp, udp, unknownFutureValue.
	Protocol *SecurityHostPortProtocol `json:"protocol,omitempty"`

	// The hostPortComponents retrieved from scanning the port.
	Services *[]SecurityHostPortComponent `json:"services,omitempty"`

	// The status of the port. The possible values are: open, filtered, closed, unknownFutureValue.
	Status *SecurityHostPortStatus `json:"status,omitempty"`

	// The total amount of times that Microsoft Defender Threat Intelligence has observed the hostPort in all its scans.
	TimesObserved nullable.Type[int64] `json:"timesObserved,omitempty"`

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

func (s SecurityHostPort) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityHostPort{}

func (s SecurityHostPort) MarshalJSON() ([]byte, error) {
	type wrapper SecurityHostPort
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityHostPort: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityHostPort: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.hostPort"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityHostPort: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityHostPort{}

func (s *SecurityHostPort) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Banners                  *[]SecurityHostPortBanner    `json:"banners,omitempty"`
		FirstSeenDateTime        nullable.Type[string]        `json:"firstSeenDateTime,omitempty"`
		LastScanDateTime         nullable.Type[string]        `json:"lastScanDateTime,omitempty"`
		LastSeenDateTime         nullable.Type[string]        `json:"lastSeenDateTime,omitempty"`
		MostRecentSslCertificate *SecuritySslCertificate      `json:"mostRecentSslCertificate,omitempty"`
		Port                     *int64                       `json:"port,omitempty"`
		Protocol                 *SecurityHostPortProtocol    `json:"protocol,omitempty"`
		Services                 *[]SecurityHostPortComponent `json:"services,omitempty"`
		Status                   *SecurityHostPortStatus      `json:"status,omitempty"`
		TimesObserved            nullable.Type[int64]         `json:"timesObserved,omitempty"`
		Id                       *string                      `json:"id,omitempty"`
		ODataId                  *string                      `json:"@odata.id,omitempty"`
		ODataType                *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Banners = decoded.Banners
	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastScanDateTime = decoded.LastScanDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.MostRecentSslCertificate = decoded.MostRecentSslCertificate
	s.Port = decoded.Port
	s.Protocol = decoded.Protocol
	s.Services = decoded.Services
	s.Status = decoded.Status
	s.TimesObserved = decoded.TimesObserved
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityHostPort into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'SecurityHostPort': %+v", err)
		}
		s.Host = &impl
	}

	return nil
}
