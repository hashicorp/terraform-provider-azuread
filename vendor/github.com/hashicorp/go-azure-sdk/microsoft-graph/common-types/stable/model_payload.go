package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Payload{}

type Payload struct {
	// The branch of a payload. Possible values are: unknown, other, americanExpress, capitalOne, dhl, docuSign, dropbox,
	// facebook, firstAmerican, microsoft, netflix, scotiabank, sendGrid, stewartTitle, tesco, wellsFargo, syrinxCloud,
	// adobe, teams, zoom, unknownFutureValue.
	Brand *PayloadBrand `json:"brand,omitempty"`

	// The complexity of a payload. Possible values are: unknown, low, medium, high, unknownFutureValue.
	Complexity *PayloadComplexity `json:"complexity,omitempty"`

	// Identity of the user who created the attack simulation and training campaign payload.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the attack simulation and training campaign payload. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the attack simulation and training campaign payload.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Additional details about the payload.
	Detail PayloadDetail `json:"detail"`

	// Display name of the attack simulation and training campaign payload. Supports $filter and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Industry of a payload. Possible values are: unknown, other, banking, businessServices, consumerServices, education,
	// energy, construction, consulting, financialServices, government, hospitality, insurance, legal, courierServices, IT,
	// healthcare, manufacturing, retail, telecom, realEstate, unknownFutureValue.
	Industry *PayloadIndustry `json:"industry,omitempty"`

	// Indicates whether the attack simulation and training campaign payload was created from an automation flow. Supports
	// $filter and $orderby.
	IsAutomated nullable.Type[bool] `json:"isAutomated,omitempty"`

	// Indicates whether the payload is controversial.
	IsControversial nullable.Type[bool] `json:"isControversial,omitempty"`

	// Indicates whether the payload is from any recent event.
	IsCurrentEvent nullable.Type[bool] `json:"isCurrentEvent,omitempty"`

	// Payload language.
	Language nullable.Type[string] `json:"language,omitempty"`

	// Identity of the user who most recently modified the attack simulation and training campaign payload.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time when the attack simulation and training campaign payload was last modified. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Free text tags for a payload.
	PayloadTags *[]string `json:"payloadTags,omitempty"`

	// The payload delivery platform for a simulation. Possible values are: unknown, sms, email, teams, unknownFutureValue.
	Platform *PayloadDeliveryPlatform `json:"platform,omitempty"`

	// Attack type of the attack simulation and training campaign. Supports $filter and $orderby. Possible values are:
	// unknown, social, cloud, endpoint, unknownFutureValue.
	SimulationAttackType *SimulationAttackType `json:"simulationAttackType,omitempty"`

	Source *SimulationContentSource `json:"source,omitempty"`

	// Simulation content status. Supports $filter and $orderby. Possible values are: unknown, draft, ready, archive,
	// delete, unknownFutureValue.
	Status *SimulationContentStatus `json:"status,omitempty"`

	// The social engineering technique used in the attack simulation and training campaign. Supports $filter and $orderby.
	// Possible values are: unknown, credentialHarvesting, attachmentMalware, driveByUrl, linkInAttachment,
	// linkToMalwareFile, unknownFutureValue, oAuthConsentGrant. Use the Prefer: include-unknown-enum-members request header
	// to get the following values from this evolvable enum: oAuthConsentGrant. For more information on the types of social
	// engineering attack techniques, see simulations.
	Technique *SimulationAttackTechnique `json:"technique,omitempty"`

	// The theme of a payload. Possible values are: unknown, other, accountActivation, accountVerification, billing,
	// cleanUpMail, controversial, documentReceived, expense, fax, financeReport, incomingMessages, invoice, itemReceived,
	// loginAlert, mailReceived, password, payment, payroll, personalizedOffer, quarantine, remoteWork, reviewMessage,
	// securityUpdate, serviceSuspended, signatureRequired, upgradeMailboxStorage, verifyMailbox, voicemail, advertisement,
	// employeeEngagement, unknownFutureValue.
	Theme *PayloadTheme `json:"theme,omitempty"`

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

func (s Payload) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Payload{}

func (s Payload) MarshalJSON() ([]byte, error) {
	type wrapper Payload
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Payload: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Payload: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.payload"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Payload: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Payload{}

func (s *Payload) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Brand                *PayloadBrand              `json:"brand,omitempty"`
		Complexity           *PayloadComplexity         `json:"complexity,omitempty"`
		CreatedBy            *EmailIdentity             `json:"createdBy,omitempty"`
		CreatedDateTime      nullable.Type[string]      `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]      `json:"description,omitempty"`
		DisplayName          nullable.Type[string]      `json:"displayName,omitempty"`
		Industry             *PayloadIndustry           `json:"industry,omitempty"`
		IsAutomated          nullable.Type[bool]        `json:"isAutomated,omitempty"`
		IsControversial      nullable.Type[bool]        `json:"isControversial,omitempty"`
		IsCurrentEvent       nullable.Type[bool]        `json:"isCurrentEvent,omitempty"`
		Language             nullable.Type[string]      `json:"language,omitempty"`
		LastModifiedBy       *EmailIdentity             `json:"lastModifiedBy,omitempty"`
		LastModifiedDateTime nullable.Type[string]      `json:"lastModifiedDateTime,omitempty"`
		PayloadTags          *[]string                  `json:"payloadTags,omitempty"`
		Platform             *PayloadDeliveryPlatform   `json:"platform,omitempty"`
		SimulationAttackType *SimulationAttackType      `json:"simulationAttackType,omitempty"`
		Source               *SimulationContentSource   `json:"source,omitempty"`
		Status               *SimulationContentStatus   `json:"status,omitempty"`
		Technique            *SimulationAttackTechnique `json:"technique,omitempty"`
		Theme                *PayloadTheme              `json:"theme,omitempty"`
		Id                   *string                    `json:"id,omitempty"`
		ODataId              *string                    `json:"@odata.id,omitempty"`
		ODataType            *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Brand = decoded.Brand
	s.Complexity = decoded.Complexity
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Industry = decoded.Industry
	s.IsAutomated = decoded.IsAutomated
	s.IsControversial = decoded.IsControversial
	s.IsCurrentEvent = decoded.IsCurrentEvent
	s.Language = decoded.Language
	s.LastModifiedBy = decoded.LastModifiedBy
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.PayloadTags = decoded.PayloadTags
	s.Platform = decoded.Platform
	s.SimulationAttackType = decoded.SimulationAttackType
	s.Source = decoded.Source
	s.Status = decoded.Status
	s.Technique = decoded.Technique
	s.Theme = decoded.Theme
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Payload into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["detail"]; ok {
		impl, err := UnmarshalPayloadDetailImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Detail' for 'Payload': %+v", err)
		}
		s.Detail = impl
	}

	return nil
}
