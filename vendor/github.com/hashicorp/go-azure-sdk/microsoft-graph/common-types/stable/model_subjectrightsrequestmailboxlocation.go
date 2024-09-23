package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestMailboxLocation interface {
	SubjectRightsRequestMailboxLocation() BaseSubjectRightsRequestMailboxLocationImpl
}

var _ SubjectRightsRequestMailboxLocation = BaseSubjectRightsRequestMailboxLocationImpl{}

type BaseSubjectRightsRequestMailboxLocationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSubjectRightsRequestMailboxLocationImpl) SubjectRightsRequestMailboxLocation() BaseSubjectRightsRequestMailboxLocationImpl {
	return s
}

var _ SubjectRightsRequestMailboxLocation = RawSubjectRightsRequestMailboxLocationImpl{}

// RawSubjectRightsRequestMailboxLocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSubjectRightsRequestMailboxLocationImpl struct {
	subjectRightsRequestMailboxLocation BaseSubjectRightsRequestMailboxLocationImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawSubjectRightsRequestMailboxLocationImpl) SubjectRightsRequestMailboxLocation() BaseSubjectRightsRequestMailboxLocationImpl {
	return s.subjectRightsRequestMailboxLocation
}

func UnmarshalSubjectRightsRequestMailboxLocationImplementation(input []byte) (SubjectRightsRequestMailboxLocation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequestMailboxLocation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestAllMailboxLocation") {
		var out SubjectRightsRequestAllMailboxLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestAllMailboxLocation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestEnumeratedMailboxLocation") {
		var out SubjectRightsRequestEnumeratedMailboxLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestEnumeratedMailboxLocation: %+v", err)
		}
		return out, nil
	}

	var parent BaseSubjectRightsRequestMailboxLocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSubjectRightsRequestMailboxLocationImpl: %+v", err)
	}

	return RawSubjectRightsRequestMailboxLocationImpl{
		subjectRightsRequestMailboxLocation: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
