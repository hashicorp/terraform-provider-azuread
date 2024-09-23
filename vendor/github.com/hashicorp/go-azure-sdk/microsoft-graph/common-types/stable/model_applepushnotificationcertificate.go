package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApplePushNotificationCertificate{}

type ApplePushNotificationCertificate struct {
	// Apple Id of the account used to create the MDM push certificate.
	AppleIdentifier nullable.Type[string] `json:"appleIdentifier,omitempty"`

	// Not yet documented
	Certificate nullable.Type[string] `json:"certificate,omitempty"`

	// Certificate serial number. This property is read-only.
	CertificateSerialNumber nullable.Type[string] `json:"certificateSerialNumber,omitempty"`

	// The reason the certificate upload failed.
	CertificateUploadFailureReason nullable.Type[string] `json:"certificateUploadFailureReason,omitempty"`

	// The certificate upload status.
	CertificateUploadStatus nullable.Type[string] `json:"certificateUploadStatus,omitempty"`

	// The expiration date and time for Apple push notification certificate.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// Last modified date and time for Apple push notification certificate.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Topic Id.
	TopicIdentifier nullable.Type[string] `json:"topicIdentifier,omitempty"`

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

func (s ApplePushNotificationCertificate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApplePushNotificationCertificate{}

func (s ApplePushNotificationCertificate) MarshalJSON() ([]byte, error) {
	type wrapper ApplePushNotificationCertificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApplePushNotificationCertificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplePushNotificationCertificate: %+v", err)
	}

	delete(decoded, "certificateSerialNumber")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.applePushNotificationCertificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApplePushNotificationCertificate: %+v", err)
	}

	return encoded, nil
}
