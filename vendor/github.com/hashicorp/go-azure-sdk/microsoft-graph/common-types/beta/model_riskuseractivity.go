package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskUserActivity struct {
	// The possible values are none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange,
	// userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe,
	// userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden,
	// adminConfirmedUserCompromised, unknownFutureValue.
	Detail *RiskDetail `json:"detail,omitempty"`

	// List of risk event types. Deprecated. Use riskEventType instead.
	EventTypes *[]RiskEventType `json:"eventTypes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of risk event detected. The possible values are: anonymizedIPAddress, investigationsThreatIntelligence,
	// investigationsThreatIntelligenceSigninLinked,leakedCredentials, maliciousIPAddress,
	// maliciousIPAddressValidCredentialsBlockedIP, malwareInfectedIPAddress, mcasImpossibleTravel,
	// mcasSuspiciousInboxManipulationRules, suspiciousAPITraffic, suspiciousIPAddress, unfamiliarFeatures, unlikelyTravel.
	// For more information about each value, see Risk types and detection.
	RiskEventTypes *[]string `json:"riskEventTypes,omitempty"`
}
