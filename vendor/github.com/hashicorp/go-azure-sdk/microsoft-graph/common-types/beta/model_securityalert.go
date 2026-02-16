package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityAlert{}

type SecurityAlert struct {
	// The adversary or activity group that is associated with this alert.
	ActorDisplayName nullable.Type[string] `json:"actorDisplayName,omitempty"`

	// A collection of other alert properties, including user-defined properties. Any custom details defined in the alert,
	// and any dynamic content in the alert details, are stored here.
	AdditionalData *SecurityDictionary `json:"additionalData,omitempty"`

	// The ID of the policy that generated the alert, and populated when there is a specific policy that generated the
	// alert, whether configured by a customer or a built-in policy.
	AlertPolicyId nullable.Type[string] `json:"alertPolicyId,omitempty"`

	// URL for the Microsoft 365 Defender portal alert page.
	AlertWebUrl nullable.Type[string] `json:"alertWebUrl,omitempty"`

	// Owner of the alert, or null if no owner is assigned.
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`

	// The attack kill-chain category that the alert belongs to. Aligned with the MITRE ATT&CK framework.
	Category nullable.Type[string] `json:"category,omitempty"`

	// Specifies whether the alert represents a true threat. Possible values are: unknown, falsePositive, truePositive,
	// informationalExpectedActivity, unknownFutureValue.
	Classification *SecurityAlertClassification `json:"classification,omitempty"`

	// Array of comments created by the Security Operations (SecOps) team during the alert management process.
	Comments *[]SecurityAlertComment `json:"comments,omitempty"`

	// Time when Microsoft 365 Defender created the alert.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// User defined custom fields with string values.
	CustomDetails *SecurityDictionary `json:"customDetails,omitempty"`

	// String value describing each alert.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Detection technology or sensor that identified the notable component or activity. Possible values are: unknown,
	// microsoftDefenderForEndpoint, antivirus, smartScreen, customTi, microsoftDefenderForOffice365,
	// automatedInvestigation, microsoftThreatExperts, customDetection, microsoftDefenderForIdentity, cloudAppSecurity,
	// microsoft365Defender, azureAdIdentityProtection, manual, microsoftDataLossPrevention, appGovernancePolicy,
	// appGovernanceDetection, unknownFutureValue, microsoftDefenderForCloud, microsoftDefenderForIoT,
	// microsoftDefenderForServers, microsoftDefenderForStorage, microsoftDefenderForDNS, microsoftDefenderForDatabases,
	// microsoftDefenderForContainers, microsoftDefenderForNetwork, microsoftDefenderForAppService,
	// microsoftDefenderForKeyVault, microsoftDefenderForResourceManager, microsoftDefenderForApiManagement,
	// microsoftSentinel, nrtAlerts, scheduledAlerts, microsoftDefenderThreatIntelligenceAnalytics, builtInMl,
	// microsoftThreatIntelligence. Use the Prefer: include-unknown-enum-members request header to get the following values
	// in this evolvable enum: microsoftDefenderForCloud, microsoftDefenderForIoT, microsoftDefenderForServers,
	// microsoftDefenderForStorage, microsoftDefenderForDNS, microsoftDefenderForDatabases, microsoftDefenderForContainers,
	// microsoftDefenderForNetwork, microsoftDefenderForAppService, microsoftDefenderForKeyVault,
	// microsoftDefenderForResourceManager, microsoftDefenderForApiManagement, microsoftSentinel, nrtAlerts,
	// scheduledAlerts, microsoftDefenderThreatIntelligenceAnalytics, builtInMl, microsoftThreatIntelligence.
	DetectionSource *SecurityDetectionSource `json:"detectionSource,omitempty"`

	// The ID of the detector that triggered the alert.
	DetectorId nullable.Type[string] `json:"detectorId,omitempty"`

	// Specifies the result of the investigation, whether the alert represents a true attack, and if so, the nature of the
	// attack. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other,
	// multiStagedAttack, compromisedAccount, phishing, maliciousUserActivity, notMalicious, notEnoughDataToValidate,
	// confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
	Determination *SecurityAlertDetermination `json:"determination,omitempty"`

	// Collection of evidence related to the alert.
	Evidence *[]SecurityAlertEvidence `json:"evidence,omitempty"`

	// The earliest activity associated with the alert.
	FirstActivityDateTime nullable.Type[string] `json:"firstActivityDateTime,omitempty"`

	// Unique identifier to represent the incident this alert resource is associated with.
	IncidentId nullable.Type[string] `json:"incidentId,omitempty"`

	// URL for the incident page in the Microsoft 365 Defender portal.
	IncidentWebUrl nullable.Type[string] `json:"incidentWebUrl,omitempty"`

	// The oldest activity associated with the alert.
	LastActivityDateTime nullable.Type[string] `json:"lastActivityDateTime,omitempty"`

	// Time when the alert was last updated at Microsoft 365 Defender.
	LastUpdateDateTime nullable.Type[string] `json:"lastUpdateDateTime,omitempty"`

	// The attack techniques, as aligned with the MITRE ATT&CK framework.
	MitreTechniques *[]string `json:"mitreTechniques,omitempty"`

	// The name of the product which published this alert.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The ID of the alert as it appears in the security provider product that generated the alert.
	ProviderAlertId nullable.Type[string] `json:"providerAlertId,omitempty"`

	// Recommended response and remediation actions to take in the event this alert was generated.
	RecommendedActions nullable.Type[string] `json:"recommendedActions,omitempty"`

	// Time when the alert was resolved.
	ResolvedDateTime nullable.Type[string] `json:"resolvedDateTime,omitempty"`

	ServiceSource *SecurityServiceSource `json:"serviceSource,omitempty"`
	Severity      *SecurityAlertSeverity `json:"severity,omitempty"`
	Status        *SecurityAlertStatus   `json:"status,omitempty"`

	// The system tags associated with the alert.
	SystemTags *[]string `json:"systemTags,omitempty"`

	// The Microsoft Entra tenant the alert was created in.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The threat associated with this alert.
	ThreatDisplayName nullable.Type[string] `json:"threatDisplayName,omitempty"`

	// Threat family associated with this alert.
	ThreatFamilyName nullable.Type[string] `json:"threatFamilyName,omitempty"`

	// Brief identifying string value describing the alert.
	Title nullable.Type[string] `json:"title,omitempty"`

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

func (s SecurityAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityAlert{}

func (s SecurityAlert) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.alert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAlert: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityAlert{}

func (s *SecurityAlert) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActorDisplayName      nullable.Type[string]        `json:"actorDisplayName,omitempty"`
		AdditionalData        *SecurityDictionary          `json:"additionalData,omitempty"`
		AlertPolicyId         nullable.Type[string]        `json:"alertPolicyId,omitempty"`
		AlertWebUrl           nullable.Type[string]        `json:"alertWebUrl,omitempty"`
		AssignedTo            nullable.Type[string]        `json:"assignedTo,omitempty"`
		Category              nullable.Type[string]        `json:"category,omitempty"`
		Classification        *SecurityAlertClassification `json:"classification,omitempty"`
		Comments              *[]SecurityAlertComment      `json:"comments,omitempty"`
		CreatedDateTime       nullable.Type[string]        `json:"createdDateTime,omitempty"`
		CustomDetails         *SecurityDictionary          `json:"customDetails,omitempty"`
		Description           nullable.Type[string]        `json:"description,omitempty"`
		DetectionSource       *SecurityDetectionSource     `json:"detectionSource,omitempty"`
		DetectorId            nullable.Type[string]        `json:"detectorId,omitempty"`
		Determination         *SecurityAlertDetermination  `json:"determination,omitempty"`
		FirstActivityDateTime nullable.Type[string]        `json:"firstActivityDateTime,omitempty"`
		IncidentId            nullable.Type[string]        `json:"incidentId,omitempty"`
		IncidentWebUrl        nullable.Type[string]        `json:"incidentWebUrl,omitempty"`
		LastActivityDateTime  nullable.Type[string]        `json:"lastActivityDateTime,omitempty"`
		LastUpdateDateTime    nullable.Type[string]        `json:"lastUpdateDateTime,omitempty"`
		MitreTechniques       *[]string                    `json:"mitreTechniques,omitempty"`
		ProductName           nullable.Type[string]        `json:"productName,omitempty"`
		ProviderAlertId       nullable.Type[string]        `json:"providerAlertId,omitempty"`
		RecommendedActions    nullable.Type[string]        `json:"recommendedActions,omitempty"`
		ResolvedDateTime      nullable.Type[string]        `json:"resolvedDateTime,omitempty"`
		ServiceSource         *SecurityServiceSource       `json:"serviceSource,omitempty"`
		Severity              *SecurityAlertSeverity       `json:"severity,omitempty"`
		Status                *SecurityAlertStatus         `json:"status,omitempty"`
		SystemTags            *[]string                    `json:"systemTags,omitempty"`
		TenantId              nullable.Type[string]        `json:"tenantId,omitempty"`
		ThreatDisplayName     nullable.Type[string]        `json:"threatDisplayName,omitempty"`
		ThreatFamilyName      nullable.Type[string]        `json:"threatFamilyName,omitempty"`
		Title                 nullable.Type[string]        `json:"title,omitempty"`
		Id                    *string                      `json:"id,omitempty"`
		ODataId               *string                      `json:"@odata.id,omitempty"`
		ODataType             *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActorDisplayName = decoded.ActorDisplayName
	s.AdditionalData = decoded.AdditionalData
	s.AlertPolicyId = decoded.AlertPolicyId
	s.AlertWebUrl = decoded.AlertWebUrl
	s.AssignedTo = decoded.AssignedTo
	s.Category = decoded.Category
	s.Classification = decoded.Classification
	s.Comments = decoded.Comments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomDetails = decoded.CustomDetails
	s.Description = decoded.Description
	s.DetectionSource = decoded.DetectionSource
	s.DetectorId = decoded.DetectorId
	s.Determination = decoded.Determination
	s.FirstActivityDateTime = decoded.FirstActivityDateTime
	s.IncidentId = decoded.IncidentId
	s.IncidentWebUrl = decoded.IncidentWebUrl
	s.LastActivityDateTime = decoded.LastActivityDateTime
	s.LastUpdateDateTime = decoded.LastUpdateDateTime
	s.MitreTechniques = decoded.MitreTechniques
	s.ProductName = decoded.ProductName
	s.ProviderAlertId = decoded.ProviderAlertId
	s.RecommendedActions = decoded.RecommendedActions
	s.ResolvedDateTime = decoded.ResolvedDateTime
	s.ServiceSource = decoded.ServiceSource
	s.Severity = decoded.Severity
	s.Status = decoded.Status
	s.SystemTags = decoded.SystemTags
	s.TenantId = decoded.TenantId
	s.ThreatDisplayName = decoded.ThreatDisplayName
	s.ThreatFamilyName = decoded.ThreatFamilyName
	s.Title = decoded.Title
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityAlert into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["evidence"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Evidence into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityAlertEvidence, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityAlertEvidenceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Evidence' for 'SecurityAlert': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Evidence = &output
	}

	return nil
}
