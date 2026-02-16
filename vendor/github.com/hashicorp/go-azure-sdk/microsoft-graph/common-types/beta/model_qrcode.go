package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = QrCode{}

type QrCode struct {
	// The date and time when the QR code was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Temporary QR code lifetime is between 1-12 hours. Standard QR code lifetime is in days and max. is 395 days (13
	// months) and default value is 365 days (12 months).
	ExpireDateTime nullable.Type[string] `json:"expireDateTime,omitempty"`

	// The QR code image's raw data that is returned when a standard or temporary QR code is created.
	Image *QrCodeImageDetails `json:"image,omitempty"`

	// The date and time when the QR code was last used for a successful sign-in.
	LastUsedDateTime nullable.Type[string] `json:"lastUsedDateTime,omitempty"`

	// The date and time when the QR code becomes active and available to use.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

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

func (s QrCode) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = QrCode{}

func (s QrCode) MarshalJSON() ([]byte, error) {
	type wrapper QrCode
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling QrCode: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling QrCode: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.qrCode"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling QrCode: %+v", err)
	}

	return encoded, nil
}
