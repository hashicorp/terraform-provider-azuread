package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceInsight interface {
	Entity
	GovernanceInsight() BaseGovernanceInsightImpl
}

var _ GovernanceInsight = BaseGovernanceInsightImpl{}

type BaseGovernanceInsightImpl struct {
	// Indicates when the insight was created.
	InsightCreatedDateTime nullable.Type[string] `json:"insightCreatedDateTime,omitempty"`

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

func (s BaseGovernanceInsightImpl) GovernanceInsight() BaseGovernanceInsightImpl {
	return s
}

func (s BaseGovernanceInsightImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GovernanceInsight = RawGovernanceInsightImpl{}

// RawGovernanceInsightImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGovernanceInsightImpl struct {
	governanceInsight BaseGovernanceInsightImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawGovernanceInsightImpl) GovernanceInsight() BaseGovernanceInsightImpl {
	return s.governanceInsight
}

func (s RawGovernanceInsightImpl) Entity() BaseEntityImpl {
	return s.governanceInsight.Entity()
}

var _ json.Marshaler = BaseGovernanceInsightImpl{}

func (s BaseGovernanceInsightImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGovernanceInsightImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGovernanceInsightImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGovernanceInsightImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceInsight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGovernanceInsightImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalGovernanceInsightImplementation(input []byte) (GovernanceInsight, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceInsight into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.membershipOutlierInsight") {
		var out MembershipOutlierInsight
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MembershipOutlierInsight: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSignInInsight") {
		var out UserSignInInsight
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSignInInsight: %+v", err)
		}
		return out, nil
	}

	var parent BaseGovernanceInsightImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGovernanceInsightImpl: %+v", err)
	}

	return RawGovernanceInsightImpl{
		governanceInsight: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
