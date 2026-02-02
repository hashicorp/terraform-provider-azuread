package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = QrPin{}

type QrPin struct {
	// PIN of the user. It is between 8-20 digits as configured in the QR code authentication method policy. The code is
	// temporary when issued by admin but permanent after the user changes it at the first login attempt. This PIN can be
	// reset by the admin but not the user.
	Code *string `json:"code,omitempty"`

	// The date and time when the PIN was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Defaults to true for a temporary PIN.
	ForceChangePinNextSignIn nullable.Type[bool] `json:"forceChangePinNextSignIn,omitempty"`

	// The date and time when the PIN was updated.
	UpdatedDateTime nullable.Type[string] `json:"updatedDateTime,omitempty"`

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

func (s QrPin) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = QrPin{}

func (s QrPin) MarshalJSON() ([]byte, error) {
	type wrapper QrPin
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling QrPin: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling QrPin: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.qrPin"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling QrPin: %+v", err)
	}

	return encoded, nil
}
