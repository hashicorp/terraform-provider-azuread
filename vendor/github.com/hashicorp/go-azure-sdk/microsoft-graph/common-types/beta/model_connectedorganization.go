package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ConnectedOrganization{}

type ConnectedOrganization struct {
	// UPN of the user who created this resource. Read-only.
	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of the connected organization.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the connected organization. Supports $filter (eq).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	ExternalSponsors *[]DirectoryObject `json:"externalSponsors,omitempty"`

	// List of OData IDs for `ExternalSponsors` to bind to this entity
	ExternalSponsors_ODataBind *[]string `json:"externalSponsors@odata.bind,omitempty"`

	// The identity sources in this connected organization, one of azureActiveDirectoryTenant,
	// crossCloudAzureActiveDirectoryTenant, domainIdentitySource, externalDomainFederation, or socialIdentitySource.
	// Read-only. Nullable. Supports $select and $filter(eq). To filter by the derived types, you must declare the resource
	// using its full OData cast, for example,
	// $filter=identitySources/any(is:is/microsoft.graph.azureActiveDirectoryTenant/tenantId eq
	// 'bcfdfff4-cbc3-43f2-9000-ba7b7515054f').
	IdentitySources *[]IdentitySource `json:"identitySources,omitempty"`

	InternalSponsors *[]DirectoryObject `json:"internalSponsors,omitempty"`

	// List of OData IDs for `InternalSponsors` to bind to this entity
	InternalSponsors_ODataBind *[]string `json:"internalSponsors@odata.bind,omitempty"`

	// UPN of the user who last modified this resource. Read-only.
	ModifiedBy nullable.Type[string] `json:"modifiedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// The state of a connected organization defines whether assignment policies with requestor scope type
	// AllConfiguredConnectedOrganizationSubjects are applicable or not. Possible values are: configured, proposed.
	State *ConnectedOrganizationState `json:"state,omitempty"`

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

func (s ConnectedOrganization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConnectedOrganization{}

func (s ConnectedOrganization) MarshalJSON() ([]byte, error) {
	type wrapper ConnectedOrganization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectedOrganization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectedOrganization: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "identitySources")
	delete(decoded, "modifiedBy")
	delete(decoded, "modifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.connectedOrganization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectedOrganization: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ConnectedOrganization{}

func (s *ConnectedOrganization) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedBy                  nullable.Type[string]       `json:"createdBy,omitempty"`
		CreatedDateTime            nullable.Type[string]       `json:"createdDateTime,omitempty"`
		Description                nullable.Type[string]       `json:"description,omitempty"`
		DisplayName                nullable.Type[string]       `json:"displayName,omitempty"`
		ExternalSponsors_ODataBind *[]string                   `json:"externalSponsors@odata.bind,omitempty"`
		InternalSponsors_ODataBind *[]string                   `json:"internalSponsors@odata.bind,omitempty"`
		ModifiedBy                 nullable.Type[string]       `json:"modifiedBy,omitempty"`
		ModifiedDateTime           nullable.Type[string]       `json:"modifiedDateTime,omitempty"`
		State                      *ConnectedOrganizationState `json:"state,omitempty"`
		Id                         *string                     `json:"id,omitempty"`
		ODataId                    *string                     `json:"@odata.id,omitempty"`
		ODataType                  *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExternalSponsors_ODataBind = decoded.ExternalSponsors_ODataBind
	s.InternalSponsors_ODataBind = decoded.InternalSponsors_ODataBind
	s.ModifiedBy = decoded.ModifiedBy
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.State = decoded.State
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ConnectedOrganization into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["externalSponsors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ExternalSponsors into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ExternalSponsors' for 'ConnectedOrganization': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ExternalSponsors = &output
	}

	if v, ok := temp["identitySources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IdentitySources into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentitySource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentitySourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IdentitySources' for 'ConnectedOrganization': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IdentitySources = &output
	}

	if v, ok := temp["internalSponsors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling InternalSponsors into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'InternalSponsors' for 'ConnectedOrganization': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.InternalSponsors = &output
	}

	return nil
}
