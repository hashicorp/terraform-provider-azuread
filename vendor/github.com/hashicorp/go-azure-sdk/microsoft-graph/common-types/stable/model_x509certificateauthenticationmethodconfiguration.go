package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = X509CertificateAuthenticationMethodConfiguration{}

type X509CertificateAuthenticationMethodConfiguration struct {
	// Defines strong authentication configurations. This configuration includes the default authentication mode and the
	// different rules for strong authentication bindings.
	AuthenticationModeConfiguration *X509CertificateAuthenticationModeConfiguration `json:"authenticationModeConfiguration,omitempty"`

	// Defines fields in the X.509 certificate that map to attributes of the Microsoft Entra user object in order to bind
	// the certificate to the user. The priority of the object determines the order in which the binding is carried out. The
	// first binding that matches will be used and the rest ignored.
	CertificateUserBindings *[]X509CertificateUserBinding `json:"certificateUserBindings,omitempty"`

	CrlValidationConfiguration *X509CertificateCRLValidationConfiguration `json:"crlValidationConfiguration,omitempty"`

	// A collection of groups that are enabled to use the authentication method.
	IncludeTargets *[]AuthenticationMethodTarget `json:"includeTargets,omitempty"`

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

func (s X509CertificateAuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s X509CertificateAuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = X509CertificateAuthenticationMethodConfiguration{}

func (s X509CertificateAuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper X509CertificateAuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling X509CertificateAuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling X509CertificateAuthenticationMethodConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.x509CertificateAuthenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling X509CertificateAuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &X509CertificateAuthenticationMethodConfiguration{}

func (s *X509CertificateAuthenticationMethodConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationModeConfiguration *X509CertificateAuthenticationModeConfiguration `json:"authenticationModeConfiguration,omitempty"`
		CertificateUserBindings         *[]X509CertificateUserBinding                   `json:"certificateUserBindings,omitempty"`
		CrlValidationConfiguration      *X509CertificateCRLValidationConfiguration      `json:"crlValidationConfiguration,omitempty"`
		ExcludeTargets                  *[]ExcludeTarget                                `json:"excludeTargets,omitempty"`
		State                           *AuthenticationMethodState                      `json:"state,omitempty"`
		Id                              *string                                         `json:"id,omitempty"`
		ODataId                         *string                                         `json:"@odata.id,omitempty"`
		ODataType                       *string                                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuthenticationModeConfiguration = decoded.AuthenticationModeConfiguration
	s.CertificateUserBindings = decoded.CertificateUserBindings
	s.CrlValidationConfiguration = decoded.CrlValidationConfiguration
	s.ExcludeTargets = decoded.ExcludeTargets
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling X509CertificateAuthenticationMethodConfiguration into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'IncludeTargets' for 'X509CertificateAuthenticationMethodConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IncludeTargets = &output
	}

	return nil
}
