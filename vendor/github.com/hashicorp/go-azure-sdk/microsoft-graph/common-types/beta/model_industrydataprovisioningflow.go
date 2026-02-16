package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataProvisioningFlow interface {
	Entity
	IndustryDataProvisioningFlow() BaseIndustryDataProvisioningFlowImpl
}

var _ IndustryDataProvisioningFlow = BaseIndustryDataProvisioningFlowImpl{}

type BaseIndustryDataProvisioningFlowImpl struct {
	// The date and time when the provisioning flow was created. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time when the provisioning flow was most recently changed. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The state of the activity from creation through to ready to do work. The possible values are: notReady, ready,
	// failed, disabled, expired, unknownFutureValue.
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

func (s BaseIndustryDataProvisioningFlowImpl) IndustryDataProvisioningFlow() BaseIndustryDataProvisioningFlowImpl {
	return s
}

func (s BaseIndustryDataProvisioningFlowImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataProvisioningFlow = RawIndustryDataProvisioningFlowImpl{}

// RawIndustryDataProvisioningFlowImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataProvisioningFlowImpl struct {
	industryDataProvisioningFlow BaseIndustryDataProvisioningFlowImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawIndustryDataProvisioningFlowImpl) IndustryDataProvisioningFlow() BaseIndustryDataProvisioningFlowImpl {
	return s.industryDataProvisioningFlow
}

func (s RawIndustryDataProvisioningFlowImpl) Entity() BaseEntityImpl {
	return s.industryDataProvisioningFlow.Entity()
}

var _ json.Marshaler = BaseIndustryDataProvisioningFlowImpl{}

func (s BaseIndustryDataProvisioningFlowImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataProvisioningFlowImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataProvisioningFlowImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataProvisioningFlowImpl: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "readinessStatus")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.provisioningFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataProvisioningFlowImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataProvisioningFlowImplementation(input []byte) (IndustryDataProvisioningFlow, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataProvisioningFlow into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.administrativeUnitProvisioningFlow") {
		var out IndustryDataAdministrativeUnitProvisioningFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataAdministrativeUnitProvisioningFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.classGroupProvisioningFlow") {
		var out IndustryDataClassGroupProvisioningFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataClassGroupProvisioningFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.securityGroupProvisioningFlow") {
		var out IndustryDataSecurityGroupProvisioningFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataSecurityGroupProvisioningFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.userProvisioningFlow") {
		var out IndustryDataUserProvisioningFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataUserProvisioningFlow: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataProvisioningFlowImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataProvisioningFlowImpl: %+v", err)
	}

	return RawIndustryDataProvisioningFlowImpl{
		industryDataProvisioningFlow: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
