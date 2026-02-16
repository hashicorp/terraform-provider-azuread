package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudCertificationAuthorityLeafCertificate{}

type CloudCertificationAuthorityLeafCertificate struct {
	// Enum type of possible leaf certificate statuses. These statuses indicate whether certificates are active and usable
	// or unusable if they have been revoked or expired.
	CertificateStatus *CloudCertificationAuthorityLeafCertificateStatus `json:"certificateStatus,omitempty"`

	// The URI of the certification authority that issued the certificate. Read-only.
	CertificationAuthorityIssuerUri nullable.Type[string] `json:"certificationAuthorityIssuerUri,omitempty"`

	// URL to find the relevant Certificate Revocation List for this certificate. Read-only.
	CrlDistributionPointUrl nullable.Type[string] `json:"crlDistributionPointUrl,omitempty"`

	// The unique identifier of the managed device for which the certificate was created. This ID is assigned at device
	// enrollment time. Read-only. Supports $select.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Name of the device for which the certificate was created. Read-only. Supports $select.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The platform of the device for which the certificate was created. Possible values are: Android, AndroidForWork, iOS,
	// MacOS, WindowsPhone81, Windows81AndLater, Windows10AndLater, AndroidWorkProfile, Unknown, AndroidAOSP,
	// AndroidMobileApplicationManagement, iOSMobileApplicationManagement. Default value: Unknown. Read-only. Supports
	// $select.
	DevicePlatform nullable.Type[string] `json:"devicePlatform,omitempty"`

	// Certificate extensions that further define the purpose of the public key contained in a certificate. Data is
	// formatted as a comma-separated list of object identifiers (OID). For example a possible value is '1.3.6.1.5.5.7.3.2'.
	// Read-only. Nullable.
	ExtendedKeyUsages *[]string `json:"extendedKeyUsages,omitempty"`

	// The globally unique identifier of the certification authority that issued the leaf certificate. Read-only.
	IssuerId nullable.Type[string] `json:"issuerId,omitempty"`

	// The name of the certification authority that issued the leaf certificate. Read-only.
	IssuerName nullable.Type[string] `json:"issuerName,omitempty"`

	// Certificate extensions that define the purpose of the public key contained in a certificate. For example possible
	// values are 'Key Encipherment' and 'Digital Signature'. Read-only. Nullable.
	KeyUsages *[]string `json:"keyUsages,omitempty"`

	// The Online Certificate Status Protocol (OCSP) responder URI that can be used to determine certificate status.
	// Read-only.
	OcspResponderUri nullable.Type[string] `json:"ocspResponderUri,omitempty"`

	// The date and time a certificate was revoked. If the certificate was not revoked, this will be null. The
	// DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For
	// example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
	RevocationDateTime nullable.Type[string] `json:"revocationDateTime,omitempty"`

	// The serial number used to uniquely identify a certificate with its issuing certification authority. Read-only.
	// Supports $select.
	SerialNumber *string `json:"serialNumber,omitempty"`

	// The subject name of the certificate. The subject is the target or intended beneficiary of the security being
	// provided, such as a user or device. Read-only. Supports $select and $orderby.
	SubjectName nullable.Type[string] `json:"subjectName,omitempty"`

	// Secure Hash Algorithm 1 digest of the certificate that can be used to identify it. Read-only. Supports $select.
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`

	// The unique identifier of the user for which the certificate was created. Null for userless devices. This is an Intune
	// user ID. Nullable. Read-only. Supports $select.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// User principal name of the user for which the certificate was created. Null for userless devices. Nullable.
	// Read-only. Supports $select.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The end date time of the validity period of a certificate. Certificates cannot be used after this date time as they
	// are longer valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in
	// UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
	// Supports $orderby.
	ValidityEndDateTime nullable.Type[string] `json:"validityEndDateTime,omitempty"`

	// The start date time of the validity period of a certificate. Certificates cannot be used before this date time as
	// they are not yet valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable.
	// Read-only. Supports $orderby.
	ValidityStartDateTime nullable.Type[string] `json:"validityStartDateTime,omitempty"`

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

func (s CloudCertificationAuthorityLeafCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudCertificationAuthorityLeafCertificate{}

func (s CloudCertificationAuthorityLeafCertificate) MarshalJSON() ([]byte, error) {
	type wrapper CloudCertificationAuthorityLeafCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudCertificationAuthorityLeafCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudCertificationAuthorityLeafCertificate: %+v", err)
	}

	delete(decoded, "certificationAuthorityIssuerUri")
	delete(decoded, "crlDistributionPointUrl")
	delete(decoded, "deviceId")
	delete(decoded, "deviceName")
	delete(decoded, "devicePlatform")
	delete(decoded, "extendedKeyUsages")
	delete(decoded, "issuerId")
	delete(decoded, "issuerName")
	delete(decoded, "keyUsages")
	delete(decoded, "ocspResponderUri")
	delete(decoded, "revocationDateTime")
	delete(decoded, "serialNumber")
	delete(decoded, "subjectName")
	delete(decoded, "thumbprint")
	delete(decoded, "userId")
	delete(decoded, "userPrincipalName")
	delete(decoded, "validityEndDateTime")
	delete(decoded, "validityStartDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudCertificationAuthorityLeafCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudCertificationAuthorityLeafCertificate: %+v", err)
	}

	return encoded, nil
}
