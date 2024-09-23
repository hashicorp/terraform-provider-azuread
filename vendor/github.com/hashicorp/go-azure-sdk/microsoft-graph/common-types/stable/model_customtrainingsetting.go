package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TrainingSetting = CustomTrainingSetting{}

type CustomTrainingSetting struct {
	// A user collection that specifies to whom the training should be assigned. Possible values are: none, allUsers,
	// clickedPayload, compromised, reportedPhish, readButNotClicked, didNothing, unknownFutureValue.
	AssignedTo *[]TrainingAssignedTo `json:"assignedTo,omitempty"`

	// The description of the custom training setting.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the custom training setting.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Training duration.
	DurationInMinutes nullable.Type[int64] `json:"durationInMinutes,omitempty"`

	// The training URL.
	Url nullable.Type[string] `json:"url,omitempty"`

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

func (s CustomTrainingSetting) TrainingSetting() BaseTrainingSettingImpl {
	return BaseTrainingSettingImpl{
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		SettingType: s.SettingType,
	}
}

var _ json.Marshaler = CustomTrainingSetting{}

func (s CustomTrainingSetting) MarshalJSON() ([]byte, error) {
	type wrapper CustomTrainingSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomTrainingSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomTrainingSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customTrainingSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomTrainingSetting: %+v", err)
	}

	return encoded, nil
}
