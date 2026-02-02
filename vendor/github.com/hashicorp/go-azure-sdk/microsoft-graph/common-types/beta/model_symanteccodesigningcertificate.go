package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SymantecCodeSigningCertificate{}

type SymantecCodeSigningCertificate struct {
	// The Windows Symantec Code-Signing Certificate in the raw data format.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The Cert Expiration Date.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// The Issuer value for the cert.
	Issuer nullable.Type[string] `json:"issuer,omitempty"`

	// The Issuer Name for the cert.
	IssuerName nullable.Type[string] `json:"issuerName,omitempty"`

	// The Password required for .pfx file.
	Password nullable.Type[string] `json:"password,omitempty"`

	Status *CertificateStatus `json:"status,omitempty"`

	// The Subject value for the cert.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// The Subject Name for the cert.
	SubjectName nullable.Type[string] `json:"subjectName,omitempty"`

	// The Type of the CodeSigning Cert as Symantec Cert.
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

func (s SymantecCodeSigningCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SymantecCodeSigningCertificate{}

func (s SymantecCodeSigningCertificate) MarshalJSON() ([]byte, error) {
	type wrapper SymantecCodeSigningCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SymantecCodeSigningCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SymantecCodeSigningCertificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.symantecCodeSigningCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SymantecCodeSigningCertificate: %+v", err)
	}

	return encoded, nil
}
