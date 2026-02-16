package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Windows10XCertificateProfile = Windows10XSCEPCertificateProfile{}

type Windows10XSCEPCertificateProfile struct {
	// CertificateStore types
	CertificateStore *CertificateStore `json:"certificateStore,omitempty"`

	// Certificate Validity Period Options.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`

	// Value for the Certificate Validity Period
	CertificateValidityPeriodValue *int64 `json:"certificateValidityPeriodValue,omitempty"`

	// Extended Key Usage (EKU) settings.
	ExtendedKeyUsages *[]ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`

	// SCEP Hash Algorithm.
	HashAlgorithm *[]HashAlgorithms `json:"hashAlgorithm,omitempty"`

	// Key Size Options.
	KeySize *KeySize `json:"keySize,omitempty"`

	// Key Storage Provider (KSP) Import Options.
	KeyStorageProvider *KeyStorageProviderOption `json:"keyStorageProvider,omitempty"`

	// Key Usage Options.
	KeyUsage *KeyUsages `json:"keyUsage,omitempty"`

	// Certificate renewal threshold percentage
	RenewalThresholdPercentage *int64 `json:"renewalThresholdPercentage,omitempty"`

	// Trusted Root Certificate ID
	RootCertificateId nullable.Type[string] `json:"rootCertificateId,omitempty"`

	// SCEP Server Url(s).
	ScepServerUrls *[]string `json:"scepServerUrls,omitempty"`

	// Custom AAD Attributes.
	SubjectAlternativeNameFormats *[]Windows10XCustomSubjectAlternativeName `json:"subjectAlternativeNameFormats,omitempty"`

	// Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise
	// Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
	SubjectNameFormatString nullable.Type[string] `json:"subjectNameFormatString,omitempty"`

	// Fields inherited from DeviceManagementResourceAccessProfileBase

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceManagementResourceAccessProfileAssignment `json:"assignments,omitempty"`

	// DateTime profile was created
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`

	// Profile description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Profile display name
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime profile was last modified
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Scope Tags
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Version of the profile
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

func (s Windows10XSCEPCertificateProfile) Windows10XCertificateProfile() BaseWindows10XCertificateProfileImpl {
	return BaseWindows10XCertificateProfileImpl{
		Assignments:          s.Assignments,
		CreationDateTime:     s.CreationDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Windows10XSCEPCertificateProfile) DeviceManagementResourceAccessProfileBase() BaseDeviceManagementResourceAccessProfileBaseImpl {
	return BaseDeviceManagementResourceAccessProfileBaseImpl{
		Assignments:          s.Assignments,
		CreationDateTime:     s.CreationDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		RoleScopeTagIds:      s.RoleScopeTagIds,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Windows10XSCEPCertificateProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10XSCEPCertificateProfile{}

func (s Windows10XSCEPCertificateProfile) MarshalJSON() ([]byte, error) {
	type wrapper Windows10XSCEPCertificateProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10XSCEPCertificateProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10XSCEPCertificateProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10XSCEPCertificateProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10XSCEPCertificateProfile: %+v", err)
	}

	return encoded, nil
}
