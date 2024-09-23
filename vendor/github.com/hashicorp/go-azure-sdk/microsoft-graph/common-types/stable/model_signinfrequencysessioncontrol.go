package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessSessionControl = SignInFrequencySessionControl{}

type SignInFrequencySessionControl struct {
	// The possible values are primaryAndSecondaryAuthentication, secondaryAuthentication, unknownFutureValue. This property
	// isn't required when using frequencyInterval with the value of timeBased.
	AuthenticationType *SignInFrequencyAuthenticationType `json:"authenticationType,omitempty"`

	// The possible values are timeBased, everyTime, unknownFutureValue. Sign-in frequency of everyTime is available for
	// risky users, risky sign-ins, and Intune device enrollment. For more information, see Require reauthentication every
	// time.
	FrequencyInterval *SignInFrequencyInterval `json:"frequencyInterval,omitempty"`

	// Possible values are: days, hours.
	Type *SigninFrequencyType `json:"type,omitempty"`

	// The number of days or hours.
	Value nullable.Type[int64] `json:"value,omitempty"`

	// Fields inherited from ConditionalAccessSessionControl

	// Specifies whether the session control is enabled.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SignInFrequencySessionControl) ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl {
	return BaseConditionalAccessSessionControlImpl{
		IsEnabled: s.IsEnabled,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SignInFrequencySessionControl{}

func (s SignInFrequencySessionControl) MarshalJSON() ([]byte, error) {
	type wrapper SignInFrequencySessionControl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SignInFrequencySessionControl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SignInFrequencySessionControl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.signInFrequencySessionControl"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SignInFrequencySessionControl: %+v", err)
	}

	return encoded, nil
}
