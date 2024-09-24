package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Simulation{}

type Simulation struct {
	// The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby.
	// Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment,
	// linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Note that you must use the Prefer:
	// include-unknown-enum-members request header to get the following values from this evolvable enum: oAuthConsentGrant.
	// For more information on the types of social engineering attack techniques, see simulations.
	AttackTechnique *SimulationAttackTechnique `json:"attackTechnique,omitempty"`

	// Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are:
	// unknown, social, cloud, endpoint, unknownFutureValue.
	AttackType *SimulationAttackType `json:"attackType,omitempty"`

	// Unique identifier for the attack simulation automation.
	AutomationId nullable.Type[string] `json:"automationId,omitempty"`

	// Date and time of completion of the attack simulation and training campaign. Supports $filter and $orderby.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// Identity of the user who created the attack simulation and training campaign.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time of creation of the attack simulation and training campaign.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the attack simulation and training campaign.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the attack simulation and training campaign. Supports $filter and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Simulation duration in days.
	DurationInDays nullable.Type[int64] `json:"durationInDays,omitempty"`

	// Details about the end user notification setting.
	EndUserNotificationSetting EndUserNotificationSetting `json:"endUserNotificationSetting"`

	// Users excluded from the simulation.
	ExcludedAccountTarget AccountTargetContent `json:"excludedAccountTarget"`

	// Users targeted in the simulation.
	IncludedAccountTarget AccountTargetContent `json:"includedAccountTarget"`

	// Flag that represents if the attack simulation and training campaign was created from a simulation automation flow.
	// Supports $filter and $orderby.
	IsAutomated nullable.Type[bool] `json:"isAutomated,omitempty"`

	// The landing page associated with a simulation during its creation.
	LandingPage *LandingPage `json:"landingPage,omitempty"`

	// Identity of the user who most recently modified the attack simulation and training campaign.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time of the most recent modification of the attack simulation and training campaign.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date and time of the launch/start of the attack simulation and training campaign. Supports $filter and $orderby.
	LaunchDateTime nullable.Type[string] `json:"launchDateTime,omitempty"`

	// The login page associated with a simulation during its creation.
	LoginPage *LoginPage `json:"loginPage,omitempty"`

	// OAuth app details for the OAuth technique.
	OAuthConsentAppDetail *OAuthConsentAppDetail `json:"oAuthConsentAppDetail,omitempty"`

	// The payload associated with a simulation during its creation.
	Payload *Payload `json:"payload,omitempty"`

	// Method of delivery of the phishing payload used in the attack simulation and training campaign. Possible values are:
	// unknown, sms, email, teams, unknownFutureValue.
	PayloadDeliveryPlatform *PayloadDeliveryPlatform `json:"payloadDeliveryPlatform,omitempty"`

	// Report of the attack simulation and training campaign.
	Report *SimulationReport `json:"report,omitempty"`

	// Status of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are: unknown,
	// draft, running, scheduled, succeeded, failed, cancelled, excluded, unknownFutureValue.
	Status *SimulationStatus `json:"status,omitempty"`

	// Details about the training settings for a simulation.
	TrainingSetting TrainingSetting `json:"trainingSetting"`

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

func (s Simulation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Simulation{}

func (s Simulation) MarshalJSON() ([]byte, error) {
	type wrapper Simulation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Simulation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Simulation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.simulation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Simulation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Simulation{}

func (s *Simulation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AttackTechnique         *SimulationAttackTechnique `json:"attackTechnique,omitempty"`
		AttackType              *SimulationAttackType      `json:"attackType,omitempty"`
		AutomationId            nullable.Type[string]      `json:"automationId,omitempty"`
		CompletionDateTime      nullable.Type[string]      `json:"completionDateTime,omitempty"`
		CreatedBy               *EmailIdentity             `json:"createdBy,omitempty"`
		CreatedDateTime         nullable.Type[string]      `json:"createdDateTime,omitempty"`
		Description             nullable.Type[string]      `json:"description,omitempty"`
		DisplayName             nullable.Type[string]      `json:"displayName,omitempty"`
		DurationInDays          nullable.Type[int64]       `json:"durationInDays,omitempty"`
		IsAutomated             nullable.Type[bool]        `json:"isAutomated,omitempty"`
		LandingPage             *LandingPage               `json:"landingPage,omitempty"`
		LastModifiedBy          *EmailIdentity             `json:"lastModifiedBy,omitempty"`
		LastModifiedDateTime    nullable.Type[string]      `json:"lastModifiedDateTime,omitempty"`
		LaunchDateTime          nullable.Type[string]      `json:"launchDateTime,omitempty"`
		LoginPage               *LoginPage                 `json:"loginPage,omitempty"`
		OAuthConsentAppDetail   *OAuthConsentAppDetail     `json:"oAuthConsentAppDetail,omitempty"`
		Payload                 *Payload                   `json:"payload,omitempty"`
		PayloadDeliveryPlatform *PayloadDeliveryPlatform   `json:"payloadDeliveryPlatform,omitempty"`
		Report                  *SimulationReport          `json:"report,omitempty"`
		Status                  *SimulationStatus          `json:"status,omitempty"`
		Id                      *string                    `json:"id,omitempty"`
		ODataId                 *string                    `json:"@odata.id,omitempty"`
		ODataType               *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AttackTechnique = decoded.AttackTechnique
	s.AttackType = decoded.AttackType
	s.AutomationId = decoded.AutomationId
	s.CompletionDateTime = decoded.CompletionDateTime
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.DurationInDays = decoded.DurationInDays
	s.IsAutomated = decoded.IsAutomated
	s.LandingPage = decoded.LandingPage
	s.LastModifiedBy = decoded.LastModifiedBy
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LaunchDateTime = decoded.LaunchDateTime
	s.LoginPage = decoded.LoginPage
	s.OAuthConsentAppDetail = decoded.OAuthConsentAppDetail
	s.Payload = decoded.Payload
	s.PayloadDeliveryPlatform = decoded.PayloadDeliveryPlatform
	s.Report = decoded.Report
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Simulation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["endUserNotificationSetting"]; ok {
		impl, err := UnmarshalEndUserNotificationSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EndUserNotificationSetting' for 'Simulation': %+v", err)
		}
		s.EndUserNotificationSetting = impl
	}

	if v, ok := temp["excludedAccountTarget"]; ok {
		impl, err := UnmarshalAccountTargetContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExcludedAccountTarget' for 'Simulation': %+v", err)
		}
		s.ExcludedAccountTarget = impl
	}

	if v, ok := temp["includedAccountTarget"]; ok {
		impl, err := UnmarshalAccountTargetContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IncludedAccountTarget' for 'Simulation': %+v", err)
		}
		s.IncludedAccountTarget = impl
	}

	if v, ok := temp["trainingSetting"]; ok {
		impl, err := UnmarshalTrainingSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TrainingSetting' for 'Simulation': %+v", err)
		}
		s.TrainingSetting = impl
	}

	return nil
}
