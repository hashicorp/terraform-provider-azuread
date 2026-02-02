package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityIndicator = SecurityArticleIndicator{}

type SecurityArticleIndicator struct {

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

func (s SecurityArticleIndicator) SecurityIndicator() BaseSecurityIndicatorImpl {
	return BaseSecurityIndicatorImpl{
		Artifact:  s.Artifact,
		Source:    s.Source,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s SecurityArticleIndicator) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityArticleIndicator{}

func (s SecurityArticleIndicator) MarshalJSON() ([]byte, error) {
	type wrapper SecurityArticleIndicator
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityArticleIndicator: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityArticleIndicator: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.articleIndicator"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityArticleIndicator: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityArticleIndicator{}

func (s *SecurityArticleIndicator) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Source    *SecurityIndicatorSource `json:"source,omitempty"`
		Id        *string                  `json:"id,omitempty"`
		ODataId   *string                  `json:"@odata.id,omitempty"`
		ODataType *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Source = decoded.Source

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityArticleIndicator into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["artifact"]; ok {
		impl, err := UnmarshalSecurityArtifactImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Artifact' for 'SecurityArticleIndicator': %+v", err)
		}
		s.Artifact = &impl
	}

	return nil
}
