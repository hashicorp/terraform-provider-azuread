package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = CertificateBasedAuthPki{}

type CertificateBasedAuthPki struct {
	// The collection of certificate authorities contained in this public key infrastructure resource.
	CertificateAuthorities *[]CertificateAuthorityDetail `json:"certificateAuthorities,omitempty"`

	// The name of the object. Maximum length is 256 characters.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the object was created or last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The status of any asynchronous jobs runs on the object which can be upload or delete.
	Status nullable.Type[string] `json:"status,omitempty"`

	// The status details of the upload/deleted operation of PKI (Public Key Infrastructure).
	StatusDetails nullable.Type[string] `json:"statusDetails,omitempty"`

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

func (s CertificateBasedAuthPki) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s CertificateBasedAuthPki) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CertificateBasedAuthPki{}

func (s CertificateBasedAuthPki) MarshalJSON() ([]byte, error) {
	type wrapper CertificateBasedAuthPki
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateBasedAuthPki: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateBasedAuthPki: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.certificateBasedAuthPki"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateBasedAuthPki: %+v", err)
	}

	return encoded, nil
}
