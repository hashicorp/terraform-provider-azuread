package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EnterpriseCodeSigningCertificate{}

type EnterpriseCodeSigningCertificate struct {
	// The Windows Enterprise Code-Signing Certificate in the raw data format. Set to null once certificate has been
	// uploaded and other properties have been populated.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The cert expiration date and time (using ISO 8601 format, in UTC time). Uploading a valid cert file through the
	// Intune admin console will automatically populate this value in the HTTP response. Supports: $filter, $select, $top,
	// $OrderBy, $skip. $Search is not supported.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// The issuer value for the cert. This might contain information such as country (C), state or province (S), locality
	// (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file
	// through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter,
	// $select, $top, $OrderBy, $skip. $Search is not supported.
	Issuer nullable.Type[string] `json:"issuer,omitempty"`

	// The issuer name for the cert. This might contain information such as country (C), state or province (S), locality
	// (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file
	// through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter,
	// $select, $top, $OrderBy, $skip. $Search is not supported.
	IssuerName nullable.Type[string] `json:"issuerName,omitempty"`

	Status *CertificateStatus `json:"status,omitempty"`

	// The subject value for the cert. This might contain information such as country (C), state or province (S), locality
	// (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file
	// through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter,
	// $select, $top, $OrderBy, $skip. $Search is not supported.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// The subject name for the cert. This might contain information such as country (C), state or province (S), locality
	// (L), common name of the cert (CN), organization (O), and organizational unit (OU). Uploading a valid cert file
	// through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter,
	// $select, $top, $OrderBy, $skip. $Search is not supported.
	SubjectName nullable.Type[string] `json:"subjectName,omitempty"`

	// The date time of CodeSigning Cert when it is uploaded (using ISO 8601 format, in UTC time). Uploading a valid cert
	// file through the Intune admin console will automatically populate this value in the HTTP response. Supports: $filter,
	// $select, $top, $OrderBy, $skip. $Search is not supported.
	UploadDateTime *string `json:"uploadDateTime,omitempty"`

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

func (s EnterpriseCodeSigningCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnterpriseCodeSigningCertificate{}

func (s EnterpriseCodeSigningCertificate) MarshalJSON() ([]byte, error) {
	type wrapper EnterpriseCodeSigningCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnterpriseCodeSigningCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnterpriseCodeSigningCertificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enterpriseCodeSigningCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnterpriseCodeSigningCertificate: %+v", err)
	}

	return encoded, nil
}
