package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Security struct {
	// Notifications for suspicious or potential security issues in a customer’s tenant.
	Alerts *[]Alert `json:"alerts,omitempty"`

	// A collection of alerts in Microsoft 365 Defender.
	Alertsv2 *[]SecurityAlert `json:"alerts_v2,omitempty"`

	// Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
	AttackSimulation *AttackSimulationRoot `json:"attackSimulation,omitempty"`

	AuditLog                 *SecurityAuditCoreRoot     `json:"auditLog,omitempty"`
	Cases                    *SecurityCasesRoot         `json:"cases,omitempty"`
	CloudAppSecurityProfiles *[]CloudAppSecurityProfile `json:"cloudAppSecurityProfiles,omitempty"`

	// Enables read and other actions on collaborative entities in Microsoft Defender.
	Collaboration *SecurityCollaborationRoot `json:"collaboration,omitempty"`

	DataDiscovery             *SecurityDataDiscoveryRoot       `json:"dataDiscovery,omitempty"`
	DataSecurityAndGovernance *TenantDataSecurityAndGovernance `json:"dataSecurityAndGovernance,omitempty"`
	DomainSecurityProfiles    *[]DomainSecurityProfile         `json:"domainSecurityProfiles,omitempty"`
	FileSecurityProfiles      *[]FileSecurityProfile           `json:"fileSecurityProfiles,omitempty"`
	HostSecurityProfiles      *[]HostSecurityProfile           `json:"hostSecurityProfiles,omitempty"`
	IPSecurityProfiles        *[]IPSecurityProfile             `json:"ipSecurityProfiles,omitempty"`

	// A container for security identities APIs.
	Identities *SecurityIdentityContainer `json:"identities,omitempty"`

	// A collection of incidents in Microsoft 365 Defender, each of which is a set of correlated alerts and associated
	// metadata that reflects the story of an attack.
	Incidents *[]SecurityIncident `json:"incidents,omitempty"`

	InformationProtection *SecurityInformationProtection `json:"informationProtection,omitempty"`
	Labels                *SecurityLabelsRoot            `json:"labels,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A container that safeguards the Microsoft Azure resources of Microsoft Cloud Solution Provider (CSP) partners’
	// customers, including alerts, scores, and all aspects of security.
	Partner *PartnerSecurityPartnerSecurity `json:"partner,omitempty"`

	ProviderTenantSettings     *[]ProviderTenantSetting     `json:"providerTenantSettings,omitempty"`
	Rules                      *SecurityRulesRoot           `json:"rules,omitempty"`
	SecureScoreControlProfiles *[]SecureScoreControlProfile `json:"secureScoreControlProfiles,omitempty"`

	// Measurements of tenants’ security posture to help protect them from threats.
	SecureScores *[]SecureScore `json:"secureScores,omitempty"`

	SecurityActions       *[]SecurityAction           `json:"securityActions,omitempty"`
	SubjectRightsRequests *[]SubjectRightsRequest     `json:"subjectRightsRequests,omitempty"`
	ThreatIntelligence    *SecurityThreatIntelligence `json:"threatIntelligence,omitempty"`

	// A threat submission sent to Microsoft; for example, a suspicious email threat, URL threat, or file threat.
	ThreatSubmission *SecurityThreatSubmissionRoot `json:"threatSubmission,omitempty"`

	TiIndicators         *[]TiIndicator            `json:"tiIndicators,omitempty"`
	TriggerTypes         *SecurityTriggerTypesRoot `json:"triggerTypes,omitempty"`
	Triggers             *SecurityTriggersRoot     `json:"triggers,omitempty"`
	UserSecurityProfiles *[]UserSecurityProfile    `json:"userSecurityProfiles,omitempty"`
}
