package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainDnsRecord interface {
	Entity
	DomainDnsRecord() BaseDomainDnsRecordImpl
}

var _ DomainDnsRecord = BaseDomainDnsRecordImpl{}

type BaseDomainDnsRecordImpl struct {
	// If false, the customer must configure this record at the DNS host for Microsoft Online Services to operate correctly
	// with the domain.
	IsOptional *bool `json:"isOptional,omitempty"`

	// Value used when configuring the name of the DNS record at the DNS host.
	Label *string `json:"label,omitempty"`

	// Indicates what type of DNS record this entity represents. The value can be CName, Mx, Srv, or Txt.
	RecordType nullable.Type[string] `json:"recordType,omitempty"`

	// Microsoft Online Service or feature that has a dependency on this DNS record. Can be one of the following values:
	// null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain,
	// FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune.
	SupportedService *string `json:"supportedService,omitempty"`

	// Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable.
	Ttl *int64 `json:"ttl,omitempty"`

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

func (s BaseDomainDnsRecordImpl) DomainDnsRecord() BaseDomainDnsRecordImpl {
	return s
}

func (s BaseDomainDnsRecordImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DomainDnsRecord = RawDomainDnsRecordImpl{}

// RawDomainDnsRecordImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDomainDnsRecordImpl struct {
	domainDnsRecord BaseDomainDnsRecordImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawDomainDnsRecordImpl) DomainDnsRecord() BaseDomainDnsRecordImpl {
	return s.domainDnsRecord
}

func (s RawDomainDnsRecordImpl) Entity() BaseEntityImpl {
	return s.domainDnsRecord.Entity()
}

var _ json.Marshaler = BaseDomainDnsRecordImpl{}

func (s BaseDomainDnsRecordImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDomainDnsRecordImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDomainDnsRecordImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDomainDnsRecordImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.domainDnsRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDomainDnsRecordImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDomainDnsRecordImplementation(input []byte) (DomainDnsRecord, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DomainDnsRecord into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsCnameRecord") {
		var out DomainDnsCnameRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsCnameRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsMxRecord") {
		var out DomainDnsMxRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsMxRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsSrvRecord") {
		var out DomainDnsSrvRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsSrvRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsTxtRecord") {
		var out DomainDnsTxtRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsTxtRecord: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainDnsUnavailableRecord") {
		var out DomainDnsUnavailableRecord
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainDnsUnavailableRecord: %+v", err)
		}
		return out, nil
	}

	var parent BaseDomainDnsRecordImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDomainDnsRecordImpl: %+v", err)
	}

	return RawDomainDnsRecordImpl{
		domainDnsRecord: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
