package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionDuration interface {
	SecurityRetentionDuration() BaseSecurityRetentionDurationImpl
}

var _ SecurityRetentionDuration = BaseSecurityRetentionDurationImpl{}

type BaseSecurityRetentionDurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityRetentionDurationImpl) SecurityRetentionDuration() BaseSecurityRetentionDurationImpl {
	return s
}

var _ SecurityRetentionDuration = RawSecurityRetentionDurationImpl{}

// RawSecurityRetentionDurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityRetentionDurationImpl struct {
	securityRetentionDuration BaseSecurityRetentionDurationImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawSecurityRetentionDurationImpl) SecurityRetentionDuration() BaseSecurityRetentionDurationImpl {
	return s.securityRetentionDuration
}

func UnmarshalSecurityRetentionDurationImplementation(input []byte) (SecurityRetentionDuration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRetentionDuration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionDurationForever") {
		var out SecurityRetentionDurationForever
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionDurationForever: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.retentionDurationInDays") {
		var out SecurityRetentionDurationInDays
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRetentionDurationInDays: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityRetentionDurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityRetentionDurationImpl: %+v", err)
	}

	return RawSecurityRetentionDurationImpl{
		securityRetentionDuration: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
