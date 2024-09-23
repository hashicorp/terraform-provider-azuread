package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessContext = WhatIfAuthenticationContext{}

type WhatIfAuthenticationContext struct {
	AuthenticationContext nullable.Type[string] `json:"authenticationContext,omitempty"`

	// Fields inherited from ConditionalAccessContext

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WhatIfAuthenticationContext) ConditionalAccessContext() BaseConditionalAccessContextImpl {
	return BaseConditionalAccessContextImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WhatIfAuthenticationContext{}

func (s WhatIfAuthenticationContext) MarshalJSON() ([]byte, error) {
	type wrapper WhatIfAuthenticationContext
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WhatIfAuthenticationContext: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WhatIfAuthenticationContext: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.whatIfAuthenticationContext"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WhatIfAuthenticationContext: %+v", err)
	}

	return encoded, nil
}
