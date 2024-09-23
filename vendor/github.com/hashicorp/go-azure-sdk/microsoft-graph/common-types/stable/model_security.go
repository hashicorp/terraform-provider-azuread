package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Security{}

type Security struct {
	Alerts *[]Alert `json:"alerts,omitempty"`

	// A collection of alerts in Microsoft 365 Defender.
	Alertsv2 *[]SecurityAlert `json:"alerts_v2,omitempty"`

	AttackSimulation *AttackSimulationRoot      `json:"attackSimulation,omitempty"`
	Cases            *SecurityCasesRoot         `json:"cases,omitempty"`
	Identities       *SecurityIdentityContainer `json:"identities,omitempty"`

	// A collection of incidents in Microsoft 365 Defender, each of which is a set of correlated alerts and associated
	// metadata that reflects the story of an attack.
	Incidents *[]SecurityIncident `json:"incidents,omitempty"`

	Labels                     *SecurityLabelsRoot          `json:"labels,omitempty"`
	SecureScoreControlProfiles *[]SecureScoreControlProfile `json:"secureScoreControlProfiles,omitempty"`
	SecureScores               *[]SecureScore               `json:"secureScores,omitempty"`
	SubjectRightsRequests      *[]SubjectRightsRequest      `json:"subjectRightsRequests,omitempty"`
	ThreatIntelligence         *SecurityThreatIntelligence  `json:"threatIntelligence,omitempty"`
	TriggerTypes               *SecurityTriggerTypesRoot    `json:"triggerTypes,omitempty"`
	Triggers                   *SecurityTriggersRoot        `json:"triggers,omitempty"`

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

func (s Security) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Security{}

func (s Security) MarshalJSON() ([]byte, error) {
	type wrapper Security
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Security: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Security: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Security: %+v", err)
	}

	return encoded, nil
}
