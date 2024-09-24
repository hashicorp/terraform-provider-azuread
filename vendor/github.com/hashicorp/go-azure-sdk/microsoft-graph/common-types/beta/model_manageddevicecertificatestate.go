package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceCertificateState{}

type ManagedDeviceCertificateState struct {
	// Extended key usage
	CertificateEnhancedKeyUsage nullable.Type[string] `json:"certificateEnhancedKeyUsage,omitempty"`

	// Error code
	CertificateErrorCode *int64 `json:"certificateErrorCode,omitempty"`

	// Certificate expiry date
	CertificateExpirationDateTime *string `json:"certificateExpirationDateTime,omitempty"`

	// Issuance date
	CertificateIssuanceDateTime *string `json:"certificateIssuanceDateTime,omitempty"`

	// Certificate Issuance State Options.
	CertificateIssuanceState *CertificateIssuanceStates `json:"certificateIssuanceState,omitempty"`

	// Issuer
	CertificateIssuer nullable.Type[string] `json:"certificateIssuer,omitempty"`

	// Key length
	CertificateKeyLength *int64 `json:"certificateKeyLength,omitempty"`

	// Key Storage Provider (KSP) Import Options.
	CertificateKeyStorageProvider *KeyStorageProviderOption `json:"certificateKeyStorageProvider,omitempty"`

	// Key Usage Options.
	CertificateKeyUsage *KeyUsages `json:"certificateKeyUsage,omitempty"`

	// Last certificate issuance state change
	CertificateLastIssuanceStateChangedDateTime *string `json:"certificateLastIssuanceStateChangedDateTime,omitempty"`

	// Certificate profile display name
	CertificateProfileDisplayName nullable.Type[string] `json:"certificateProfileDisplayName,omitempty"`

	// Certificate Revocation Status.
	CertificateRevokeStatus *CertificateRevocationStatus `json:"certificateRevokeStatus,omitempty"`

	// Serial number
	CertificateSerialNumber nullable.Type[string] `json:"certificateSerialNumber,omitempty"`

	// Subject Alternative Name Options.
	CertificateSubjectAlternativeNameFormat *SubjectAlternativeNameType `json:"certificateSubjectAlternativeNameFormat,omitempty"`

	// Subject alternative name format string for custom formats
	CertificateSubjectAlternativeNameFormatString nullable.Type[string] `json:"certificateSubjectAlternativeNameFormatString,omitempty"`

	// Subject Name Format Options.
	CertificateSubjectNameFormat *SubjectNameFormat `json:"certificateSubjectNameFormat,omitempty"`

	// Subject name format string for custom subject name formats
	CertificateSubjectNameFormatString nullable.Type[string] `json:"certificateSubjectNameFormatString,omitempty"`

	// Thumbprint
	CertificateThumbprint nullable.Type[string] `json:"certificateThumbprint,omitempty"`

	// Validity period
	CertificateValidityPeriod *int64 `json:"certificateValidityPeriod,omitempty"`

	// Certificate Validity Period Options.
	CertificateValidityPeriodUnits *CertificateValidityPeriodScale `json:"certificateValidityPeriodUnits,omitempty"`

	// Device display name
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// Supported platform types.
	DevicePlatform *DevicePlatformType `json:"devicePlatform,omitempty"`

	// Last certificate issuance state change
	LastCertificateStateChangeDateTime *string `json:"lastCertificateStateChangeDateTime,omitempty"`

	// User display name
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

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

func (s ManagedDeviceCertificateState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceCertificateState{}

func (s ManagedDeviceCertificateState) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceCertificateState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceCertificateState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceCertificateState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedDeviceCertificateState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceCertificateState: %+v", err)
	}

	return encoded, nil
}
