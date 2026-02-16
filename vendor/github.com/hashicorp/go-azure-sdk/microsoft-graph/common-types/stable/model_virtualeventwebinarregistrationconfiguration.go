package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VirtualEventRegistrationConfiguration = VirtualEventWebinarRegistrationConfiguration{}

type VirtualEventWebinarRegistrationConfiguration struct {
	IsManualApprovalEnabled nullable.Type[bool] `json:"isManualApprovalEnabled,omitempty"`
	IsWaitlistEnabled       nullable.Type[bool] `json:"isWaitlistEnabled,omitempty"`

	// Fields inherited from VirtualEventRegistrationConfiguration

	// Total capacity of the virtual event.
	Capacity nullable.Type[int64] `json:"capacity,omitempty"`

	// Registration questions.
	Questions *[]VirtualEventRegistrationQuestionBase `json:"questions,omitempty"`

	// Registration URL of the virtual event.
	RegistrationWebUrl nullable.Type[string] `json:"registrationWebUrl,omitempty"`

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

func (s VirtualEventWebinarRegistrationConfiguration) VirtualEventRegistrationConfiguration() BaseVirtualEventRegistrationConfigurationImpl {
	return BaseVirtualEventRegistrationConfigurationImpl{
		Capacity:           s.Capacity,
		Questions:          s.Questions,
		RegistrationWebUrl: s.RegistrationWebUrl,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s VirtualEventWebinarRegistrationConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventWebinarRegistrationConfiguration{}

func (s VirtualEventWebinarRegistrationConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventWebinarRegistrationConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventWebinarRegistrationConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventWebinarRegistrationConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventWebinarRegistrationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventWebinarRegistrationConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VirtualEventWebinarRegistrationConfiguration{}

func (s *VirtualEventWebinarRegistrationConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsManualApprovalEnabled nullable.Type[bool]   `json:"isManualApprovalEnabled,omitempty"`
		IsWaitlistEnabled       nullable.Type[bool]   `json:"isWaitlistEnabled,omitempty"`
		Capacity                nullable.Type[int64]  `json:"capacity,omitempty"`
		RegistrationWebUrl      nullable.Type[string] `json:"registrationWebUrl,omitempty"`
		Id                      *string               `json:"id,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsManualApprovalEnabled = decoded.IsManualApprovalEnabled
	s.IsWaitlistEnabled = decoded.IsWaitlistEnabled
	s.Capacity = decoded.Capacity
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RegistrationWebUrl = decoded.RegistrationWebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualEventWebinarRegistrationConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["questions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Questions into list []json.RawMessage: %+v", err)
		}

		output := make([]VirtualEventRegistrationQuestionBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalVirtualEventRegistrationQuestionBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Questions' for 'VirtualEventWebinarRegistrationConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Questions = &output
	}

	return nil
}
