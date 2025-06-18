package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionBehaviorOnError = FallbackToMicrosoftProviderOnError{}

type FallbackToMicrosoftProviderOnError struct {

	// Fields inherited from CustomExtensionBehaviorOnError

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s FallbackToMicrosoftProviderOnError) CustomExtensionBehaviorOnError() BaseCustomExtensionBehaviorOnErrorImpl {
	return BaseCustomExtensionBehaviorOnErrorImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FallbackToMicrosoftProviderOnError{}

func (s FallbackToMicrosoftProviderOnError) MarshalJSON() ([]byte, error) {
	type wrapper FallbackToMicrosoftProviderOnError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FallbackToMicrosoftProviderOnError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FallbackToMicrosoftProviderOnError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fallbackToMicrosoftProviderOnError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FallbackToMicrosoftProviderOnError: %+v", err)
	}

	return encoded, nil
}
