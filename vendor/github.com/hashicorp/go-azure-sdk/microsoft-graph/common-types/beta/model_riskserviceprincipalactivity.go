package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskServicePrincipalActivity struct {
	// Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium
	// customers. Events in tenants without that license will be returned hidden. The possible values are: none, hidden,
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the
	// Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum:
	// adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
	Detail *RiskDetail `json:"detail,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of risk event detected. The possible values are: investigationsThreatIntelligence, generic,
	// adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, anomalousServicePrincipalActivity,
	// maliciousApplication, suspiciousApplication.
	RiskEventTypes *[]string `json:"riskEventTypes,omitempty"`
}
