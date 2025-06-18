package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ServicePrincipalRiskDetection{}

type ServicePrincipalRiskDetection struct {
	// Indicates the activity type the detected risk is linked to. The possible values are: signin, servicePrincipal. Use
	// the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum:
	// servicePrincipal.
	Activity *ActivityType `json:"activity,omitempty"`

	// Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ActivityDateTime nullable.Type[string] `json:"activityDateTime,omitempty"`

	// Additional information associated with the risk detection. This string value is represented as a JSON object with the
	// quotations escaped.
	AdditionalInfo nullable.Type[string] `json:"additionalInfo,omitempty"`

	// The unique identifier for the associated application.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk
	// detection is not associated with a sign-in activity.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	DetectedDateTime nullable.Type[string] `json:"detectedDateTime,omitempty"`

	// Timing of the detected risk , whether real-time or offline. The possible values are: notDefined, realtime,
	// nearRealtime, offline, unknownFutureValue.
	DetectionTimingType *RiskDetectionTimingType `json:"detectionTimingType,omitempty"`

	// Provides the IP address of the client from where the risk occurred.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The unique identifier for the key credential associated with the risk detection.
	KeyIds *[]string `json:"keyIds,omitempty"`

	// Date and time when the risk detection was last updated.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// Location from where the sign-in was initiated.
	Location *SignInLocation `json:"location,omitempty"`

	// Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk
	// detection is not associated with a sign-in activity. Supports $filter (eq).
	RequestId nullable.Type[string] `json:"requestId,omitempty"`

	// Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium
	// customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden,
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Use the Prefer:
	// include-unknown-enum-members request header to get the following value(s) in this evolvable enum:
	// adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The type of risk event detected. The possible values are: investigationsThreatIntelligence, generic,
	// adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, anomalousServicePrincipalActivity,
	// maliciousApplication, suspiciousApplication.
	RiskEventType nullable.Type[string] `json:"riskEventType,omitempty"`

	// Level of the detected risk. Note: Details for this property are only available for Workload Identities Premium
	// customers. Events in tenants without this license will be returned hidden. The possible values are: low, medium,
	// high, hidden, none.
	RiskLevel *RiskLevel `json:"riskLevel,omitempty"`

	// The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed,
	// atRisk, confirmedCompromised.
	RiskState *RiskState `json:"riskState,omitempty"`

	// The display name for the service principal.
	ServicePrincipalDisplayName nullable.Type[string] `json:"servicePrincipalDisplayName,omitempty"`

	// The unique identifier for the service principal. Supports $filter (eq).
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Source of the risk detection. For example, identityProtection.
	Source nullable.Type[string] `json:"source,omitempty"`

	// Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD.
	TokenIssuerType *TokenIssuerType `json:"tokenIssuerType,omitempty"`

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

func (s ServicePrincipalRiskDetection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServicePrincipalRiskDetection{}

func (s ServicePrincipalRiskDetection) MarshalJSON() ([]byte, error) {
	type wrapper ServicePrincipalRiskDetection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePrincipalRiskDetection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePrincipalRiskDetection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.servicePrincipalRiskDetection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePrincipalRiskDetection: %+v", err)
	}

	return encoded, nil
}
