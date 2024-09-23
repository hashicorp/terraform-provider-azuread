package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DomainDnsRecord = DomainDnsTxtRecord{}

type DomainDnsTxtRecord struct {
	// Value used when configuring the text property at the DNS host.
	Text *string `json:"text,omitempty"`

	// Fields inherited from DomainDnsRecord

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

func (s DomainDnsTxtRecord) DomainDnsRecord() BaseDomainDnsRecordImpl {
	return BaseDomainDnsRecordImpl{
		IsOptional:       s.IsOptional,
		Label:            s.Label,
		RecordType:       s.RecordType,
		SupportedService: s.SupportedService,
		Ttl:              s.Ttl,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s DomainDnsTxtRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DomainDnsTxtRecord{}

func (s DomainDnsTxtRecord) MarshalJSON() ([]byte, error) {
	type wrapper DomainDnsTxtRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DomainDnsTxtRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DomainDnsTxtRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.domainDnsTxtRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DomainDnsTxtRecord: %+v", err)
	}

	return encoded, nil
}
