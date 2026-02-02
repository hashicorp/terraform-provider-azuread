package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TrainingCampaign{}

type TrainingCampaign struct {
	// Details about the schedule and current status for a training campaign
	CampaignSchedule *CampaignSchedule `json:"campaignSchedule,omitempty"`

	// Identity of the user who created the training campaign
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time of creation of the training campaign.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the training campaign.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the training campaign. Supports $filter and $orderby.
	DisplayName *string `json:"displayName,omitempty"`

	// Details about the end user notification setting.
	EndUserNotificationSetting EndUserNotificationSetting `json:"endUserNotificationSetting"`

	// Users excluded from the training campaign.
	ExcludedAccountTarget AccountTargetContent `json:"excludedAccountTarget"`

	// Users targeted in the training campaign.
	IncludedAccountTarget AccountTargetContent `json:"includedAccountTarget"`

	// Identity of the user who most recently modified the training campaign.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time of the most recent modification of the training campaign.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Report of the training campaign.
	Report *TrainingCampaignReport `json:"report,omitempty"`

	// Details about the training settings for a training campaign.
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

func (s TrainingCampaign) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TrainingCampaign{}

func (s TrainingCampaign) MarshalJSON() ([]byte, error) {
	type wrapper TrainingCampaign
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TrainingCampaign: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TrainingCampaign: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trainingCampaign"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TrainingCampaign: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TrainingCampaign{}

func (s *TrainingCampaign) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CampaignSchedule     *CampaignSchedule       `json:"campaignSchedule,omitempty"`
		CreatedBy            *EmailIdentity          `json:"createdBy,omitempty"`
		CreatedDateTime      nullable.Type[string]   `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]   `json:"description,omitempty"`
		DisplayName          *string                 `json:"displayName,omitempty"`
		LastModifiedBy       *EmailIdentity          `json:"lastModifiedBy,omitempty"`
		LastModifiedDateTime nullable.Type[string]   `json:"lastModifiedDateTime,omitempty"`
		Report               *TrainingCampaignReport `json:"report,omitempty"`
		Id                   *string                 `json:"id,omitempty"`
		ODataId              *string                 `json:"@odata.id,omitempty"`
		ODataType            *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CampaignSchedule = decoded.CampaignSchedule
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedBy = decoded.LastModifiedBy
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Report = decoded.Report
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TrainingCampaign into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["endUserNotificationSetting"]; ok {
		impl, err := UnmarshalEndUserNotificationSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EndUserNotificationSetting' for 'TrainingCampaign': %+v", err)
		}
		s.EndUserNotificationSetting = impl
	}

	if v, ok := temp["excludedAccountTarget"]; ok {
		impl, err := UnmarshalAccountTargetContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExcludedAccountTarget' for 'TrainingCampaign': %+v", err)
		}
		s.ExcludedAccountTarget = impl
	}

	if v, ok := temp["includedAccountTarget"]; ok {
		impl, err := UnmarshalAccountTargetContentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IncludedAccountTarget' for 'TrainingCampaign': %+v", err)
		}
		s.IncludedAccountTarget = impl
	}

	if v, ok := temp["trainingSetting"]; ok {
		impl, err := UnmarshalTrainingSettingImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TrainingSetting' for 'TrainingCampaign': %+v", err)
		}
		s.TrainingSetting = impl
	}

	return nil
}
