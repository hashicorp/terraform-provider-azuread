package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityIncident{}

type SecurityIncident struct {
	// The list of related alerts. Supports $expand.
	Alerts *[]SecurityAlert `json:"alerts,omitempty"`

	// Owner of the incident, or null if no owner is assigned. Free editable text.
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`

	// The specification for the incident. Possible values are: unknown, falsePositive, truePositive,
	// informationalExpectedActivity, unknownFutureValue.
	Classification *SecurityAlertClassification `json:"classification,omitempty"`

	// Array of comments created by the Security Operations (SecOps) team when the incident is managed.
	Comments *[]SecurityAlertComment `json:"comments,omitempty"`

	// Time when the incident was first created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The collection of custom tags that are associated with an incident.
	CustomTags *[]string `json:"customTags,omitempty"`

	// Description of the incident.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Specifies the determination of the incident. Possible values are: unknown, apt, malware, securityPersonnel,
	// securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean,
	// insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
	Determination *SecurityAlertDetermination `json:"determination,omitempty"`

	// The incident name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The URL for the incident page in the Microsoft 365 Defender portal.
	IncidentWebUrl nullable.Type[string] `json:"incidentWebUrl,omitempty"`

	// The identity that last modified the incident.
	LastModifiedBy nullable.Type[string] `json:"lastModifiedBy,omitempty"`

	// Time when the incident was last updated.
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`

	// A rich text string that represents the actions that are reccomnded to take in order to resolve the incident.
	RecommendedActions nullable.Type[string] `json:"recommendedActions,omitempty"`

	// List of hunting Kusto Query Language (KQL) queries related to the incident.
	RecommendedHuntingQueries *[]SecurityRecommendedHuntingQuery `json:"recommendedHuntingQueries,omitempty"`

	// Only populated in case an incident is grouped together with another incident, as part of the logic that processes
	// incidents. In such a case, the status property is redirected.
	RedirectIncidentId nullable.Type[string] `json:"redirectIncidentId,omitempty"`

	// User input that explains the resolution of the incident and the classification choice. This property contains free
	// editable text.
	ResolvingComment nullable.Type[string] `json:"resolvingComment,omitempty"`

	Severity *SecurityAlertSeverity  `json:"severity,omitempty"`
	Status   *SecurityIncidentStatus `json:"status,omitempty"`

	// The overview of an attack. When applicable, the summary contains details of what occurred, impacted assets, and the
	// type of attack.
	Summary nullable.Type[string] `json:"summary,omitempty"`

	// The collection of system tags that are associated with the incident.
	SystemTags *[]string `json:"systemTags,omitempty"`

	// The Microsoft Entra tenant in which the alert was created.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s SecurityIncident) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityIncident{}

func (s SecurityIncident) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIncident
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIncident: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIncident: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.incident"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIncident: %+v", err)
	}

	return encoded, nil
}
