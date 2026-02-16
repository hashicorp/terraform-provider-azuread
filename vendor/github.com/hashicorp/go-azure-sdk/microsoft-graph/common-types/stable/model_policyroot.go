package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PolicyRoot{}

type PolicyRoot struct {
	// The policy that controls the idle time out for web sessions for applications.
	ActivityBasedTimeoutPolicies *[]ActivityBasedTimeoutPolicy `json:"activityBasedTimeoutPolicies,omitempty"`

	// The policy by which consent requests are created and managed for the entire tenant.
	AdminConsentRequestPolicy *AdminConsentRequestPolicy `json:"adminConsentRequestPolicy,omitempty"`

	// The policies that enforce app management restrictions for specific applications and service principals, overriding
	// the defaultAppManagementPolicy.
	AppManagementPolicies *[]AppManagementPolicy `json:"appManagementPolicies,omitempty"`

	// The policy configuration of the self-service sign-up experience of external users.
	AuthenticationFlowsPolicy *AuthenticationFlowsPolicy `json:"authenticationFlowsPolicy,omitempty"`

	// The authentication methods and the users that are allowed to use them to sign in and perform multifactor
	// authentication (MFA) in Microsoft Entra ID.
	AuthenticationMethodsPolicy *AuthenticationMethodsPolicy `json:"authenticationMethodsPolicy,omitempty"`

	// The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional
	// Access.
	AuthenticationStrengthPolicies *[]AuthenticationStrengthPolicy `json:"authenticationStrengthPolicies,omitempty"`

	// The policy that controls Microsoft Entra authorization settings.
	AuthorizationPolicy *AuthorizationPolicy `json:"authorizationPolicy,omitempty"`

	// The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific
	// application.
	ClaimsMappingPolicies *[]ClaimsMappingPolicy `json:"claimsMappingPolicies,omitempty"`

	// The custom rules that define an access scenario.
	ConditionalAccessPolicies *[]ConditionalAccessPolicy `json:"conditionalAccessPolicies,omitempty"`

	// The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
	CrossTenantAccessPolicy *CrossTenantAccessPolicy `json:"crossTenantAccessPolicy,omitempty"`

	// The tenant-wide policy that enforces app management restrictions for all applications and service principals.
	DefaultAppManagementPolicy *TenantAppManagementPolicy `json:"defaultAppManagementPolicy,omitempty"`

	DeviceRegistrationPolicy *DeviceRegistrationPolicy `json:"deviceRegistrationPolicy,omitempty"`

	// The feature rollout policy associated with a directory object.
	FeatureRolloutPolicies *[]FeatureRolloutPolicy `json:"featureRolloutPolicies,omitempty"`

	// The policy to control Microsoft Entra authentication behavior for federated users.
	HomeRealmDiscoveryPolicies *[]HomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`

	// The policy that represents the security defaults that protect against common attacks.
	IdentitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy `json:"identitySecurityDefaultsEnforcementPolicy,omitempty"`

	// The policy that specifies the conditions under which consent can be granted.
	PermissionGrantPolicies *[]PermissionGrantPolicy `json:"permissionGrantPolicies,omitempty"`

	// Specifies the various policies associated with scopes and roles.
	RoleManagementPolicies *[]UnifiedRoleManagementPolicy `json:"roleManagementPolicies,omitempty"`

	// The assignment of a role management policy to a role definition object.
	RoleManagementPolicyAssignments *[]UnifiedRoleManagementPolicyAssignment `json:"roleManagementPolicyAssignments,omitempty"`

	// The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
	TokenIssuancePolicies *[]TokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`

	// The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft
	// Entra ID.
	TokenLifetimePolicies *[]TokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`

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

func (s PolicyRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PolicyRoot{}

func (s PolicyRoot) MarshalJSON() ([]byte, error) {
	type wrapper PolicyRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicyRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policyRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicyRoot: %+v", err)
	}

	return encoded, nil
}
