package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CertificateAuthorityAsEntity{}

type CertificateAuthorityAsEntity struct {
	// The trusted certificate.
	Certificate *string `json:"certificate,omitempty"`

	// Indicates if the certificate is a root authority. In a certificateBasedApplicationConfiguration object, at least one
	// object in the trustedCertificateAuthorities collection must be a root authority.
	IsRootAuthority *bool `json:"isRootAuthority,omitempty"`

	// The issuer of the trusted certificate.
	Issuer *string `json:"issuer,omitempty"`

	// The subject key identifier of the trusted certificate.
	IssuerSubjectKeyIdentifier *string `json:"issuerSubjectKeyIdentifier,omitempty"`

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

func (s CertificateAuthorityAsEntity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CertificateAuthorityAsEntity{}

func (s CertificateAuthorityAsEntity) MarshalJSON() ([]byte, error) {
	type wrapper CertificateAuthorityAsEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateAuthorityAsEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateAuthorityAsEntity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.certificateAuthorityAsEntity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateAuthorityAsEntity: %+v", err)
	}

	return encoded, nil
}
