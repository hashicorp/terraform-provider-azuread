package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RiskDetection{}

type RiskDetection struct {
	// Indicates the activity type the detected risk is linked to. Possible values are: signin, user, unknownFutureValue.
	Activity *ActivityType `json:"activity,omitempty"`

	// Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this:
	// 2014-01-01T00:00:00Z
	ActivityDateTime nullable.Type[string] `json:"activityDateTime,omitempty"`

	// Additional information associated with the risk detection in JSON format. For example,
	// '[{/'Key/':/'userAgent/',/'Value/':/'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)
	// Chrome/68.0.3440.106 Safari/537.36/'}]'. Possible keys in the additionalInfo JSON string are: userAgent, alertUrl,
	// relatedEventTimeInUtc, relatedUserAgent, deviceInformation, relatedLocation, requestId, correlationId,
	// lastActivityTimeInUtc, malwareName, clientLocation, clientIp, riskReasons. For more information about riskReasons and
	// possible values, see riskReasons values.
	AdditionalInfo nullable.Type[string] `json:"additionalInfo,omitempty"`

	// Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not
	// associated with a sign-in.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: 2014-01-01T00:00:00Z
	DetectedDateTime nullable.Type[string] `json:"detectedDateTime,omitempty"`

	// Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline,
	// unknownFutureValue.
	DetectionTimingType *RiskDetectionTimingType `json:"detectionTimingType,omitempty"`

	// Provides the IP address of the client from where the risk occurred.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this:
	// 2014-01-01T00:00:00Z
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// Location of the sign-in.
	Location *SignInLocation `json:"location,omitempty"`

	// Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not
	// associated with a sign-in.
	RequestId nullable.Type[string] `json:"requestId,omitempty"`

	// Details of the detected risk. The possible values are: none, adminGeneratedTemporaryPassword,
	// userChangedPasswordOnPremises, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset,
	// adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser,
	// adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue,
	// m365DAdminDismissedDetection. Use the Prefer: include - unknown -enum-members request header to get the following
	// value(s) in this evolvable enum: m365DAdminDismissedDetection.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The type of risk event detected. The possible values are adminConfirmedUserCompromised, anomalousToken,
	// anomalousUserActivity, anonymizedIPAddress, generic, impossibleTravel, investigationsThreatIntelligence,
	// suspiciousSendingPatterns, leakedCredentials, maliciousIPAddress,malwareInfectedIPAddress,
	// mcasSuspiciousInboxManipulationRules, newCountry, passwordSpray,riskyIPAddress, suspiciousAPITraffic,
	// suspiciousBrowser,suspiciousInboxForwarding, suspiciousIPAddress, tokenIssuerAnomaly, unfamiliarFeatures,
	// unlikelyTravel. If the risk detection is a premium detection, will show generic. For more information about each
	// value, see Risk types and detection.
	RiskEventType nullable.Type[string] `json:"riskEventType,omitempty"`

	// Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
	RiskLevel *RiskLevel `json:"riskLevel,omitempty"`

	// The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed,
	// atRisk, confirmedCompromised, unknownFutureValue.
	RiskState *RiskState `json:"riskState,omitempty"`

	// Source of the risk detection. For example, activeDirectory.
	Source nullable.Type[string] `json:"source,omitempty"`

	// Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices,
	// UnknownFutureValue.
	TokenIssuerType *TokenIssuerType `json:"tokenIssuerType,omitempty"`

	// The user principal name (UPN) of the user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// Unique ID of the user.
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
