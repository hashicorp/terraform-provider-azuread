package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystemIdentitySource = GsuiteSource{}

type GsuiteSource struct {
	// Domain name
	Domain nullable.Type[string] `json:"domain,omitempty"`

	// Fields inherited from AuthorizationSystemIdentitySource

	// Type of identity provider. Read-only.
	IdentityProviderType nullable.Type[string] `json:"identityProviderType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s GsuiteSource) AuthorizationSystemIdentitySource() BaseAuthorizationSystemIdentitySourceImpl {
	return BaseAuthorizationSystemIdentitySourceImpl{
		IdentityProviderType: s.IdentityProviderType,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

var _ json.Marshaler = GsuiteSource{}

func (s GsuiteSource) MarshalJSON() ([]byte, error) {
	type wrapper GsuiteSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GsuiteSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GsuiteSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gsuiteSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GsuiteSource: %+v", err)
	}

	return encoded, nil
}
