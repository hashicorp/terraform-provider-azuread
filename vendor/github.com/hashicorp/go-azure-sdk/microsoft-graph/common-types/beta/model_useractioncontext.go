package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignInContext = UserActionContext{}

type UserActionContext struct {
	// Represents the user action that the authenticating identity is performing. The possible values are:
	// registerSecurityInformation, registerOrJoinDevices, unknownFutureValue.
	UserAction *UserAction `json:"userAction,omitempty"`

	// Fields inherited from SignInContext

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserActionContext) SignInContext() BaseSignInContextImpl {
	return BaseSignInContextImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserActionContext{}

func (s UserActionContext) MarshalJSON() ([]byte, error) {
	type wrapper UserActionContext
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserActionContext: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserActionContext: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userActionContext"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserActionContext: %+v", err)
	}

	return encoded, nil
}
