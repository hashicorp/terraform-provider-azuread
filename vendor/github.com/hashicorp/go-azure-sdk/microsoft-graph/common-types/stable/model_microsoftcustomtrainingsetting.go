package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TrainingSetting = MicrosoftCustomTrainingSetting{}

type MicrosoftCustomTrainingSetting struct {
	// The completion date and time of the training. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// The mapping details of the associated training.
	TrainingAssignmentMappings *[]MicrosoftTrainingAssignmentMapping `json:"trainingAssignmentMappings,omitempty"`

	// The training completion duration that needs to be provided before scheduling the training. Possible values are: week,
	// fortnite, month, unknownFutureValue.
	TrainingCompletionDuration *TrainingCompletionDuration `json:"trainingCompletionDuration,omitempty"`

	// Fields inherited from TrainingSetting

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of setting. Possible values are: microsoftCustom, microsoftManaged, noTraining, custom, unknownFutureValue.
	SettingType *TrainingSettingType `json:"settingType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MicrosoftCustomTrainingSetting) TrainingSetting() BaseTrainingSettingImpl {
	return BaseTrainingSettingImpl{
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		SettingType: s.SettingType,
	}
}

var _ json.Marshaler = MicrosoftCustomTrainingSetting{}

func (s MicrosoftCustomTrainingSetting) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftCustomTrainingSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftCustomTrainingSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftCustomTrainingSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftCustomTrainingSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftCustomTrainingSetting: %+v", err)
	}

	return encoded, nil
}
