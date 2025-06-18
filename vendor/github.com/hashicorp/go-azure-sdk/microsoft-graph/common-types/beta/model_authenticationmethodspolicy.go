package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationMethodsPolicy{}

type AuthenticationMethodsPolicy struct {
	// Represents the settings for each authentication method. Automatically expanded on GET
	// /policies/authenticationMethodsPolicy.
	AuthenticationMethodConfigurations *[]AuthenticationMethodConfiguration `json:"authenticationMethodConfigurations,omitempty"`

	// A description of the policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the policy.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date and time of the last update to the policy.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	MicrosoftAuthenticatorPlatformSettings *MicrosoftAuthenticatorPlatformSettings `json:"microsoftAuthenticatorPlatformSettings,omitempty"`

	// The state of migration of the authentication methods policy from the legacy multifactor authentication and
	// self-service password reset (SSPR) policies. The possible values are: premigration - means the authentication methods
	// policy is used for authentication only, legacy policies are respected. migrationInProgress - means the authentication
	// methods policy is used for both authentication and SSPR, legacy policies are respected. migrationComplete - means the
	// authentication methods policy is used for authentication and SSPR, legacy policies are ignored. unknownFutureValue -
	// Evolvable enumeration sentinel value. Don't use.
	PolicyMigrationState *AuthenticationMethodsPolicyMigrationState `json:"policyMigrationState,omitempty"`

	// The version of the policy in use.
	PolicyVersion nullable.Type[string] `json:"policyVersion,omitempty"`

	// Days before the user will be asked to reconfirm their method.
	ReconfirmationInDays nullable.Type[int64] `json:"reconfirmationInDays,omitempty"`

	// Enforce registration at sign-in time. This property can be used to remind users to set up targeted authentication
	// methods.
	RegistrationEnforcement *RegistrationEnforcement `json:"registrationEnforcement,omitempty"`

	// Enable users to report unexpected voice call or phone app notification multi-factor authentication prompts as
	// suspicious.
	ReportSuspiciousActivitySettings *ReportSuspiciousActivitySettings `json:"reportSuspiciousActivitySettings,omitempty"`

	// Prompt users with their most-preferred credential for multifactor authentication.
	SystemCredentialPreferences *SystemCredentialPreferences `json:"systemCredentialPreferences,omitempty"`

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

func (s AuthenticationMethodsPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationMethodsPolicy{}

func (s AuthenticationMethodsPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationMethodsPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationMethodsPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationMethodsPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationMethodsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationMethodsPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AuthenticationMethodsPolicy{}

func (s *AuthenticationMethodsPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description                            nullable.Type[string]                      `json:"description,omitempty"`
		DisplayName                            nullable.Type[string]                      `json:"displayName,omitempty"`
		LastModifiedDateTime                   nullable.Type[string]                      `json:"lastModifiedDateTime,omitempty"`
		MicrosoftAuthenticatorPlatformSettings *MicrosoftAuthenticatorPlatformSettings    `json:"microsoftAuthenticatorPlatformSettings,omitempty"`
		PolicyMigrationState                   *AuthenticationMethodsPolicyMigrationState `json:"policyMigrationState,omitempty"`
		PolicyVersion                          nullable.Type[string]                      `json:"policyVersion,omitempty"`
		ReconfirmationInDays                   nullable.Type[int64]                       `json:"reconfirmationInDays,omitempty"`
		RegistrationEnforcement                *RegistrationEnforcement                   `json:"registrationEnforcement,omitempty"`
		ReportSuspiciousActivitySettings       *ReportSuspiciousActivitySettings          `json:"reportSuspiciousActivitySettings,omitempty"`
		SystemCredentialPreferences            *SystemCredentialPreferences               `json:"systemCredentialPreferences,omitempty"`
		Id                                     *string                                    `json:"id,omitempty"`
		ODataId                                *string                                    `json:"@odata.id,omitempty"`
		ODataType                              *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.MicrosoftAuthenticatorPlatformSettings = decoded.MicrosoftAuthenticatorPlatformSettings
	s.PolicyMigrationState = decoded.PolicyMigrationState
	s.PolicyVersion = decoded.PolicyVersion
	s.ReconfirmationInDays = decoded.ReconfirmationInDays
	s.RegistrationEnforcement = decoded.RegistrationEnforcement
	s.ReportSuspiciousActivitySettings = decoded.ReportSuspiciousActivitySettings
	s.SystemCredentialPreferences = decoded.SystemCredentialPreferences
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AuthenticationMethodsPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authenticationMethodConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AuthenticationMethodConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationMethodConfiguration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationMethodConfigurationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AuthenticationMethodConfigurations' for 'AuthenticationMethodsPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AuthenticationMethodConfigurations = &output
	}

	return nil
}
