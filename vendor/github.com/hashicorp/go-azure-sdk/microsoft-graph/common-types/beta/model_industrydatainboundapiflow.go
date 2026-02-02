package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataInboundFlow = IndustryDataInboundApiFlow{}

type IndustryDataInboundApiFlow struct {

	// Fields inherited from IndustryDataInboundFlow

	DataConnector *IndustryDataIndustryDataConnector `json:"dataConnector,omitempty"`
	DataDomain    *IndustryDataInboundDomain         `json:"dataDomain,omitempty"`

	// The start of the time window when the flow is allowed to run. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`

	// The end of the time window when the flow is allowed to run. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	Year *IndustryDataYearTimePeriodDefinition `json:"year,omitempty"`

	// Fields inherited from IndustryDataIndustryDataActivity

	// The name of the activity. Maximum supported length is 100 characters.
	DisplayName *string `json:"displayName,omitempty"`

	ReadinessStatus *IndustryDataReadinessStatus `json:"readinessStatus,omitempty"`

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

func (s IndustryDataInboundApiFlow) IndustryDataInboundFlow() BaseIndustryDataInboundFlowImpl {
	return BaseIndustryDataInboundFlowImpl{
		DataConnector:      s.DataConnector,
		DataDomain:         s.DataDomain,
		EffectiveDateTime:  s.EffectiveDateTime,
		ExpirationDateTime: s.ExpirationDateTime,
		Year:               s.Year,
		DisplayName:        s.DisplayName,
		ReadinessStatus:    s.ReadinessStatus,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s IndustryDataInboundApiFlow) IndustryDataIndustryDataActivity() BaseIndustryDataIndustryDataActivityImpl {
	return BaseIndustryDataIndustryDataActivityImpl{
		DisplayName:     s.DisplayName,
		ReadinessStatus: s.ReadinessStatus,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s IndustryDataInboundApiFlow) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataInboundApiFlow{}

func (s IndustryDataInboundApiFlow) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataInboundApiFlow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataInboundApiFlow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataInboundApiFlow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.inboundApiFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataInboundApiFlow: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataInboundApiFlow{}

func (s *IndustryDataInboundApiFlow) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DataDomain         *IndustryDataInboundDomain            `json:"dataDomain,omitempty"`
		EffectiveDateTime  *string                               `json:"effectiveDateTime,omitempty"`
		ExpirationDateTime nullable.Type[string]                 `json:"expirationDateTime,omitempty"`
		Year               *IndustryDataYearTimePeriodDefinition `json:"year,omitempty"`
		DisplayName        *string                               `json:"displayName,omitempty"`
		ReadinessStatus    *IndustryDataReadinessStatus          `json:"readinessStatus,omitempty"`
		Id                 *string                               `json:"id,omitempty"`
		ODataId            *string                               `json:"@odata.id,omitempty"`
		ODataType          *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataDomain = decoded.DataDomain
	s.DisplayName = decoded.DisplayName
	s.EffectiveDateTime = decoded.EffectiveDateTime
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ReadinessStatus = decoded.ReadinessStatus
	s.Year = decoded.Year

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataInboundApiFlow into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataConnector"]; ok {
		impl, err := UnmarshalIndustryDataIndustryDataConnectorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataConnector' for 'IndustryDataInboundApiFlow': %+v", err)
		}
		s.DataConnector = &impl
	}

	return nil
}
