package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = QrCodePinAuthenticationMethodConfiguration{}

type QrCodePinAuthenticationMethodConfiguration struct {
	// A collection of groups that are enabled to use the authentication method.
	IncludeTargets *[]AuthenticationMethodTarget `json:"includeTargets,omitempty"`

	// A memorized alphanumeric secret code. Minimum length is 8 as per NIST 800-63B and can't be longer than 20 digits.
	PinLength nullable.Type[int64] `json:"pinLength,omitempty"`

	// The maximum value is 395 days and the default value is 365 days.
	StandardQRCodeLifetimeInDays nullable.Type[int64] `json:"standardQRCodeLifetimeInDays,omitempty"`

	// Fields inherited from AuthenticationMethodConfiguration

	// Groups of users that are excluded from a policy.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// The state of the policy. Possible values are: enabled, disabled.
	State *AuthenticationMethodState `json:"state,omitempty"`

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

func (s QrCodePinAuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s QrCodePinAuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = QrCodePinAuthenticationMethodConfiguration{}

func (s QrCodePinAuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper QrCodePinAuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling QrCodePinAuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling QrCodePinAuthenticationMethodConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.qrCodePinAuthenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling QrCodePinAuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &QrCodePinAuthenticationMethodConfiguration{}

func (s *QrCodePinAuthenticationMethodConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		PinLength                    nullable.Type[int64]       `json:"pinLength,omitempty"`
		StandardQRCodeLifetimeInDays nullable.Type[int64]       `json:"standardQRCodeLifetimeInDays,omitempty"`
		ExcludeTargets               *[]ExcludeTarget           `json:"excludeTargets,omitempty"`
		State                        *AuthenticationMethodState `json:"state,omitempty"`
		Id                           *string                    `json:"id,omitempty"`
		ODataId                      *string                    `json:"@odata.id,omitempty"`
		ODataType                    *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.PinLength = decoded.PinLength
	s.StandardQRCodeLifetimeInDays = decoded.StandardQRCodeLifetimeInDays
	s.ExcludeTargets = decoded.ExcludeTargets
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling QrCodePinAuthenticationMethodConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["includeTargets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IncludeTargets into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationMethodTarget, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationMethodTargetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IncludeTargets' for 'QrCodePinAuthenticationMethodConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IncludeTargets = &output
	}

	return nil
}
