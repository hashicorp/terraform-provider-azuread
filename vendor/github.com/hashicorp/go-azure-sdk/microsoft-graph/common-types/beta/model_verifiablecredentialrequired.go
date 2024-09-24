package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VerifiableCredentialRequirementStatus = VerifiableCredentialRequired{}

type VerifiableCredentialRequired struct {
	// When the presentation request will expire and a new one will need to be generated.
	ExpiryDateTime *string `json:"expiryDateTime,omitempty"`

	// A URL that launches the digital wallet and starts the presentation process. You can present this URL to the user if
	// they can't scan the QR code.
	Url nullable.Type[string] `json:"url,omitempty"`

	// Fields inherited from VerifiableCredentialRequirementStatus

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s VerifiableCredentialRequired) VerifiableCredentialRequirementStatus() BaseVerifiableCredentialRequirementStatusImpl {
	return BaseVerifiableCredentialRequirementStatusImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VerifiableCredentialRequired{}

func (s VerifiableCredentialRequired) MarshalJSON() ([]byte, error) {
	type wrapper VerifiableCredentialRequired
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VerifiableCredentialRequired: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VerifiableCredentialRequired: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.verifiableCredentialRequired"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VerifiableCredentialRequired: %+v", err)
	}

	return encoded, nil
}
