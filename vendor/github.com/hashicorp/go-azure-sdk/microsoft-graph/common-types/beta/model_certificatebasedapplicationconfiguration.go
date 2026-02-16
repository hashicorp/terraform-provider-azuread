package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TrustedCertificateAuthorityAsEntityBase = CertificateBasedApplicationConfiguration{}

type CertificateBasedApplicationConfiguration struct {
	// The description of the trusted certificate authorities.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the trusted certificate authorities.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from TrustedCertificateAuthorityAsEntityBase

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

func (s CertificateBasedApplicationConfiguration) TrustedCertificateAuthorityAsEntityBase() BaseTrustedCertificateAuthorityAsEntityBaseImpl {
	return BaseTrustedCertificateAuthorityAsEntityBaseImpl{
		TrustedCertificateAuthorities: s.TrustedCertificateAuthorities,
		DeletedDateTime:               s.DeletedDateTime,
		Id:                            s.Id,
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
	}
}

func (s CertificateBasedApplicationConfiguration) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s CertificateBasedApplicationConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CertificateBasedApplicationConfiguration{}

func (s CertificateBasedApplicationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CertificateBasedApplicationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateBasedApplicationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateBasedApplicationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.certificateBasedApplicationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateBasedApplicationConfiguration: %+v", err)
	}

	return encoded, nil
}
