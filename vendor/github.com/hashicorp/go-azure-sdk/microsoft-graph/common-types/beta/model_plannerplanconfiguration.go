package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerPlanConfiguration{}

type PlannerPlanConfiguration struct {
	// List the buckets that should be created in the plan.
	Buckets *[]PlannerPlanConfigurationBucketDefinition `json:"buckets,omitempty"`

	// The identity of the creator of the plan configuration.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time when the plan configuration was created. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The language code for the default language to be used for the names of the objects created for the plan.
	DefaultLanguage nullable.Type[string] `json:"defaultLanguage,omitempty"`

	// The identity of the user who last modified the plan configuration.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the plan configuration was last modified. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Localized names for the plan configuration.
	Localizations *[]PlannerPlanConfigurationLocalization `json:"localizations,omitempty"`

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

func (s PlannerPlanConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerPlanConfiguration{}

func (s PlannerPlanConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper PlannerPlanConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerPlanConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlanConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerPlanConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerPlanConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerPlanConfiguration{}

func (s *PlannerPlanConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Buckets              *[]PlannerPlanConfigurationBucketDefinition `json:"buckets,omitempty"`
		CreatedDateTime      *string                                     `json:"createdDateTime,omitempty"`
		DefaultLanguage      nullable.Type[string]                       `json:"defaultLanguage,omitempty"`
		LastModifiedDateTime *string                                     `json:"lastModifiedDateTime,omitempty"`
		Localizations        *[]PlannerPlanConfigurationLocalization     `json:"localizations,omitempty"`
		Id                   *string                                     `json:"id,omitempty"`
		ODataId              *string                                     `json:"@odata.id,omitempty"`
		ODataType            *string                                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Buckets = decoded.Buckets
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefaultLanguage = decoded.DefaultLanguage
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Localizations = decoded.Localizations
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerPlanConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PlannerPlanConfiguration': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'PlannerPlanConfiguration': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
