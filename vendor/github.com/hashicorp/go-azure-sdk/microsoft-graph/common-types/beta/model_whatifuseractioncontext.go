package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessContext = WhatIfUserActionContext{}

type WhatIfUserActionContext struct {
	UserAction *UserAction `json:"userAction,omitempty"`

	// Fields inherited from ConditionalAccessContext

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WhatIfUserActionContext) ConditionalAccessContext() BaseConditionalAccessContextImpl {
	return BaseConditionalAccessContextImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WhatIfUserActionContext{}

func (s WhatIfUserActionContext) MarshalJSON() ([]byte, error) {
	type wrapper WhatIfUserActionContext
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WhatIfUserActionContext: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WhatIfUserActionContext: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.whatIfUserActionContext"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WhatIfUserActionContext: %+v", err)
	}

	return encoded, nil
}
