package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserPFXCertificate{}

type UserPFXCertificate struct {
	// Date/time when this PFX certificate was imported.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Encrypted PFX blob.
	EncryptedPfxBlob nullable.Type[string] `json:"encryptedPfxBlob,omitempty"`

	// Encrypted PFX password.
	EncryptedPfxPassword nullable.Type[string] `json:"encryptedPfxPassword,omitempty"`

	// Certificate's validity expiration date/time.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Supported values for the intended purpose of a user PFX certificate.
	IntendedPurpose *UserPfxIntendedPurpose `json:"intendedPurpose,omitempty"`

	// Name of the key (within the provider) used to encrypt the blob.
	KeyName nullable.Type[string] `json:"keyName,omitempty"`

	// Date/time when this PFX certificate was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Supported values for the padding scheme used by encryption provider.
	PaddingScheme *UserPfxPaddingScheme `json:"paddingScheme,omitempty"`

	// Crypto provider used to encrypt this blob.
	ProviderName nullable.Type[string] `json:"providerName,omitempty"`

	// Certificate's validity start date/time.
	StartDateTime *string `json:"startDateTime,omitempty"`

	// SHA-1 thumbprint of the PFX certificate.
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`

	// User Principal Name of the PFX certificate.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s UserPFXCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserPFXCertificate{}

func (s UserPFXCertificate) MarshalJSON() ([]byte, error) {
	type wrapper UserPFXCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserPFXCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserPFXCertificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userPFXCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserPFXCertificate: %+v", err)
	}

	return encoded, nil
}
