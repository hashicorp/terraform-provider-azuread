package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DataClassificationService{}

type DataClassificationService struct {
	ClassifyFileJobs        *[]JobResponseBase       `json:"classifyFileJobs,omitempty"`
	ClassifyTextJobs        *[]JobResponseBase       `json:"classifyTextJobs,omitempty"`
	EvaluateDlpPoliciesJobs *[]JobResponseBase       `json:"evaluateDlpPoliciesJobs,omitempty"`
	EvaluateLabelJobs       *[]JobResponseBase       `json:"evaluateLabelJobs,omitempty"`
	ExactMatchDataStores    *[]ExactMatchDataStore   `json:"exactMatchDataStores,omitempty"`
	ExactMatchUploadAgents  *[]ExactMatchUploadAgent `json:"exactMatchUploadAgents,omitempty"`
	Jobs                    *[]JobResponseBase       `json:"jobs,omitempty"`
	SensitiveTypes          *[]SensitiveType         `json:"sensitiveTypes,omitempty"`
	SensitivityLabels       *[]SensitivityLabel      `json:"sensitivityLabels,omitempty"`

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

func (s DataClassificationService) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DataClassificationService{}

func (s DataClassificationService) MarshalJSON() ([]byte, error) {
	type wrapper DataClassificationService
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataClassificationService: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataClassificationService: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dataClassificationService"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataClassificationService: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DataClassificationService{}

func (s *DataClassificationService) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExactMatchDataStores   *[]ExactMatchDataStore   `json:"exactMatchDataStores,omitempty"`
		ExactMatchUploadAgents *[]ExactMatchUploadAgent `json:"exactMatchUploadAgents,omitempty"`
		SensitiveTypes         *[]SensitiveType         `json:"sensitiveTypes,omitempty"`
		SensitivityLabels      *[]SensitivityLabel      `json:"sensitivityLabels,omitempty"`
		Id                     *string                  `json:"id,omitempty"`
		ODataId                *string                  `json:"@odata.id,omitempty"`
		ODataType              *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExactMatchDataStores = decoded.ExactMatchDataStores
	s.ExactMatchUploadAgents = decoded.ExactMatchUploadAgents
	s.SensitiveTypes = decoded.SensitiveTypes
	s.SensitivityLabels = decoded.SensitivityLabels
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataClassificationService into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["classifyFileJobs"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ClassifyFileJobs into list []json.RawMessage: %+v", err)
		}

		output := make([]JobResponseBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalJobResponseBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ClassifyFileJobs' for 'DataClassificationService': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ClassifyFileJobs = &output
	}

	if v, ok := temp["classifyTextJobs"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ClassifyTextJobs into list []json.RawMessage: %+v", err)
		}

		output := make([]JobResponseBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalJobResponseBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ClassifyTextJobs' for 'DataClassificationService': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ClassifyTextJobs = &output
	}

	if v, ok := temp["evaluateDlpPoliciesJobs"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EvaluateDlpPoliciesJobs into list []json.RawMessage: %+v", err)
		}

		output := make([]JobResponseBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalJobResponseBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EvaluateDlpPoliciesJobs' for 'DataClassificationService': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EvaluateDlpPoliciesJobs = &output
	}

	if v, ok := temp["evaluateLabelJobs"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EvaluateLabelJobs into list []json.RawMessage: %+v", err)
		}

		output := make([]JobResponseBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalJobResponseBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EvaluateLabelJobs' for 'DataClassificationService': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EvaluateLabelJobs = &output
	}

	if v, ok := temp["jobs"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Jobs into list []json.RawMessage: %+v", err)
		}

		output := make([]JobResponseBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalJobResponseBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Jobs' for 'DataClassificationService': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Jobs = &output
	}

	return nil
}
