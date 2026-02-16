package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TermsOfUseContainer{}

type TermsOfUseContainer struct {
	// Represents the current status of a user's response to a company's customizable terms of use agreement.
	AgreementAcceptances *[]AgreementAcceptance `json:"agreementAcceptances,omitempty"`

	// Represents a tenant's customizable terms of use agreement that's created and managed with Microsoft Entra ID
	// Governance.
	Agreements *[]Agreement `json:"agreements,omitempty"`

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

func (s TermsOfUseContainer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TermsOfUseContainer{}

func (s TermsOfUseContainer) MarshalJSON() ([]byte, error) {
	type wrapper TermsOfUseContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TermsOfUseContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TermsOfUseContainer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.termsOfUseContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TermsOfUseContainer: %+v", err)
	}

	return encoded, nil
}
