package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySource = ExternalDomainFederation{}

type ExternalDomainFederation struct {
	// The name of the identity source, typically also the domain name. Read only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The domain name. Read only.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// The issuerURI of the incoming federation. Read only.
	IssuerUri nullable.Type[string] `json:"issuerUri,omitempty"`

	// Fields inherited from IdentitySource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ExternalDomainFederation) IdentitySource() BaseIdentitySourceImpl {
	return BaseIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalDomainFederation{}

func (s ExternalDomainFederation) MarshalJSON() ([]byte, error) {
	type wrapper ExternalDomainFederation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalDomainFederation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalDomainFederation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalDomainFederation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalDomainFederation: %+v", err)
	}

	return encoded, nil
}
