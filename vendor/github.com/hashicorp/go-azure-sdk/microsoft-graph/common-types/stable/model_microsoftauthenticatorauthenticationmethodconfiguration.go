package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = MicrosoftAuthenticatorAuthenticationMethodConfiguration{}

type MicrosoftAuthenticatorAuthenticationMethodConfiguration struct {
	// A collection of Microsoft Authenticator settings such as application context and location context, and whether they
	// are enabled for all users or specific users only.
	FeatureSettings *MicrosoftAuthenticatorFeatureSettings `json:"featureSettings,omitempty"`

	// A collection of groups that are enabled to use the authentication method. Expanded by default.
	IncludeTargets *[]MicrosoftAuthenticatorAuthenticationMethodTarget `json:"includeTargets,omitempty"`

	IsSoftwareOathEnabled nullable.Type[bool] `json:"isSoftwareOathEnabled,omitempty"`

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

func (s MicrosoftAuthenticatorAuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s MicrosoftAuthenticatorAuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftAuthenticatorAuthenticationMethodConfiguration{}

func (s MicrosoftAuthenticatorAuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftAuthenticatorAuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftAuthenticatorAuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftAuthenticatorAuthenticationMethodConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftAuthenticatorAuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}
