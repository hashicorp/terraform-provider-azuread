package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWhoisBaseRecord interface {
	Entity
	SecurityWhoisBaseRecord() BaseSecurityWhoisBaseRecordImpl
}

var _ SecurityWhoisBaseRecord = BaseSecurityWhoisBaseRecordImpl{}

type BaseSecurityWhoisBaseRecordImpl struct {
	// The contact information for the abuse contact.
	Abuse *SecurityWhoisContact `json:"abuse,omitempty"`

	// The contact information for the admin contact.
	Admin *SecurityWhoisContact `json:"admin,omitempty"`

	// The contact information for the billing contact.
	Billing *SecurityWhoisContact `json:"billing,omitempty"`

	// The domain status for this WHOIS object.
	DomainStatus nullable.Type[string] `json:"domainStatus,omitempty"`

	// The date and time when this WHOIS record expires with the registrar. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The first seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	FirstSeenDateTime nullable.Type[string] `json:"firstSeenDateTime,omitempty"`

	Host *SecurityHost `json:"host,omitempty"`

	// The last seen date and time of this WHOIS record. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The date and time when this WHOIS record was last modified. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastUpdateDateTime nullable.Type[string] `json:"lastUpdateDateTime,omitempty"`

	// The nameservers for this WHOIS object.
	Nameservers *[]SecurityWhoisNameserver `json:"nameservers,omitempty"`

	// The contact information for the noc contact.
	Noc *SecurityWhoisContact `json:"noc,omitempty"`

	// The raw WHOIS details for this WHOIS object.
	RawWhoisText nullable.Type[string] `json:"rawWhoisText,omitempty"`

	// The contact information for the registrant contact.
	Registrant *SecurityWhoisContact `json:"registrant,omitempty"`

	// The contact information for the registrar contact.
	Registrar *SecurityWhoisContact `json:"registrar,omitempty"`

	// The date and time when this WHOIS record was registered with a registrar. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	RegistrationDateTime nullable.Type[string] `json:"registrationDateTime,omitempty"`

	// The contact information for the technical contact.
	Technical *SecurityWhoisContact `json:"technical,omitempty"`

	// The WHOIS server that provides the details.
	WhoisServer nullable.Type[string] `json:"whoisServer,omitempty"`

	// The contact information for the zone contact.
	Zone *SecurityWhoisContact `json:"zone,omitempty"`

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

func (s BaseSecurityWhoisBaseRecordImpl) SecurityWhoisBaseRecord() BaseSecurityWhoisBaseRecordImpl {
	return s
}

func (s BaseSecurityWhoisBaseRecordImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityWhoisBaseRecord = RawSecurityWhoisBaseRecordImpl{}

// RawSecurityWhoisBaseRecordImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityWhoisBaseRecordImpl struct {
	securityWhoisBaseRecord BaseSecurityWhoisBaseRecordImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawSecurityWhoisBaseRecordImpl) SecurityWhoisBaseRecord() BaseSecurityWhoisBaseRecordImpl {
	return s.securityWhoisBaseRecord
}

func (s RawSecurityWhoisBaseRecordImpl) Entity() BaseEntityImpl {
	return s.securityWhoisBaseRecord.Entity()
}

var _ json.Marshaler = BaseSecurityWhoisBaseRecordImpl{}

func (s BaseSecurityWhoisBaseRecordImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityWhoisBaseRecordImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityWhoisBaseRecordImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityWhoisBaseRecordImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.whoisBaseRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityWhoisBaseRecordImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityWhoisBaseRecordImpl{}

func (s *BaseSecurityWhoisBaseRecordImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Abuse                *SecurityWhoisContact      `json:"abuse,omitempty"`
		Admin                *SecurityWhoisContact      `json:"admin,omitempty"`
		Billing              *SecurityWhoisContact      `json:"billing,omitempty"`
		DomainStatus         nullable.Type[string]      `json:"domainStatus,omitempty"`
		ExpirationDateTime   nullable.Type[string]      `json:"expirationDateTime,omitempty"`
		FirstSeenDateTime    nullable.Type[string]      `json:"firstSeenDateTime,omitempty"`
		LastSeenDateTime     nullable.Type[string]      `json:"lastSeenDateTime,omitempty"`
		LastUpdateDateTime   nullable.Type[string]      `json:"lastUpdateDateTime,omitempty"`
		Nameservers          *[]SecurityWhoisNameserver `json:"nameservers,omitempty"`
		Noc                  *SecurityWhoisContact      `json:"noc,omitempty"`
		RawWhoisText         nullable.Type[string]      `json:"rawWhoisText,omitempty"`
		Registrant           *SecurityWhoisContact      `json:"registrant,omitempty"`
		Registrar            *SecurityWhoisContact      `json:"registrar,omitempty"`
		RegistrationDateTime nullable.Type[string]      `json:"registrationDateTime,omitempty"`
		Technical            *SecurityWhoisContact      `json:"technical,omitempty"`
		WhoisServer          nullable.Type[string]      `json:"whoisServer,omitempty"`
		Zone                 *SecurityWhoisContact      `json:"zone,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Abuse = decoded.Abuse
	s.Admin = decoded.Admin
	s.Billing = decoded.Billing
	s.DomainStatus = decoded.DomainStatus
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.LastUpdateDateTime = decoded.LastUpdateDateTime
	s.Nameservers = decoded.Nameservers
	s.Noc = decoded.Noc
	s.RawWhoisText = decoded.RawWhoisText
	s.Registrant = decoded.Registrant
	s.Registrar = decoded.Registrar
	s.RegistrationDateTime = decoded.RegistrationDateTime
	s.Technical = decoded.Technical
	s.WhoisServer = decoded.WhoisServer
	s.Zone = decoded.Zone
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityWhoisBaseRecordImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["host"]; ok {
		impl, err := UnmarshalSecurityHostImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Host' for 'BaseSecurityWhoisBaseRecordImpl': %+v", err)
		}
		s.Host = &impl
	}

	return nil
}

func UnmarshalSecurityWhoisBaseRecordImplementation(input []byte) (SecurityWhoisBaseRecord, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityWhoisBaseRecord into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.whoisHistoryRecord") {
		var out SecurityWhoisHistoryRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWhoisHistoryRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.whoisRecord") {
		var out SecurityWhoisRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityWhoisRecord: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityWhoisBaseRecordImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityWhoisBaseRecordImpl: %+v", err)
	}

	return RawSecurityWhoisBaseRecordImpl{
		securityWhoisBaseRecord: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
