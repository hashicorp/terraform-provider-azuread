package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AndroidForWorkEnrollmentProfile{}

type AndroidForWorkEnrollmentProfile struct {
	// Tenant GUID the enrollment profile belongs to.
	AccountId nullable.Type[string] `json:"accountId,omitempty"`

	// Date time the enrollment profile was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Description for the enrollment profile.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for the enrollment profile.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Total number of Android devices that have enrolled using this enrollment profile.
	EnrolledDeviceCount *int64 `json:"enrolledDeviceCount,omitempty"`

	// Date time the enrollment profile was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// String used to generate a QR code for the token.
	QrCodeContent nullable.Type[string] `json:"qrCodeContent,omitempty"`

	// String used to generate a QR code for the token.
	QrCodeImage *MimeContent `json:"qrCodeImage,omitempty"`

	// Date time the most recently created token will expire.
	TokenExpirationDateTime *string `json:"tokenExpirationDateTime,omitempty"`

	// Value of the most recently created token for this enrollment profile.
	TokenValue nullable.Type[string] `json:"tokenValue,omitempty"`

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

func (s AndroidForWorkEnrollmentProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidForWorkEnrollmentProfile{}

func (s AndroidForWorkEnrollmentProfile) MarshalJSON() ([]byte, error) {
	type wrapper AndroidForWorkEnrollmentProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidForWorkEnrollmentProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidForWorkEnrollmentProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidForWorkEnrollmentProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidForWorkEnrollmentProfile: %+v", err)
	}

	return encoded, nil
}
