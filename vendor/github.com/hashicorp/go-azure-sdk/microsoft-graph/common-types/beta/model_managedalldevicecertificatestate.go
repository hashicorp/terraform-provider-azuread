package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedAllDeviceCertificateState{}

type ManagedAllDeviceCertificateState struct {
	// Certificate expiry date
	CertificateExpirationDateTime *string `json:"certificateExpirationDateTime,omitempty"`

	// Enhanced Key Usage
	CertificateExtendedKeyUsages nullable.Type[string] `json:"certificateExtendedKeyUsages,omitempty"`

	// Issuance date
	CertificateIssuanceDateTime *string `json:"certificateIssuanceDateTime,omitempty"`

	// Issuer
	CertificateIssuerName nullable.Type[string] `json:"certificateIssuerName,omitempty"`

	// Key Usage
	CertificateKeyUsages nullable.Type[int64] `json:"certificateKeyUsages,omitempty"`

	// Certificate Revocation Status.
	CertificateRevokeStatus *CertificateRevocationStatus `json:"certificateRevokeStatus,omitempty"`

	// The time the revoke status was last changed
	CertificateRevokeStatusLastChangeDateTime *string `json:"certificateRevokeStatusLastChangeDateTime,omitempty"`

	// Serial number
	CertificateSerialNumber nullable.Type[string] `json:"certificateSerialNumber,omitempty"`

	// Certificate subject name
	CertificateSubjectName nullable.Type[string] `json:"certificateSubjectName,omitempty"`

	// Thumbprint
	CertificateThumbprint nullable.Type[string] `json:"certificateThumbprint,omitempty"`

	// Device display name
	ManagedDeviceDisplayName nullable.Type[string] `json:"managedDeviceDisplayName,omitempty"`

	// User principal name
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

func (s ManagedAllDeviceCertificateState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedAllDeviceCertificateState{}

func (s ManagedAllDeviceCertificateState) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAllDeviceCertificateState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAllDeviceCertificateState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAllDeviceCertificateState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAllDeviceCertificateState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAllDeviceCertificateState: %+v", err)
	}

	return encoded, nil
}
