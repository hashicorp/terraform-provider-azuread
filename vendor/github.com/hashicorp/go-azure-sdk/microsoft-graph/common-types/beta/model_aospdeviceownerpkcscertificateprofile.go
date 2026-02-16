package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AospDeviceOwnerCertificateProfileBase = AospDeviceOwnerPkcsCertificateProfile{}

type AospDeviceOwnerPkcsCertificateProfile struct {
	// CertificateStore types
	CertificateStore *CertificateStore `json:"certificateStore,omitempty"`

	// PKCS Certificate Template Name
	CertificateTemplateName *string `json:"certificateTemplateName,omitempty"`

	// PKCS Certification Authority
	CertificationAuthority *string `json:"certificationAuthority,omitempty"`

	// PKCS Certification Authority Name
	CertificationAuthorityName nullable.Type[string] `json:"certificationAuthorityName,omitempty"`

	// Device Management Certification Authority Types.
	CertificationAuthorityType *DeviceManagementCertificationAuthority `json:"certificationAuthorityType,omitempty"`

	// Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
	CustomSubjectAlternativeNames *[]CustomSubjectAlternativeName `json:"customSubjectAlternativeNames,omitempty"`

	// Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
	ManagedDeviceCertificateStates *[]ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`

	// Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString nullable.Type[string] `json:"subjectAlternativeNameFormatString,omitempty"`

	// Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise
	// Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
	SubjectNameFormatString *string `json:"subjectNameFormatString,omitempty"`

	// Fields inherited from AospDeviceOwnerCertificateProfileBase

	// Certificate Validity Period Options.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`

	// Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int64 `json:"certificateValidityPeriodValue,omitempty"`

	// Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages *[]ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`

	// Certificate renewal threshold percentage. Valid values 1 to 99
	RenewalThresholdPercentage *int64 `json:"renewalThresholdPercentage,omitempty"`

	// Trusted Root Certificate.
	RootCertificate *AospDeviceOwnerTrustedRootCertificate `json:"rootCertificate,omitempty"`

	// Certificate Subject Alternative Name Type. This collection can contain a maximum of 500 elements. Possible values
	// are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`

	// Certificate Subject Name Format. This collection can contain a maximum of 500 elements. Possible values are:
	// commonName, commonNameIncludingEmail, commonNameAsEmail, custom, commonNameAsIMEI, commonNameAsSerialNumber,
	// commonNameAsAadDeviceId, commonNameAsIntuneDeviceId, commonNameAsDurableDeviceId.
	SubjectNameFormat *SubjectNameFormat `json:"subjectNameFormat,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

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

func (s AospDeviceOwnerPkcsCertificateProfile) AospDeviceOwnerCertificateProfileBase() BaseAospDeviceOwnerCertificateProfileBaseImpl {
	return BaseAospDeviceOwnerCertificateProfileBaseImpl{
		CertificateValidityPeriodScale:              s.CertificateValidityPeriodScale,
		CertificateValidityPeriodValue:              s.CertificateValidityPeriodValue,
		ExtendedKeyUsages:                           s.ExtendedKeyUsages,
		RenewalThresholdPercentage:                  s.RenewalThresholdPercentage,
		RootCertificate:                             s.RootCertificate,
		SubjectAlternativeNameType:                  s.SubjectAlternativeNameType,
		SubjectNameFormat:                           s.SubjectNameFormat,
		Assignments:                                 s.Assignments,
		CreatedDateTime:                             s.CreatedDateTime,
		Description:                                 s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s AospDeviceOwnerPkcsCertificateProfile) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s AospDeviceOwnerPkcsCertificateProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AospDeviceOwnerPkcsCertificateProfile{}

func (s AospDeviceOwnerPkcsCertificateProfile) MarshalJSON() ([]byte, error) {
	type wrapper AospDeviceOwnerPkcsCertificateProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AospDeviceOwnerPkcsCertificateProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AospDeviceOwnerPkcsCertificateProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aospDeviceOwnerPkcsCertificateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AospDeviceOwnerPkcsCertificateProfile: %+v", err)
	}

	return encoded, nil
}
