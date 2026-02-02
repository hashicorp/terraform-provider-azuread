package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequest interface {
	Entity
	ThreatAssessmentRequest() BaseThreatAssessmentRequestImpl
}

var _ ThreatAssessmentRequest = BaseThreatAssessmentRequestImpl{}

type BaseThreatAssessmentRequestImpl struct {
	Category *ThreatCategory `json:"category,omitempty"`

	// The content type of threat assessment. Possible values are: mail, url, file.
	ContentType *ThreatAssessmentContentType `json:"contentType,omitempty"`

	// The threat assessment request creator.
	CreatedBy IdentitySet `json:"createdBy"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	ExpectedAssessment *ThreatExpectedAssessment `json:"expectedAssessment,omitempty"`

	// The source of the threat assessment request. Possible values are: user, administrator.
	RequestSource *ThreatAssessmentRequestSource `json:"requestSource,omitempty"`

	// A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not
	// return this property unless you apply $expand on it.
	Results *[]ThreatAssessmentResult `json:"results,omitempty"`

	// The assessment process status. Possible values are: pending, completed.
	Status *ThreatAssessmentStatus `json:"status,omitempty"`

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

func (s BaseThreatAssessmentRequestImpl) ThreatAssessmentRequest() BaseThreatAssessmentRequestImpl {
	return s
}

func (s BaseThreatAssessmentRequestImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ThreatAssessmentRequest = RawThreatAssessmentRequestImpl{}

// RawThreatAssessmentRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawThreatAssessmentRequestImpl struct {
	threatAssessmentRequest BaseThreatAssessmentRequestImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawThreatAssessmentRequestImpl) ThreatAssessmentRequest() BaseThreatAssessmentRequestImpl {
	return s.threatAssessmentRequest
}

func (s RawThreatAssessmentRequestImpl) Entity() BaseEntityImpl {
	return s.threatAssessmentRequest.Entity()
}

var _ json.Marshaler = BaseThreatAssessmentRequestImpl{}

func (s BaseThreatAssessmentRequestImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseThreatAssessmentRequestImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseThreatAssessmentRequestImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseThreatAssessmentRequestImpl: %+v", err)
	}

	delete(decoded, "results")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.threatAssessmentRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseThreatAssessmentRequestImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseThreatAssessmentRequestImpl{}

func (s *BaseThreatAssessmentRequestImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Category           *ThreatCategory                `json:"category,omitempty"`
		ContentType        *ThreatAssessmentContentType   `json:"contentType,omitempty"`
		CreatedDateTime    nullable.Type[string]          `json:"createdDateTime,omitempty"`
		ExpectedAssessment *ThreatExpectedAssessment      `json:"expectedAssessment,omitempty"`
		RequestSource      *ThreatAssessmentRequestSource `json:"requestSource,omitempty"`
		Results            *[]ThreatAssessmentResult      `json:"results,omitempty"`
		Status             *ThreatAssessmentStatus        `json:"status,omitempty"`
		Id                 *string                        `json:"id,omitempty"`
		ODataId            *string                        `json:"@odata.id,omitempty"`
		ODataType          *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Category = decoded.Category
	s.ContentType = decoded.ContentType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.ExpectedAssessment = decoded.ExpectedAssessment
	s.RequestSource = decoded.RequestSource
	s.Results = decoded.Results
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseThreatAssessmentRequestImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseThreatAssessmentRequestImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalThreatAssessmentRequestImplementation(input []byte) (ThreatAssessmentRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ThreatAssessmentRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.emailFileAssessmentRequest") {
		var out EmailFileAssessmentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailFileAssessmentRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileAssessmentRequest") {
		var out FileAssessmentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileAssessmentRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailAssessmentRequest") {
		var out MailAssessmentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailAssessmentRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.urlAssessmentRequest") {
		var out UrlAssessmentRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UrlAssessmentRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseThreatAssessmentRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseThreatAssessmentRequestImpl: %+v", err)
	}

	return RawThreatAssessmentRequestImpl{
		threatAssessmentRequest: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
