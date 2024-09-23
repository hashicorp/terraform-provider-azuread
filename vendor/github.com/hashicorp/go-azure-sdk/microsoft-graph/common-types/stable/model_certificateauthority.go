package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthority struct {
	// Required. The base64 encoded string representing the public certificate.
	Certificate string `json:"certificate"`

	// The URL of the certificate revocation list.
	CertificateRevocationListUrl nullable.Type[string] `json:"certificateRevocationListUrl,omitempty"`

	// The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was
	// created.
	DeltaCertificateRevocationListUrl nullable.Type[string] `json:"deltaCertificateRevocationListUrl,omitempty"`

	// Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate
	// authority.
	IsRootAuthority bool `json:"isRootAuthority"`

	// The issuer of the certificate, calculated from the certificate value. Read-only.
	Issuer *string `json:"issuer,omitempty"`

	// The subject key identifier of the certificate, calculated from the certificate value. Read-only.
	IssuerSki *string `json:"issuerSki,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = CertificateAuthority{}

func (s CertificateAuthority) MarshalJSON() ([]byte, error) {
	type wrapper CertificateAuthority
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateAuthority: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateAuthority: %+v", err)
	}

	delete(decoded, "issuer")
	delete(decoded, "issuerSki")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateAuthority: %+v", err)
	}

	return encoded, nil
}
