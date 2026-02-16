package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Domain{}

type Domain struct {
	// Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed
	// indicates a cloud managed domain where Microsoft Entra ID performs user authentication. Federated indicates
	// authentication is federated with an identity provider such as the tenant's on-premises Active Directory via Active
	// Directory Federation Services. Not nullable. To update this property in delegated scenarios, the calling app must be
	// assigned the Directory.AccessAsUser.All delegated permission.
	AuthenticationType *string `json:"authenticationType,omitempty"`

	// This property is always null except when the verify action is used. When the verify action is used, a domain entity
	// is returned in the response. The availabilityStatus property of the domain entity in the response is either
	// AvailableImmediately or EmailVerifiedDomainTakeoverScheduled.
	AvailabilityStatus nullable.Type[string] `json:"availabilityStatus,omitempty"`

	// The objects such as users and groups that reference the domain ID. Read-only, Nullable. Does not support $expand.
	// Supports $filter by the OData type of objects returned. For example,
	// /domains/{domainId}/domainNameReferences/microsoft.graph.user and
	// /domains/{domainId}/domainNameReferences/microsoft.graph.group.
	DomainNameReferences *[]DirectoryObject `json:"domainNameReferences,omitempty"`

	// List of OData IDs for `DomainNameReferences` to bind to this entity
	DomainNameReferences_ODataBind *[]string `json:"domainNameReferences@odata.bind,omitempty"`

	// Domain settings configured by customer when federated with Microsoft Entra ID. Does not support $expand.
	FederationConfiguration *[]InternalDomainFederation `json:"federationConfiguration,omitempty"`

	// The value of the property is false if the DNS record management of the domain is delegated to Microsoft 365.
	// Otherwise, the value is true. Not nullable.
	IsAdminManaged *bool `json:"isAdminManaged,omitempty"`

	// true for the default domain that is used for user creation. There's only one default domain per company. Not
	// nullable.
	IsDefault *bool `json:"isDefault,omitempty"`

	// true for the initial domain created by Microsoft Online Services. For example, contoso.onmicrosoft.com. There's only
	// one initial domain per company. Not nullable.
	IsInitial *bool `json:"isInitial,omitempty"`

	// true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified. Not
	// nullable.
	IsRoot *bool `json:"isRoot,omitempty"`

	// true for verified domains. Not nullable.
	IsVerified *bool `json:"isVerified,omitempty"`

	// Specifies the number of days before a user receives a password expiry notification. 14 days by default.
	PasswordNotificationWindowInDays nullable.Type[int64] `json:"passwordNotificationWindowInDays,omitempty"`

	// Specifies the length of time that a password is valid before it must be changed. 90 days by default.
	PasswordValidityPeriodInDays nullable.Type[int64] `json:"passwordValidityPeriodInDays,omitempty"`

	// Root domain of a subdomain. Read-only, Nullable. Supports $expand.
	RootDomain *Domain `json:"rootDomain,omitempty"`

	// DNS records the customer adds to the DNS zone file of the domain before the domain can be used by Microsoft Online
	// services. Read-only, Nullable. Does not support $expand.
	ServiceConfigurationRecords *[]DomainDnsRecord `json:"serviceConfigurationRecords,omitempty"`

	SharedEmailDomainInvitations *[]SharedEmailDomainInvitation `json:"sharedEmailDomainInvitations,omitempty"`

	// Status of asynchronous operations scheduled for the domain.
	State *DomainState `json:"state,omitempty"`

	// The capabilities assigned to the domain. Can include 0, 1, or more of following values: Email, Sharepoint,
	// EmailInternalRelayOnly, OfficeCommunicationsOnline,SharePointDefaultDomain, FullRedelegation, SharePointPublic,
	// OrgIdAuthentication, Yammer, Intune, CustomUrlDomain. The values that you can add or remove using the API include:
	// Email, OfficeCommunicationsOnline, Yammer, and CustomUrlDomain. Not nullable. For more information about
	// CustomUrlDomain, see Custom URL domains in external tenants.
	SupportedServices *[]string `json:"supportedServices,omitempty"`

	// DNS records that the customer adds to the DNS zone file of the domain before the customer can complete domain
	// ownership verification with Microsoft Entra ID. Read-only, Nullable. Does not support $expand.
	VerificationDnsRecords *[]DomainDnsRecord `json:"verificationDnsRecords,omitempty"`

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

func (s Domain) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Domain{}

func (s Domain) MarshalJSON() ([]byte, error) {
	type wrapper Domain
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Domain: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Domain: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.domain"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Domain: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Domain{}

func (s *Domain) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationType               *string                        `json:"authenticationType,omitempty"`
		AvailabilityStatus               nullable.Type[string]          `json:"availabilityStatus,omitempty"`
		DomainNameReferences_ODataBind   *[]string                      `json:"domainNameReferences@odata.bind,omitempty"`
		FederationConfiguration          *[]InternalDomainFederation    `json:"federationConfiguration,omitempty"`
		IsAdminManaged                   *bool                          `json:"isAdminManaged,omitempty"`
		IsDefault                        *bool                          `json:"isDefault,omitempty"`
		IsInitial                        *bool                          `json:"isInitial,omitempty"`
		IsRoot                           *bool                          `json:"isRoot,omitempty"`
		IsVerified                       *bool                          `json:"isVerified,omitempty"`
		PasswordNotificationWindowInDays nullable.Type[int64]           `json:"passwordNotificationWindowInDays,omitempty"`
		PasswordValidityPeriodInDays     nullable.Type[int64]           `json:"passwordValidityPeriodInDays,omitempty"`
		RootDomain                       *Domain                        `json:"rootDomain,omitempty"`
		SharedEmailDomainInvitations     *[]SharedEmailDomainInvitation `json:"sharedEmailDomainInvitations,omitempty"`
		State                            *DomainState                   `json:"state,omitempty"`
		SupportedServices                *[]string                      `json:"supportedServices,omitempty"`
		Id                               *string                        `json:"id,omitempty"`
		ODataId                          *string                        `json:"@odata.id,omitempty"`
		ODataType                        *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuthenticationType = decoded.AuthenticationType
	s.AvailabilityStatus = decoded.AvailabilityStatus
	s.DomainNameReferences_ODataBind = decoded.DomainNameReferences_ODataBind
	s.FederationConfiguration = decoded.FederationConfiguration
	s.IsAdminManaged = decoded.IsAdminManaged
	s.IsDefault = decoded.IsDefault
	s.IsInitial = decoded.IsInitial
	s.IsRoot = decoded.IsRoot
	s.IsVerified = decoded.IsVerified
	s.PasswordNotificationWindowInDays = decoded.PasswordNotificationWindowInDays
	s.PasswordValidityPeriodInDays = decoded.PasswordValidityPeriodInDays
	s.RootDomain = decoded.RootDomain
	s.SharedEmailDomainInvitations = decoded.SharedEmailDomainInvitations
	s.State = decoded.State
	s.SupportedServices = decoded.SupportedServices
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Domain into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["domainNameReferences"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DomainNameReferences into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DomainNameReferences' for 'Domain': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DomainNameReferences = &output
	}

	if v, ok := temp["serviceConfigurationRecords"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ServiceConfigurationRecords into list []json.RawMessage: %+v", err)
		}

		output := make([]DomainDnsRecord, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDomainDnsRecordImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ServiceConfigurationRecords' for 'Domain': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ServiceConfigurationRecords = &output
	}

	if v, ok := temp["verificationDnsRecords"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling VerificationDnsRecords into list []json.RawMessage: %+v", err)
		}

		output := make([]DomainDnsRecord, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDomainDnsRecordImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'VerificationDnsRecords' for 'Domain': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.VerificationDnsRecords = &output
	}

	return nil
}
