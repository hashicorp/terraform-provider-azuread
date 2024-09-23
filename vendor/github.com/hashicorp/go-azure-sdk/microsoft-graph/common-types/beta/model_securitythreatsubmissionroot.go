package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityThreatSubmissionRoot{}

type SecurityThreatSubmissionRoot struct {
	EmailThreatSubmissionPolicies *[]SecurityEmailThreatSubmissionPolicy `json:"emailThreatSubmissionPolicies,omitempty"`
	EmailThreats                  *[]SecurityEmailThreatSubmission       `json:"emailThreats,omitempty"`
	FileThreats                   *[]SecurityFileThreatSubmission        `json:"fileThreats,omitempty"`
	UrlThreats                    *[]SecurityUrlThreatSubmission         `json:"urlThreats,omitempty"`

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

func (s SecurityThreatSubmissionRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityThreatSubmissionRoot{}

func (s SecurityThreatSubmissionRoot) MarshalJSON() ([]byte, error) {
	type wrapper SecurityThreatSubmissionRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityThreatSubmissionRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityThreatSubmissionRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.threatSubmissionRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityThreatSubmissionRoot: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityThreatSubmissionRoot{}

func (s *SecurityThreatSubmissionRoot) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EmailThreatSubmissionPolicies *[]SecurityEmailThreatSubmissionPolicy `json:"emailThreatSubmissionPolicies,omitempty"`
		UrlThreats                    *[]SecurityUrlThreatSubmission         `json:"urlThreats,omitempty"`
		Id                            *string                                `json:"id,omitempty"`
		ODataId                       *string                                `json:"@odata.id,omitempty"`
		ODataType                     *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EmailThreatSubmissionPolicies = decoded.EmailThreatSubmissionPolicies
	s.UrlThreats = decoded.UrlThreats
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityThreatSubmissionRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["emailThreats"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EmailThreats into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityEmailThreatSubmission, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityEmailThreatSubmissionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EmailThreats' for 'SecurityThreatSubmissionRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EmailThreats = &output
	}

	if v, ok := temp["fileThreats"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling FileThreats into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityFileThreatSubmission, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityFileThreatSubmissionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'FileThreats' for 'SecurityThreatSubmissionRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.FileThreats = &output
	}

	return nil
}
