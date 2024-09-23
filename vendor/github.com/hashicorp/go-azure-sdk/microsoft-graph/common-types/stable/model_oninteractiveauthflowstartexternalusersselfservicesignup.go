package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnInteractiveAuthFlowStartHandler = OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp{}

type OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp struct {
	// Optional. Specifies whether the authentication flow includes an option to sign up (create account) and sign in.
	// Default value is false meaning only sign in is enabled.
	IsSignUpAllowed *bool `json:"isSignUpAllowed,omitempty"`

	// Fields inherited from OnInteractiveAuthFlowStartHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) OnInteractiveAuthFlowStartHandler() BaseOnInteractiveAuthFlowStartHandlerImpl {
	return BaseOnInteractiveAuthFlowStartHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp{}

func (s OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) MarshalJSON() ([]byte, error) {
	type wrapper OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onInteractiveAuthFlowStartExternalUsersSelfServiceSignUp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp: %+v", err)
	}

	return encoded, nil
}
