package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustedCertificateAuthorityBase interface {
	Entity
	DirectoryObject
	TrustedCertificateAuthorityBase() BaseTrustedCertificateAuthorityBaseImpl
}

var _ TrustedCertificateAuthorityBase = BaseTrustedCertificateAuthorityBaseImpl{}

type BaseTrustedCertificateAuthorityBaseImpl struct {
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

func (s BaseTrustedCertificateAuthorityBaseImpl) TrustedCertificateAuthorityBase() BaseTrustedCertificateAuthorityBaseImpl {
	return s
}

func (s BaseTrustedCertificateAuthorityBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseTrustedCertificateAuthorityBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ TrustedCertificateAuthorityBase = RawTrustedCertificateAuthorityBaseImpl{}

// RawTrustedCertificateAuthorityBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTrustedCertificateAuthorityBaseImpl struct {
	trustedCertificateAuthorityBase BaseTrustedCertificateAuthorityBaseImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawTrustedCertificateAuthorityBaseImpl) TrustedCertificateAuthorityBase() BaseTrustedCertificateAuthorityBaseImpl {
	return s.trustedCertificateAuthorityBase
}

func (s RawTrustedCertificateAuthorityBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.trustedCertificateAuthorityBase.DirectoryObject()
}

func (s RawTrustedCertificateAuthorityBaseImpl) Entity() BaseEntityImpl {
	return s.trustedCertificateAuthorityBase.Entity()
}

var _ json.Marshaler = BaseTrustedCertificateAuthorityBaseImpl{}

func (s BaseTrustedCertificateAuthorityBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseTrustedCertificateAuthorityBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseTrustedCertificateAuthorityBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseTrustedCertificateAuthorityBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trustedCertificateAuthorityBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseTrustedCertificateAuthorityBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalTrustedCertificateAuthorityBaseImplementation(input []byte) (TrustedCertificateAuthorityBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TrustedCertificateAuthorityBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.mutualTlsOauthConfiguration") {
		var out MutualTlsOauthConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MutualTlsOauthConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseTrustedCertificateAuthorityBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTrustedCertificateAuthorityBaseImpl: %+v", err)
	}

	return RawTrustedCertificateAuthorityBaseImpl{
		trustedCertificateAuthorityBase: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
