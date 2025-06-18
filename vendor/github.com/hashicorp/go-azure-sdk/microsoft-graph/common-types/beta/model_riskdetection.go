package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RiskDetection{}

type RiskDetection struct {
	// Indicates the activity type the detected risk is linked to. The possible values are signin, user, unknownFutureValue.
	Activity *ActivityType `json:"activity,omitempty"`

	// Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ActivityDateTime nullable.Type[string] `json:"activityDateTime,omitempty"`

	// Additional information associated with the risk detection in JSON format.
	AdditionalInfo nullable.Type[string] `json:"additionalInfo,omitempty"`

	// Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not
	// associated with a sign-in.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DetectedDateTime nullable.Type[string] `json:"detectedDateTime,omitempty"`

	// Timing of the detected risk (real-time/offline). The possible values are notDefined, realtime, nearRealtime, offline,
	// unknownFutureValue.
	DetectionTimingType *RiskDetectionTimingType `json:"detectionTimingType,omitempty"`

	// Provides the IP address of the client from where the risk occurred.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Date and time that the risk detection was last updated.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// Location of the sign-in.
	Location *SignInLocation `json:"location,omitempty"`

	MitreTechniqueId nullable.Type[string] `json:"mitreTechniqueId,omitempty"`

	// Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not
	// associated with a sign-in.
	RequestId nullable.Type[string] `json:"requestId,omitempty"`

	// Details of the detected risk. The possible values are: none, adminGeneratedTemporaryPassword,
	// userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe,
	// aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser,
	// adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue,
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection.
	// Use the Prefer: include - unknown -enum-members request header to get the following value(s) in this evolvable enum:
	// adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal , m365DAdminDismissedDetection.
	// Note: Details for this property are only available for Microsoft Entra ID P2 customers. P1 customers will be returned
	// hidden.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The type of risk event detected. The possible values are adminConfirmedUserCompromised, anomalousUserActivity,
	// anomalousToken, anonymizedIPAddress,attackerinTheMiddle,attemptedPRTAccess, generic,
	// investigationsThreatIntelligence, investigationsThreatIntelligenceSigninLinked,leakedCredentials, maliciousIPAddress,
	// maliciousIPAddressValidCredentialsBlockedIP, malwareInfectedIPAddress,
	// mcasImpossibleTravel,mcasFinSuspiciousFileAccess, mcasSuspiciousInboxManipulationRules,nationStateIP, newCountry,
	// passwordSpray, riskyIPAddress, suspiciousAPITraffic, suspiciousBrowser, suspiciousInboxForwarding,
	// suspiciousIPAddress,suspiciousSendingPatterns, tokenIssuerAnomaly, unfamiliarFeatures, unlikelyTravel,
	// userReportedSuspiciousActivity. For more information about each value, see Risk types and detection.
	RiskEventType nullable.Type[string] `json:"riskEventType,omitempty"`

	// Level of the detected risk. The possible values are low, medium, high, hidden, none, unknownFutureValue. Note:
	// Details for this property are only available for Microsoft Entra ID P2 customers. P1 customers will be returned
	// hidden.
	RiskLevel *RiskLevel `json:"riskLevel,omitempty"`

	// The state of a detected risky user or sign-in. The possible values are none, confirmedSafe, remediated, dismissed,
	// atRisk, confirmedCompromised, and unknownFutureValue.
	RiskState *RiskState `json:"riskState,omitempty"`

	// List of risk event types.Note: This property is deprecated. Use riskEventType instead.
	RiskType *RiskEventType `json:"riskType,omitempty"`

	// Source of the risk detection. For example, activeDirectory.
	Source nullable.Type[string] `json:"source,omitempty"`

	// Indicates the type of token issuer for the detected sign-in risk. The possible values are AzureAD,
	// ADFederationServices, and unknownFutureValue.
	TokenIssuerType *TokenIssuerType `json:"tokenIssuerType,omitempty"`

	// Name of the user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Unique ID of the user. The DateTimeOffset type represents date and time information using ISO 8601 format and is
	// always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user principal name (UPN) of the user.
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

func (s RiskDetection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RiskDetection{}

func (s RiskDetection) MarshalJSON() ([]byte, error) {
	type wrapper RiskDetection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RiskDetection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RiskDetection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.riskDetection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RiskDetection: %+v", err)
	}

	return encoded, nil
}
