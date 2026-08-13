package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskCreation interface {
	PlannerTaskCreation() BasePlannerTaskCreationImpl
}

var _ PlannerTaskCreation = BasePlannerTaskCreationImpl{}

type BasePlannerTaskCreationImpl struct {
	// Specifies what kind of creation source the task is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information about the publication process that created this task. This field is deprecated and clients should move to
	// using the new inheritance model.
	TeamsPublicationInfo *PlannerTeamsPublicationInfo `json:"teamsPublicationInfo,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePlannerTaskCreationImpl) PlannerTaskCreation() BasePlannerTaskCreationImpl {
	return s
}

var _ PlannerTaskCreation = RawPlannerTaskCreationImpl{}

// RawPlannerTaskCreationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawPlannerTaskCreationImpl struct {
	plannerTaskCreation BasePlannerTaskCreationImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawPlannerTaskCreationImpl) PlannerTaskCreation() BasePlannerTaskCreationImpl {
	return s.plannerTaskCreation
}

func (s RawPlannerTaskCreationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalPlannerTaskCreationImplementation(input []byte) (PlannerTaskCreation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerTaskCreation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerExternalTaskSource") {
		var out PlannerExternalTaskSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerExternalTaskSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTeamsPublicationInfo") {
		var out PlannerTeamsPublicationInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTeamsPublicationInfo: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerTaskCreationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerTaskCreationImpl: %+v", err)
	}

	return RawPlannerTaskCreationImpl{
		plannerTaskCreation: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
