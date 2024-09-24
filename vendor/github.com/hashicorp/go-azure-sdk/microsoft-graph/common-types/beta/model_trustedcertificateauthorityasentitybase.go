package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustedCertificateAuthorityAsEntityBase interface {
	Entity
	DirectoryObject
	TrustedCertificateAuthorityAsEntityBase() BaseTrustedCertificateAuthorityAsEntityBaseImpl
}

var _ TrustedCertificateAuthorityAsEntityBase = BaseTrustedCertificateAuthorityAsEntityBaseImpl{}

type BaseTrustedCertificateAuthorityAsEntityBaseImpl struct {
	// Collection of trusted certificate authorities.
	TrustedCertificateAuthorities *[]CertificateAuthorityAsEntity `json:"trustedCertificateAuthorities,omitempty"`

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

func (s BaseTrustedCertificateAuthorityAsEntityBaseImpl) TrustedCertificateAuthorityAsEntityBase() BaseTrustedCertificateAuthorityAsEntityBaseImpl {
	return s
}

func (s BaseTrustedCertificateAuthorityAsEntityBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseTrustedCertificateAuthorityAsEntityBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ TrustedCertificateAuthorityAsEntityBase = RawTrustedCertificateAuthorityAsEntityBaseImpl{}

// RawTrustedCertificateAuthorityAsEntityBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTrustedCertificateAuthorityAsEntityBaseImpl struct {
	trustedCertificateAuthorityAsEntityBase BaseTrustedCertificateAuthorityAsEntityBaseImpl
	Type                                    string
	Values                                  map[string]interface{}
}

func (s RawTrustedCertificateAuthorityAsEntityBaseImpl) TrustedCertificateAuthorityAsEntityBase() BaseTrustedCertificateAuthorityAsEntityBaseImpl {
	return s.trustedCertificateAuthorityAsEntityBase
}

func (s RawTrustedCertificateAuthorityAsEntityBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.trustedCertificateAuthorityAsEntityBase.DirectoryObject()
}

func (s RawTrustedCertificateAuthorityAsEntityBaseImpl) Entity() BaseEntityImpl {
	return s.trustedCertificateAuthorityAsEntityBase.Entity()
}

var _ json.Marshaler = BaseTrustedCertificateAuthorityAsEntityBaseImpl{}

func (s BaseTrustedCertificateAuthorityAsEntityBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseTrustedCertificateAuthorityAsEntityBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseTrustedCertificateAuthorityAsEntityBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseTrustedCertificateAuthorityAsEntityBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trustedCertificateAuthorityAsEntityBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseTrustedCertificateAuthorityAsEntityBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalTrustedCertificateAuthorityAsEntityBaseImplementation(input []byte) (TrustedCertificateAuthorityAsEntityBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TrustedCertificateAuthorityAsEntityBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.certificateBasedApplicationConfiguration") {
		var out CertificateBasedApplicationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CertificateBasedApplicationConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseTrustedCertificateAuthorityAsEntityBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTrustedCertificateAuthorityAsEntityBaseImpl: %+v", err)
	}

	return RawTrustedCertificateAuthorityAsEntityBaseImpl{
		trustedCertificateAuthorityAsEntityBase: parent,
		Type:                                    value,
		Values:                                  temp,
	}, nil

}
