package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SignIn{}

type SignIn struct {
	// App name displayed in the Microsoft Entra admin center. Supports $filter (eq, startsWith).
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// Unique GUID that represents the app ID in the Microsoft Entra ID. Supports $filter (eq).
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Provides a list of conditional access policies that the corresponding sign-in activity triggers. Apps need more
	// Conditional Access-related privileges to read the details of this property. For more information, see Permissions for
	// viewing applied conditional access (CA) policies in sign-ins.
	AppliedConditionalAccessPolicies *[]AppliedConditionalAccessPolicy `json:"appliedConditionalAccessPolicies,omitempty"`

	// Identifies the client used for the sign-in activity. Modern authentication clients include Browser, modern clients.
	// Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter
	// (eq).
	ClientAppUsed nullable.Type[string] `json:"clientAppUsed,omitempty"`

	// Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and
	// unknownFutureValue. Supports $filter (eq).
	ConditionalAccessStatus *ConditionalAccessStatus `json:"conditionalAccessStatus,omitempty"`

	// The request ID sent from the client when the sign-in is initiated. Used to troubleshoot sign-in activity. Supports
	// $filter (eq).
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z.
	// Supports $orderby, $filter (eq, le, and ge).
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports
	// $filter (eq, startsWith) on browser and operatingSytem properties.
	DeviceDetail *DeviceDetail `json:"deviceDetail,omitempty"`

	// IP address of the client used to sign in. Supports $filter (eq, startsWith).
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Indicates whether a sign-in is interactive.
	IsInteractive nullable.Type[bool] `json:"isInteractive,omitempty"`

	// Provides the city, state, and country code where the sign-in originated. Supports $filter (eq, startsWith) on city,
	// state, and countryOrRegion properties.
	Location *SignInLocation `json:"location,omitempty"`

	// Name of the resource the user signed into. Supports $filter (eq).
	ResourceDisplayName nullable.Type[string] `json:"resourceDisplayName,omitempty"`

	// ID of the resource that the user signed into. Supports $filter (eq).
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// The reason behind a specific state of a risky user, sign-in, or a risk event. The possible values are none,
	// adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset,
	// adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser,
	// adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue,
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection,
	// userChangedPasswordOnPremises, adminDismissedRiskForSignIn, adminConfirmedAccountSafe. Use the Prefer:
	// include-unknown-enum-members request header to get the following value or values in this evolvable enum:
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection,
	// userChangedPasswordOnPremises, adminDismissedRiskForSignIn, adminConfirmedAccountSafe.The value none means that
	// Microsoft Entra risk detection did not flag the user or the sign-in as a risky event so far. Supports $filter (eq).
	// Note: Details for this property are only available for Microsoft Entra ID P2 customers. All other customers are
	// returned hidden.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	RiskEventTypes *[]RiskEventType `json:"riskEventTypes,omitempty"`

	// The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress,
	// maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials,
	// investigationsThreatIntelligence, generic, or unknownFutureValue. Supports $filter (eq, startsWith).
	RiskEventTypesv2 *[]string `json:"riskEventTypes_v2,omitempty"`

	// Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value
	// hidden means the user or sign-in wasn't enabled for Microsoft Entra ID Protection. Supports $filter (eq). Note:
	// Details for this property are only available for Microsoft Entra ID P2 customers. All other customers are returned
	// hidden.
	RiskLevelAggregated *RiskLevel `json:"riskLevelAggregated,omitempty"`

	// Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The
	// value hidden means the user or sign-in wasn't enabled for Microsoft Entra ID Protection. Supports $filter (eq). Note:
	// Details for this property are only available for Microsoft Entra ID P2 customers. All other customers are returned
	// hidden.
	RiskLevelDuringSignIn *RiskLevel `json:"riskLevelDuringSignIn,omitempty"`

	// Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated,
	// dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq).
	RiskState *RiskState `json:"riskState,omitempty"`

	// Sign-in status. Includes the error code and description of the error (if a sign-in failure occurs). Supports $filter
	// (eq) on errorCode property.
	Status *SignInStatus `json:"status,omitempty"`

	// Display name of the user that initiated the sign-in. Supports $filter (eq, startsWith).
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// ID of the user that initiated the sign-in. Supports $filter (eq).
	UserId *string `json:"userId,omitempty"`

	// User principal name of the user that initiated the sign-in. This value is always in lowercase. For guest users whose
	// values in the user object typically contain #EXT# before the domain part, this property stores the value in both
	// lowercase and the 'true' format. For example, while the user object stores AdeleVance_fabrikam.com#EXT#@contoso.com,
	// the sign-in logs store adelevance@fabrikam.com. Supports $filter (eq, startsWith).
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s SignIn) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SignIn{}

func (s SignIn) MarshalJSON() ([]byte, error) {
	type wrapper SignIn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SignIn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SignIn: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.signIn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SignIn: %+v", err)
	}

	return encoded, nil
}
