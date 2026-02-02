package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestSiteLocation interface {
	SubjectRightsRequestSiteLocation() BaseSubjectRightsRequestSiteLocationImpl
}

var _ SubjectRightsRequestSiteLocation = BaseSubjectRightsRequestSiteLocationImpl{}

type BaseSubjectRightsRequestSiteLocationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSubjectRightsRequestSiteLocationImpl) SubjectRightsRequestSiteLocation() BaseSubjectRightsRequestSiteLocationImpl {
	return s
}

var _ SubjectRightsRequestSiteLocation = RawSubjectRightsRequestSiteLocationImpl{}

// RawSubjectRightsRequestSiteLocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSubjectRightsRequestSiteLocationImpl struct {
	subjectRightsRequestSiteLocation BaseSubjectRightsRequestSiteLocationImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawSubjectRightsRequestSiteLocationImpl) SubjectRightsRequestSiteLocation() BaseSubjectRightsRequestSiteLocationImpl {
	return s.subjectRightsRequestSiteLocation
}

func UnmarshalSubjectRightsRequestSiteLocationImplementation(input []byte) (SubjectRightsRequestSiteLocation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequestSiteLocation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestAllSiteLocation") {
		var out SubjectRightsRequestAllSiteLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestAllSiteLocation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.subjectRightsRequestEnumeratedSiteLocation") {
		var out SubjectRightsRequestEnumeratedSiteLocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubjectRightsRequestEnumeratedSiteLocation: %+v", err)
		}
		return out, nil
	}

	var parent BaseSubjectRightsRequestSiteLocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSubjectRightsRequestSiteLocationImpl: %+v", err)
	}

	return RawSubjectRightsRequestSiteLocationImpl{
		subjectRightsRequestSiteLocation: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
