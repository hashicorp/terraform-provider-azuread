package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityIndicator = SecurityIntelligenceProfileIndicator{}

type SecurityIntelligenceProfileIndicator struct {
	// Designate when an artifact was first used actively in an attack, when a particular sample was compiled, or if neither
	// of those could be ascertained when the file was first seen in public repositories (for example, VirusTotal, ANY.RUN,
	// Hybrid Analysis) or reported publicly.
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`

	// Designate when an artifact was most recently used actively in an attack, when a particular sample was compiled, or if
	// neither of those could be ascertained when the file was first seen in public repositories (for example, VirusTotal,
	// ANY.RUN, Hybrid Analysis) or reported publicly.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// Fields inherited from SecurityIndicator

	Artifact *SecurityArtifact        `json:"artifact,omitempty"`
	Source   *SecurityIndicatorSource `json:"source,omitempty"`

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

func (s SecurityIntelligenceProfileIndicator) SecurityIndicator() BaseSecurityIndicatorImpl {
	return BaseSecurityIndicatorImpl{
		Artifact:  s.Artifact,
		Source:    s.Source,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityIntelligenceProfileIndicator) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityIntelligenceProfileIndicator{}

func (s SecurityIntelligenceProfileIndicator) MarshalJSON() ([]byte, error) {
	type wrapper SecurityIntelligenceProfileIndicator
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityIntelligenceProfileIndicator: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIntelligenceProfileIndicator: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.intelligenceProfileIndicator"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityIntelligenceProfileIndicator: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityIntelligenceProfileIndicator{}

func (s *SecurityIntelligenceProfileIndicator) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FirstSeenDateTime *string                  `json:"firstSeenDateTime,omitempty"`
		LastSeenDateTime  nullable.Type[string]    `json:"lastSeenDateTime,omitempty"`
		Source            *SecurityIndicatorSource `json:"source,omitempty"`
		Id                *string                  `json:"id,omitempty"`
		ODataId           *string                  `json:"@odata.id,omitempty"`
		ODataType         *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FirstSeenDateTime = decoded.FirstSeenDateTime
	s.LastSeenDateTime = decoded.LastSeenDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityIntelligenceProfileIndicator into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["artifact"]; ok {
		impl, err := UnmarshalSecurityArtifactImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Artifact' for 'SecurityIntelligenceProfileIndicator': %+v", err)
		}
		s.Artifact = &impl
	}

	return nil
}
