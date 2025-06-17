package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TrustedCertificateAuthorityBase = MutualTlsOauthConfiguration{}

type MutualTlsOauthConfiguration struct {
	// Friendly name. Supports $filter (eq, in).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	TlsClientAuthParameter *TlsClientRegistrationMetadata `json:"tlsClientAuthParameter,omitempty"`

	// Fields inherited from TrustedCertificateAuthorityBase

	// Multi-value property that represents a list of trusted certificate authorities.
	CertificateAuthorities *[]CertificateAuthority `json:"certificateAuthorities,omitempty"`

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

func (s MutualTlsOauthConfiguration) TrustedCertificateAuthorityBase() BaseTrustedCertificateAuthorityBaseImpl {
	return BaseTrustedCertificateAuthorityBaseImpl{
		CertificateAuthorities: s.CertificateAuthorities,
		DeletedDateTime:        s.DeletedDateTime,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s MutualTlsOauthConfiguration) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s MutualTlsOauthConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MutualTlsOauthConfiguration{}

func (s MutualTlsOauthConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MutualTlsOauthConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MutualTlsOauthConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MutualTlsOauthConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mutualTlsOauthConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MutualTlsOauthConfiguration: %+v", err)
	}

	return encoded, nil
}
