package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = CertificateAuthorityDetail{}

type CertificateAuthorityDetail struct {
	// The public key of the certificate authority.
	Certificate *string `json:"certificate,omitempty"`

	// The type of certificate authority. The possible values are: root, intermediate, unknownFutureValue. Supports $filter
	// (eq).
	CertificateAuthorityType *CertificateAuthorityType `json:"certificateAuthorityType,omitempty"`

	// The URL to check if the certificate is revoked.
	CertificateRevocationListUrl nullable.Type[string] `json:"certificateRevocationListUrl,omitempty"`

	// The date and time when the certificate authority was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	DeltaCertificateRevocationListUrl nullable.Type[string] `json:"deltaCertificateRevocationListUrl,omitempty"`

	// The display name of the certificate authority.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time when the certificate authority expires. Supports $filter (eq) and $orderby.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Indicates whether the certificate picker presents the certificate authority to the user to use for authentication.
	// Default value is false. Optional.
	IsIssuerHintEnabled nullable.Type[bool] `json:"isIssuerHintEnabled,omitempty"`

	// The issuer of the certificate authority.
	Issuer nullable.Type[string] `json:"issuer,omitempty"`

	// The subject key identifier of certificate authority.
	IssuerSubjectKeyIdentifier nullable.Type[string] `json:"issuerSubjectKeyIdentifier,omitempty"`

	// The thumbprint of certificate authority certificate. Supports $filter (eq, startswith).
	Thumbprint *string `json:"thumbprint,omitempty"`

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

func (s CertificateAuthorityDetail) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s CertificateAuthorityDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CertificateAuthorityDetail{}

func (s CertificateAuthorityDetail) MarshalJSON() ([]byte, error) {
	type wrapper CertificateAuthorityDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateAuthorityDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateAuthorityDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.certificateAuthorityDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateAuthorityDetail: %+v", err)
	}

	return encoded, nil
}
