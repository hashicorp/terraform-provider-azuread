package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudCertificationAuthority{}

type CloudCertificationAuthority struct {
	// The URL to download the certification authority certificate. Read-only.
	CertificateDownloadUrl nullable.Type[string] `json:"certificateDownloadUrl,omitempty"`

	// Enum of possible cloud certification authority certificate cryptography and key size combinations.
	CertificateKeySize *CloudCertificationAuthorityCertificateKeySize `json:"certificateKeySize,omitempty"`

	// The cloud certification authority's Certificate Revocation List URL that can be used to determine revocation status.
	// Read-only.
	CertificateRevocationListUrl nullable.Type[string] `json:"certificateRevocationListUrl,omitempty"`

	// The certificate signing request used to create an issuing certification authority with a root certification authority
	// external to Microsoft Cloud PKI. The based-64 encoded certificate signing request can be downloaded through this
	// property. After downloading the certificate signing request, it must be signed by the external root certifcation
	// authority. Read-only.
	CertificateSigningRequest nullable.Type[string] `json:"certificateSigningRequest,omitempty"`

	// Issuer (parent) certification authority identifier. Nullable. Read-only. Supports $orderby and $select.
	CertificationAuthorityIssuerId nullable.Type[string] `json:"certificationAuthorityIssuerId,omitempty"`

	// The URI of the issuing certification authority of a subordinate certification authority. Returns null if a root
	// certification authority. Nullable. Read-only.
	CertificationAuthorityIssuerUri nullable.Type[string] `json:"certificationAuthorityIssuerUri,omitempty"`

	// Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is
	// currently able to issue certificates or temporarily paused or permanently revoked.
	CertificationAuthorityStatus *CloudCertificationAuthorityStatus `json:"certificationAuthorityStatus,omitempty"`

	// Enum type of possible certificate hashing algorithms used by the certification authority to create certificates.
	CloudCertificationAuthorityHashingAlgorithm *CloudCertificationAuthorityHashingAlgorithm `json:"cloudCertificationAuthorityHashingAlgorithm,omitempty"`

	// Required OData property to expose leaf certificate API.
	CloudCertificationAuthorityLeafCertificate *[]CloudCertificationAuthorityLeafCertificate `json:"cloudCertificationAuthorityLeafCertificate,omitempty"`

	// Enum type of possible certificate authority types. This feature supports a two-tier certification authority model
	// with a root certification authority and one or more child issuing (intermediate) certification authorities.
	CloudCertificationAuthorityType *CloudCertificationAuthorityType `json:"cloudCertificationAuthorityType,omitempty"`

	// The common name of the certificate subject name, which must be unique. This property is a relative distinguished name
	// used to compose the certificate subject name. Read-only. Supports $select.
	CommonName nullable.Type[string] `json:"commonName,omitempty"`

	// The country name that is used to compose the subject name of a certification authority certificate in the form 'C='.
	// Nullable. Example: US. Read-only.
	CountryName nullable.Type[string] `json:"countryName,omitempty"`

	// Creation date of this cloud certification authority entity instance. The DateTimeOffset type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like
	// this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The certification authority description displayed in the Intune admin console. Nullable. Read/write. Returns null if
	// not set.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The certification authority display name the Intune admin console. Read/write. Supports $select and $orderby.
	DisplayName *string `json:"displayName,omitempty"`

	// ETag for optimistic concurrency control. Read/write.
	ETag nullable.Type[string] `json:"eTag,omitempty"`

	// The certificate extended key usages, which specify the usage capabilities of the certificate. Read-only.
	ExtendedKeyUsages *[]ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`

	IssuerCommonName nullable.Type[string] `json:"issuerCommonName,omitempty"`

	// Enum type of possible key platforms used by the certification authority.
	KeyPlatform *CloudCertificationAuthorityKeyPlatformType `json:"keyPlatform,omitempty"`

	// Last modification date and time of this certification authority entity instance. The DateTimeOffset type represents
	// date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014
	// would look like this: '2014-01-01T00:00:00Z'. Nullable. Read/write.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The locality (town, city, etc.) name that is used to compose the subject name of a certification authority
	// certificate in the form 'L='. This is Nullable. Example: Redmond. Read-only.
	LocalityName nullable.Type[string] `json:"localityName,omitempty"`

	// The Online Certificate Status Protocol (OCSP) responder URI that can be used to determine certificate status.
	// Read-only.
	OcspResponderUri nullable.Type[string] `json:"ocspResponderUri,omitempty"`

	// The organization name that is used as a distinguished name in the subject name of a certification authority
	// certificate in the form 'O='. Nullable. Example: Microsoft. Read-only.
	OrganizationName nullable.Type[string] `json:"organizationName,omitempty"`

	// The organization unit name that is used as a distinguished name in the subject name of a certification authority
	// certificate in the form 'OU='. Nullable. Example: Security. Read-only.
	OrganizationUnit nullable.Type[string] `json:"organizationUnit,omitempty"`

	// List of Scope Tags for this entity instance. Scope tags limit access to an entity instance. Nullable. Read/write.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The common name of the certificate subject name of the certification authority issuer. This property can be used to
	// identify the certification authority that issued the current certification authority. For issuing certification
	// authorities, this is the common name of the certificate subject name of the root certification authority to which it
	// is anchored. For externally signed certification authorities, this is the common name of the certificate subject name
	// of the signing certification authority. For root certification authorities, this is the common name of the
	// certification authority's own certificate subject name. Read-only.
	RootCertificateCommonName nullable.Type[string] `json:"rootCertificateCommonName,omitempty"`

	// The SCEP server URL for device SCEP connections to request certificates. Read-only.
	ScepServerUrl nullable.Type[string] `json:"scepServerUrl,omitempty"`

	// The serial number used to uniquely identify a certificate with its issuing certification authority. Read-only.
	// Supports $select.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// The state or province name that is used to compose the subject name of a certification authority certificate in the
	// form 'ST='. Nullable. Example: Washington. Read-only.
	StateName nullable.Type[string] `json:"stateName,omitempty"`

	// The subject name of the certificate. The subject is the target or intended beneficiary of the security being
	// provided, such as a company or government entity. Read-only. Supports $orderby and $select.
	SubjectName nullable.Type[string] `json:"subjectName,omitempty"`

	// Secure Hash Algorithm 1 digest of the certificate that can be used to identify it. Read-only. Supports $select.
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`

	// The end date time of the validity period of a certification authority certificate. Certificates cannot be used after
	// this date time as they are longer valid. The DateTimeOffset type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
	ValidityEndDateTime nullable.Type[string] `json:"validityEndDateTime,omitempty"`

	// The certification authority validity period in years configured by admins.
	ValidityPeriodInYears *int64 `json:"validityPeriodInYears,omitempty"`

	// The start date time of the validity period of a certification authority certificate. Certificates cannot be used
	// before this date time as they are not yet valid. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
	ValidityStartDateTime nullable.Type[string] `json:"validityStartDateTime,omitempty"`

	// The certification authority version, which is incremented each time the certification authority is renewed.
	// Read-only.
	VersionNumber *int64 `json:"versionNumber,omitempty"`

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

func (s CloudCertificationAuthority) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudCertificationAuthority{}

func (s CloudCertificationAuthority) MarshalJSON() ([]byte, error) {
	type wrapper CloudCertificationAuthority
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudCertificationAuthority: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudCertificationAuthority: %+v", err)
	}

	delete(decoded, "certificateDownloadUrl")
	delete(decoded, "certificateRevocationListUrl")
	delete(decoded, "certificateSigningRequest")
	delete(decoded, "certificationAuthorityIssuerId")
	delete(decoded, "certificationAuthorityIssuerUri")
	delete(decoded, "commonName")
	delete(decoded, "countryName")
	delete(decoded, "createdDateTime")
	delete(decoded, "extendedKeyUsages")
	delete(decoded, "localityName")
	delete(decoded, "ocspResponderUri")
	delete(decoded, "organizationName")
	delete(decoded, "organizationUnit")
	delete(decoded, "rootCertificateCommonName")
	delete(decoded, "scepServerUrl")
	delete(decoded, "serialNumber")
	delete(decoded, "stateName")
	delete(decoded, "subjectName")
	delete(decoded, "thumbprint")
	delete(decoded, "validityEndDateTime")
	delete(decoded, "validityStartDateTime")
	delete(decoded, "versionNumber")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudCertificationAuthority"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudCertificationAuthority: %+v", err)
	}

	return encoded, nil
}
