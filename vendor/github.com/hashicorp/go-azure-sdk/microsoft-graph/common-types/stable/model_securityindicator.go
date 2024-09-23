package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIndicator interface {
	Entity
	SecurityIndicator() BaseSecurityIndicatorImpl
}

var _ SecurityIndicator = BaseSecurityIndicatorImpl{}

type BaseSecurityIndicatorImpl struct {
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

func (s BaseSecurityIndicatorImpl) SecurityIndicator() BaseSecurityIndicatorImpl {
	return s
}

func (s BaseSecurityIndicatorImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityIndicator = RawSecurityIndicatorImpl{}

// RawSecurityIndicatorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityIndicatorImpl struct {
	securityIndicator BaseSecurityIndicatorImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawSecurityIndicatorImpl) SecurityIndicator() BaseSecurityIndicatorImpl {
	return s.securityIndicator
}

func (s RawSecurityIndicatorImpl) Entity() BaseEntityImpl {
	return s.securityIndicator.Entity()
}

var _ json.Marshaler = BaseSecurityIndicatorImpl{}

func (s BaseSecurityIndicatorImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityIndicatorImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityIndicatorImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityIndicatorImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.indicator"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityIndicatorImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityIndicatorImpl{}

func (s *BaseSecurityIndicatorImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Source    *SecurityIndicatorSource `json:"source,omitempty"`
		Id        *string                  `json:"id,omitempty"`
		ODataId   *string                  `json:"@odata.id,omitempty"`
		ODataType *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Source = decoded.Source
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityIndicatorImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["artifact"]; ok {
		impl, err := UnmarshalSecurityArtifactImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Artifact' for 'BaseSecurityIndicatorImpl': %+v", err)
		}
		s.Artifact = &impl
	}

	return nil
}

func UnmarshalSecurityIndicatorImplementation(input []byte) (SecurityIndicator, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityIndicator into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.articleIndicator") {
		var out SecurityArticleIndicator
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityArticleIndicator: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.intelligenceProfileIndicator") {
		var out SecurityIntelligenceProfileIndicator
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIntelligenceProfileIndicator: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityIndicatorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityIndicatorImpl: %+v", err)
	}

	return RawSecurityIndicatorImpl{
		securityIndicator: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
