package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VerifiableCredentialRequirementStatus = VerifiableCredentialRetrieved{}

type VerifiableCredentialRetrieved struct {
	// The specific date and time that the presentation request will expire and a new one will need to be generated.
	ExpiryDateTime *string `json:"expiryDateTime,omitempty"`

	// Fields inherited from VerifiableCredentialRequirementStatus

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s VerifiableCredentialRetrieved) VerifiableCredentialRequirementStatus() BaseVerifiableCredentialRequirementStatusImpl {
	return BaseVerifiableCredentialRequirementStatusImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VerifiableCredentialRetrieved{}

func (s VerifiableCredentialRetrieved) MarshalJSON() ([]byte, error) {
	type wrapper VerifiableCredentialRetrieved
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VerifiableCredentialRetrieved: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VerifiableCredentialRetrieved: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.verifiableCredentialRetrieved"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VerifiableCredentialRetrieved: %+v", err)
	}

	return encoded, nil
}
