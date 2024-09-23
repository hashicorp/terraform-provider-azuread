package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySource = SocialIdentitySource{}

type SocialIdentitySource struct {
	DisplayName              nullable.Type[string]     `json:"displayName,omitempty"`
	SocialIdentitySourceType *SocialIdentitySourceType `json:"socialIdentitySourceType,omitempty"`

	// Fields inherited from IdentitySource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SocialIdentitySource) IdentitySource() BaseIdentitySourceImpl {
	return BaseIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SocialIdentitySource{}

func (s SocialIdentitySource) MarshalJSON() ([]byte, error) {
	type wrapper SocialIdentitySource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SocialIdentitySource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SocialIdentitySource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.socialIdentitySource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SocialIdentitySource: %+v", err)
	}

	return encoded, nil
}
