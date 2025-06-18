package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = OrgContact{}

type OrgContact struct {
	// Postal addresses for this organizational contact. For now a contact can only have one physical address.
	Addresses *[]PhysicalOfficeAddress `json:"addresses,omitempty"`

	// Name of the company that this organizational contact belong to. Supports $filter (eq, ne, not, ge, le, in,
	// startsWith, and eq for null values).
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`

	// The name for the department in which the contact works. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq
	// for null values).
	Department nullable.Type[string] `json:"department,omitempty"`

	// The contact's direct reports. (The users and contacts that have their manager property set to this contact.)
	// Read-only. Nullable. Supports $expand.
	DirectReports *[]DirectoryObject `json:"directReports,omitempty"`

	// List of OData IDs for `DirectReports` to bind to this entity
	DirectReports_ODataBind *[]string `json:"directReports@odata.bind,omitempty"`

	// Display name for this organizational contact. Maximum length is 256 characters. Supports $filter (eq, ne, not, ge,
	// le, in, startsWith, and eq for null values), $search, and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// First name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null
	// values).
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// Job title for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null
	// values).
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	// The SMTP address for the contact, for example, 'jeff@contoso.com'. Supports $filter (eq, ne, not, ge, le, in,
	// startsWith, and eq for null values).
	Mail nullable.Type[string] `json:"mail,omitempty"`

	// Email alias (portion of email address pre-pending the @ symbol) for this organizational contact. Supports $filter
	// (eq, ne, not, ge, le, in, startsWith, and eq for null values).
	MailNickname nullable.Type[string] `json:"mailNickname,omitempty"`

	// The user or contact that is this contact's manager. Read-only. Supports $expand and $filter (eq) by id.
	Manager *DirectoryObject `json:"manager,omitempty"`

	// OData ID for `Manager` to bind to this entity
	Manager_ODataBind *string `json:"manager@odata.bind,omitempty"`

	// Groups that this contact is a member of. Read-only. Nullable. Supports $expand.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`

	// List of OData IDs for `MemberOf` to bind to this entity
	MemberOf_ODataBind *[]string `json:"memberOf@odata.bind,omitempty"`

	// Date and time when this organizational contact was last synchronized from on-premises AD. The Timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, not, ge, le, in).
	OnPremisesLastSyncDateTime nullable.Type[string] `json:"onPremisesLastSyncDateTime,omitempty"`

	// List of any synchronization provisioning errors for this organizational contact. Supports $filter (eq, not for
	// category and propertyCausingError), /$count eq 0, /$count ne 0.
	OnPremisesProvisioningErrors *[]OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`

	// true if this object is synced from an on-premises directory; false if this object was originally synced from an
	// on-premises directory but is no longer synced and now mastered in Exchange; null if this object has never been synced
	// from an on-premises directory (default). Supports $filter (eq, ne, not, in, and eq for null values).
	OnPremisesSyncEnabled nullable.Type[bool] `json:"onPremisesSyncEnabled,omitempty"`

	// List of phones for this organizational contact. Phone types can be mobile, business, and businessFax. Only one of
	// each type can ever be present in the collection. Supports $filter (eq, ne, not, in).
	Phones *[]Phone `json:"phones,omitempty"`

	// For example: 'SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com'. The any operator is required for filter
	// expressions on multi-valued properties. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0).
	ProxyAddresses *[]string `json:"proxyAddresses,omitempty"`

	// Errors published by a federated service describing a non-transient, service-specific error regarding the properties
	// or link from an orgContact object . Supports $filter (eq, not, for isResolved and serviceInstance).
	ServiceProvisioningErrors *[]ServiceProvisioningError `json:"serviceProvisioningErrors,omitempty"`

	// Last name for this organizational contact. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq for null
	// values)
	Surname nullable.Type[string] `json:"surname,omitempty"`

	TransitiveMemberOf *[]DirectoryObject `json:"transitiveMemberOf,omitempty"`

	// List of OData IDs for `TransitiveMemberOf` to bind to this entity
	TransitiveMemberOf_ODataBind *[]string `json:"transitiveMemberOf@odata.bind,omitempty"`

	// The transitive reports for a contact. Read-only.
	TransitiveReports *[]DirectoryObject `json:"transitiveReports,omitempty"`

	// List of OData IDs for `TransitiveReports` to bind to this entity
	TransitiveReports_ODataBind *[]string `json:"transitiveReports@odata.bind,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s OrgContact) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s OrgContact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OrgContact{}

func (s OrgContact) MarshalJSON() ([]byte, error) {
	type wrapper OrgContact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OrgContact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OrgContact: %+v", err)
	}

	delete(decoded, "directReports")
	delete(decoded, "manager")
	delete(decoded, "memberOf")
	delete(decoded, "transitiveReports")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.orgContact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OrgContact: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OrgContact{}

func (s *OrgContact) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Addresses                    *[]PhysicalOfficeAddress       `json:"addresses,omitempty"`
		CompanyName                  nullable.Type[string]          `json:"companyName,omitempty"`
		Department                   nullable.Type[string]          `json:"department,omitempty"`
		DirectReports_ODataBind      *[]string                      `json:"directReports@odata.bind,omitempty"`
		DisplayName                  nullable.Type[string]          `json:"displayName,omitempty"`
		GivenName                    nullable.Type[string]          `json:"givenName,omitempty"`
		JobTitle                     nullable.Type[string]          `json:"jobTitle,omitempty"`
		Mail                         nullable.Type[string]          `json:"mail,omitempty"`
		MailNickname                 nullable.Type[string]          `json:"mailNickname,omitempty"`
		Manager_ODataBind            *string                        `json:"manager@odata.bind,omitempty"`
		MemberOf_ODataBind           *[]string                      `json:"memberOf@odata.bind,omitempty"`
		OnPremisesLastSyncDateTime   nullable.Type[string]          `json:"onPremisesLastSyncDateTime,omitempty"`
		OnPremisesProvisioningErrors *[]OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
		OnPremisesSyncEnabled        nullable.Type[bool]            `json:"onPremisesSyncEnabled,omitempty"`
		Phones                       *[]Phone                       `json:"phones,omitempty"`
		ProxyAddresses               *[]string                      `json:"proxyAddresses,omitempty"`
		Surname                      nullable.Type[string]          `json:"surname,omitempty"`
		TransitiveMemberOf_ODataBind *[]string                      `json:"transitiveMemberOf@odata.bind,omitempty"`
		TransitiveReports_ODataBind  *[]string                      `json:"transitiveReports@odata.bind,omitempty"`
		DeletedDateTime              nullable.Type[string]          `json:"deletedDateTime,omitempty"`
		Id                           *string                        `json:"id,omitempty"`
		ODataId                      *string                        `json:"@odata.id,omitempty"`
		ODataType                    *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Addresses = decoded.Addresses
	s.CompanyName = decoded.CompanyName
	s.Department = decoded.Department
	s.DirectReports_ODataBind = decoded.DirectReports_ODataBind
	s.DisplayName = decoded.DisplayName
	s.GivenName = decoded.GivenName
	s.JobTitle = decoded.JobTitle
	s.Mail = decoded.Mail
	s.MailNickname = decoded.MailNickname
	s.Manager_ODataBind = decoded.Manager_ODataBind
	s.MemberOf_ODataBind = decoded.MemberOf_ODataBind
	s.OnPremisesLastSyncDateTime = decoded.OnPremisesLastSyncDateTime
	s.OnPremisesProvisioningErrors = decoded.OnPremisesProvisioningErrors
	s.OnPremisesSyncEnabled = decoded.OnPremisesSyncEnabled
	s.Phones = decoded.Phones
	s.ProxyAddresses = decoded.ProxyAddresses
	s.Surname = decoded.Surname
	s.TransitiveMemberOf_ODataBind = decoded.TransitiveMemberOf_ODataBind
	s.TransitiveReports_ODataBind = decoded.TransitiveReports_ODataBind
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OrgContact into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["directReports"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DirectReports into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DirectReports' for 'OrgContact': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DirectReports = &output
	}

	if v, ok := temp["manager"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Manager' for 'OrgContact': %+v", err)
		}
		s.Manager = &impl
	}

	if v, ok := temp["memberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MemberOf' for 'OrgContact': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MemberOf = &output
	}

	if v, ok := temp["serviceProvisioningErrors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ServiceProvisioningErrors into list []json.RawMessage: %+v", err)
		}

		output := make([]ServiceProvisioningError, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalServiceProvisioningErrorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ServiceProvisioningErrors' for 'OrgContact': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ServiceProvisioningErrors = &output
	}

	if v, ok := temp["transitiveMemberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveMemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveMemberOf' for 'OrgContact': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveMemberOf = &output
	}

	if v, ok := temp["transitiveReports"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveReports into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveReports' for 'OrgContact': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveReports = &output
	}

	return nil
}
