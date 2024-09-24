package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Organization{}

type Organization struct {
	// The collection of service plans associated with the tenant. Not nullable.
	AssignedPlans *[]AssignedPlan `json:"assignedPlans,omitempty"`

	// Branding for the organization. Nullable.
	Branding *OrganizationalBranding `json:"branding,omitempty"`

	// Telephone number for the organization. Although this property is a string collection, only one number can be set.
	BusinessPhones *[]string `json:"businessPhones,omitempty"`

	// Navigation property to manage certificate-based authentication configuration. Only a single instance of
	// certificateBasedAuthConfiguration can be created in the collection.
	CertificateBasedAuthConfiguration *[]CertificateBasedAuthConfiguration `json:"certificateBasedAuthConfiguration,omitempty"`

	// City name of the address for the organization.
	City nullable.Type[string] `json:"city,omitempty"`

	// Country or region name of the address for the organization.
	Country nullable.Type[string] `json:"country,omitempty"`

	// Country or region abbreviation for the organization in ISO 3166-2 format.
	CountryLetterCode nullable.Type[string] `json:"countryLetterCode,omitempty"`

	// Timestamp of when the organization was created. The value can't be modified and is automatically populated when the
	// organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always
	// in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Two-letter ISO 3166 country code indicating the default service usage location of an organization.
	DefaultUsageLocation nullable.Type[string] `json:"defaultUsageLocation,omitempty"`

	// The display name for the tenant.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The collection of open extensions defined for the organization. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// Not nullable.
	MarketingNotificationEmails *[]string `json:"marketingNotificationEmails,omitempty"`

	// Mobile device management authority.
	MobileDeviceManagementAuthority *MdmAuthority `json:"mobileDeviceManagementAuthority,omitempty"`

	// The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014
	// is 2014-01-01T00:00:00Z. Read-only.
	OnPremisesLastSyncDateTime nullable.Type[string] `json:"onPremisesLastSyncDateTime,omitempty"`

	// true if this object is synced from an on-premises directory; false if this object was originally synced from an
	// on-premises directory but is no longer synced. Nullable. null if this object isn't synced from on-premises active
	// directory (default).
	OnPremisesSyncEnabled nullable.Type[bool] `json:"onPremisesSyncEnabled,omitempty"`

	// The type of partnership this tenant has with Microsoft. The possible values are: microsoftSupport, syndicatePartner,
	// breadthPartner, breadthPartnerDelegatedAdmin, resellerPartnerDelegatedAdmin, valueAddedResellerPartnerDelegatedAdmin,
	// unknownFutureValue. Nullable. For more information about the possible types, see partnerTenantType values.
	PartnerTenantType *PartnerTenantType `json:"partnerTenantType,omitempty"`

	// Postal code of the address for the organization.
	PostalCode nullable.Type[string] `json:"postalCode,omitempty"`

	// The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
	PreferredLanguage nullable.Type[string] `json:"preferredLanguage,omitempty"`

	// The privacy profile of an organization.
	PrivacyProfile *PrivacyProfile `json:"privacyProfile,omitempty"`

	// Not nullable.
	ProvisionedPlans *[]ProvisionedPlan `json:"provisionedPlans,omitempty"`

	// Not nullable.
	SecurityComplianceNotificationMails *[]string `json:"securityComplianceNotificationMails,omitempty"`

	// Not nullable.
	SecurityComplianceNotificationPhones *[]string `json:"securityComplianceNotificationPhones,omitempty"`

	// State name of the address for the organization.
	State nullable.Type[string] `json:"state,omitempty"`

	// Street name of the address for organization.
	Street nullable.Type[string] `json:"street,omitempty"`

	// Not nullable.
	TechnicalNotificationMails *[]string `json:"technicalNotificationMails,omitempty"`

	// Not nullable. Can be one of the following types: AAD - An enterprise identity access management (IAM) service that
	// serves business-to-employee and business-to-business (B2B) scenarios. AAD B2C An identity access management (IAM)
	// service that serves business-to-consumer (B2C) scenarios. CIAM - A customer identity & access management (CIAM)
	// solution that provides an integrated platform to serve consumers, partners, and citizen scenarios.
	TenantType nullable.Type[string] `json:"tenantType,omitempty"`

	// The collection of domains associated with this tenant. Not nullable.
	VerifiedDomains *[]VerifiedDomain `json:"verifiedDomains,omitempty"`

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

func (s Organization) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Organization) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Organization{}

func (s Organization) MarshalJSON() ([]byte, error) {
	type wrapper Organization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Organization: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Organization: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "extensions")
	delete(decoded, "onPremisesLastSyncDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.organization"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Organization: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Organization{}

func (s *Organization) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedPlans                        *[]AssignedPlan                      `json:"assignedPlans,omitempty"`
		Branding                             *OrganizationalBranding              `json:"branding,omitempty"`
		BusinessPhones                       *[]string                            `json:"businessPhones,omitempty"`
		CertificateBasedAuthConfiguration    *[]CertificateBasedAuthConfiguration `json:"certificateBasedAuthConfiguration,omitempty"`
		City                                 nullable.Type[string]                `json:"city,omitempty"`
		Country                              nullable.Type[string]                `json:"country,omitempty"`
		CountryLetterCode                    nullable.Type[string]                `json:"countryLetterCode,omitempty"`
		CreatedDateTime                      nullable.Type[string]                `json:"createdDateTime,omitempty"`
		DefaultUsageLocation                 nullable.Type[string]                `json:"defaultUsageLocation,omitempty"`
		DisplayName                          nullable.Type[string]                `json:"displayName,omitempty"`
		MarketingNotificationEmails          *[]string                            `json:"marketingNotificationEmails,omitempty"`
		MobileDeviceManagementAuthority      *MdmAuthority                        `json:"mobileDeviceManagementAuthority,omitempty"`
		OnPremisesLastSyncDateTime           nullable.Type[string]                `json:"onPremisesLastSyncDateTime,omitempty"`
		OnPremisesSyncEnabled                nullable.Type[bool]                  `json:"onPremisesSyncEnabled,omitempty"`
		PartnerTenantType                    *PartnerTenantType                   `json:"partnerTenantType,omitempty"`
		PostalCode                           nullable.Type[string]                `json:"postalCode,omitempty"`
		PreferredLanguage                    nullable.Type[string]                `json:"preferredLanguage,omitempty"`
		PrivacyProfile                       *PrivacyProfile                      `json:"privacyProfile,omitempty"`
		ProvisionedPlans                     *[]ProvisionedPlan                   `json:"provisionedPlans,omitempty"`
		SecurityComplianceNotificationMails  *[]string                            `json:"securityComplianceNotificationMails,omitempty"`
		SecurityComplianceNotificationPhones *[]string                            `json:"securityComplianceNotificationPhones,omitempty"`
		State                                nullable.Type[string]                `json:"state,omitempty"`
		Street                               nullable.Type[string]                `json:"street,omitempty"`
		TechnicalNotificationMails           *[]string                            `json:"technicalNotificationMails,omitempty"`
		TenantType                           nullable.Type[string]                `json:"tenantType,omitempty"`
		VerifiedDomains                      *[]VerifiedDomain                    `json:"verifiedDomains,omitempty"`
		DeletedDateTime                      nullable.Type[string]                `json:"deletedDateTime,omitempty"`
		Id                                   *string                              `json:"id,omitempty"`
		ODataId                              *string                              `json:"@odata.id,omitempty"`
		ODataType                            *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedPlans = decoded.AssignedPlans
	s.Branding = decoded.Branding
	s.BusinessPhones = decoded.BusinessPhones
	s.CertificateBasedAuthConfiguration = decoded.CertificateBasedAuthConfiguration
	s.City = decoded.City
	s.Country = decoded.Country
	s.CountryLetterCode = decoded.CountryLetterCode
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefaultUsageLocation = decoded.DefaultUsageLocation
	s.DisplayName = decoded.DisplayName
	s.MarketingNotificationEmails = decoded.MarketingNotificationEmails
	s.MobileDeviceManagementAuthority = decoded.MobileDeviceManagementAuthority
	s.OnPremisesLastSyncDateTime = decoded.OnPremisesLastSyncDateTime
	s.OnPremisesSyncEnabled = decoded.OnPremisesSyncEnabled
	s.PartnerTenantType = decoded.PartnerTenantType
	s.PostalCode = decoded.PostalCode
	s.PreferredLanguage = decoded.PreferredLanguage
	s.PrivacyProfile = decoded.PrivacyProfile
	s.ProvisionedPlans = decoded.ProvisionedPlans
	s.SecurityComplianceNotificationMails = decoded.SecurityComplianceNotificationMails
	s.SecurityComplianceNotificationPhones = decoded.SecurityComplianceNotificationPhones
	s.State = decoded.State
	s.Street = decoded.Street
	s.TechnicalNotificationMails = decoded.TechnicalNotificationMails
	s.TenantType = decoded.TenantType
	s.VerifiedDomains = decoded.VerifiedDomains
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Organization into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'Organization': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	return nil
}
