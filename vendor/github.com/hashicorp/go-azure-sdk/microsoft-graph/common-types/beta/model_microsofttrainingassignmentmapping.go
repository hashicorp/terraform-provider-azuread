package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TrainingSetting = MicrosoftTrainingAssignmentMapping{}

type MicrosoftTrainingAssignmentMapping struct {
	// A user collection that specifies to whom the training should be assigned. Possible values are: none, allUsers,
	// clickedPayload, compromised, reportedPhish, readButNotClicked, didNothing, unknownFutureValue.
	AssignedTo *[]TrainingAssignedTo `json:"assignedTo,omitempty"`

	Training *Training `json:"training,omitempty"`

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

func (s MicrosoftTrainingAssignmentMapping) TrainingSetting() BaseTrainingSettingImpl {
	return BaseTrainingSettingImpl{
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		SettingType: s.SettingType,
	}
}

var _ json.Marshaler = MicrosoftTrainingAssignmentMapping{}

func (s MicrosoftTrainingAssignmentMapping) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftTrainingAssignmentMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftTrainingAssignmentMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftTrainingAssignmentMapping: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftTrainingAssignmentMapping"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftTrainingAssignmentMapping: %+v", err)
	}

	return encoded, nil
}
