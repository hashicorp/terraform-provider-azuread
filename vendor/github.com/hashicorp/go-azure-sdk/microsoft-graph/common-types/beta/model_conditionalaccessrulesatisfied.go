package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRuleSatisfied struct {
	// Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application,
	// users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client, ipAddressSeenByAzureAD,
	// ipAddressSeenByResourceProvider, unknownFutureValue, servicePrincipals, servicePrincipalRisk, authenticationFlows,
	// insiderRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following
	// values in this evolvable enum: servicePrincipals, servicePrincipalRisk, authenticationFlows, insiderRisk.
	// conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a
	// comma-separated list.
	ConditionalAccessCondition *ConditionalAccessConditions `json:"conditionalAccessCondition,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Refers to the conditional access policy conditions that were satisfied. The possible values are: allApps,
	// firstPartyApps, office365, appId, acr, appFilter, allUsers, guest, groupId, roleId, userId, allDevicePlatforms,
	// devicePlatform, allLocations, insideCorpnet, allTrustedLocations, locationId, allDevices, deviceFilter, deviceState,
	// unknownFutureValue, deviceFilterIncludeRuleNotMatched, allDeviceStates, anonymizedIPAddress, unfamiliarFeatures,
	// nationStateIPAddress, realTimeThreatIntelligence, internalGuest, b2bCollaborationGuest, b2bCollaborationMember,
	// b2bDirectConnectUser, otherExternalUser, serviceProvider, microsoftAdminPortals, deviceCodeFlow, accountTransfer,
	// insiderRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following
	// values in this evolvable enum: deviceFilterIncludeRuleNotMatched, allDeviceStates, anonymizedIPAddress,
	// unfamiliarFeatures, nationStateIPAddress, realTimeThreatIntelligence, internalGuest, b2bCollaborationGuest,
	// b2bCollaborationMember, b2bDirectConnectUser, otherExternalUser, serviceProvider, microsoftAdminPortals,
	// deviceCodeFlow, accountTransfer, insiderRisk.
	RuleSatisfied *ConditionalAccessRule `json:"ruleSatisfied,omitempty"`
}
