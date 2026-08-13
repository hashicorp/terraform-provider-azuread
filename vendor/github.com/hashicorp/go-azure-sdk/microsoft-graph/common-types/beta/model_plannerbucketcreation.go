package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerBucketCreation interface {
	PlannerBucketCreation() BasePlannerBucketCreationImpl
}

var _ PlannerBucketCreation = BasePlannerBucketCreationImpl{}

type BasePlannerBucketCreationImpl struct {
	// Specifies what kind of creation source the bucket is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePlannerBucketCreationImpl) PlannerBucketCreation() BasePlannerBucketCreationImpl {
	return s
}

var _ PlannerBucketCreation = RawPlannerBucketCreationImpl{}

// RawPlannerBucketCreationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawPlannerBucketCreationImpl struct {
	plannerBucketCreation BasePlannerBucketCreationImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawPlannerBucketCreationImpl) PlannerBucketCreation() BasePlannerBucketCreationImpl {
	return s.plannerBucketCreation
}

func (s RawPlannerBucketCreationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalPlannerBucketCreationImplementation(input []byte) (PlannerBucketCreation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerBucketCreation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerExternalBucketSource") {
		var out PlannerExternalBucketSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerExternalBucketSource: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerBucketCreationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerBucketCreationImpl: %+v", err)
	}

	return RawPlannerBucketCreationImpl{
		plannerBucketCreation: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
