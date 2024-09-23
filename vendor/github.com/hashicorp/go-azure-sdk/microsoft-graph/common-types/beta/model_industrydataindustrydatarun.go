package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IndustryDataIndustryDataRun{}

type IndustryDataIndustryDataRun struct {
	// The set of activities performed during the run.
	Activities *[]IndustryDataIndustryDataRunActivity `json:"activities,omitempty"`

	// An error object to diagnose critical failures in the run.
	BlockingError *PublicError `json:"blockingError,omitempty"`

	// The name of the run for rendering in a user interface.
	DisplayName *string `json:"displayName,omitempty"`

	// The date and time when the run finished or null if the run is still in-progress. The Timestamp type represents date
	// and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The date and time when the run started. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime *string `json:"startDateTime,omitempty"`

	Status *IndustryDataIndustryDataRunStatus `json:"status,omitempty"`

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

func (s IndustryDataIndustryDataRun) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataIndustryDataRun{}

func (s IndustryDataIndustryDataRun) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataIndustryDataRun
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataIndustryDataRun: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataRun: %+v", err)
	}

	delete(decoded, "blockingError")
	delete(decoded, "displayName")
	delete(decoded, "endDateTime")
	delete(decoded, "startDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.industryDataRun"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataIndustryDataRun: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataIndustryDataRun{}

func (s *IndustryDataIndustryDataRun) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BlockingError *PublicError                       `json:"blockingError,omitempty"`
		DisplayName   *string                            `json:"displayName,omitempty"`
		EndDateTime   nullable.Type[string]              `json:"endDateTime,omitempty"`
		StartDateTime *string                            `json:"startDateTime,omitempty"`
		Status        *IndustryDataIndustryDataRunStatus `json:"status,omitempty"`
		Id            *string                            `json:"id,omitempty"`
		ODataId       *string                            `json:"@odata.id,omitempty"`
		ODataType     *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BlockingError = decoded.BlockingError
	s.DisplayName = decoded.DisplayName
	s.EndDateTime = decoded.EndDateTime
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataIndustryDataRun into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Activities into list []json.RawMessage: %+v", err)
		}

		output := make([]IndustryDataIndustryDataRunActivity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIndustryDataIndustryDataRunActivityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Activities' for 'IndustryDataIndustryDataRun': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Activities = &output
	}

	return nil
}
