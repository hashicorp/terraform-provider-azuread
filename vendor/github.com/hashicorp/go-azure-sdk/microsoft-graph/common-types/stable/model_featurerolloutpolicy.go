package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = FeatureRolloutPolicy{}

type FeatureRolloutPolicy struct {
	// Nullable. Specifies a list of directoryObject resources that feature is enabled for.
	AppliesTo *[]DirectoryObject `json:"appliesTo,omitempty"`

	// List of OData IDs for `AppliesTo` to bind to this entity
	AppliesTo_ODataBind *[]string `json:"appliesTo@odata.bind,omitempty"`

	// A description for this feature rollout policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for this feature rollout policy.
	DisplayName *string `json:"displayName,omitempty"`

	Feature *StagedFeatureName `json:"feature,omitempty"`

	// Indicates whether this feature rollout policy should be applied to the entire organization.
	IsAppliedToOrganization *bool `json:"isAppliedToOrganization,omitempty"`

	// Indicates whether the feature rollout is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`

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

func (s FeatureRolloutPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FeatureRolloutPolicy{}

func (s FeatureRolloutPolicy) MarshalJSON() ([]byte, error) {
	type wrapper FeatureRolloutPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FeatureRolloutPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureRolloutPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.featureRolloutPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FeatureRolloutPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &FeatureRolloutPolicy{}

func (s *FeatureRolloutPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppliesTo_ODataBind     *[]string             `json:"appliesTo@odata.bind,omitempty"`
		Description             nullable.Type[string] `json:"description,omitempty"`
		DisplayName             *string               `json:"displayName,omitempty"`
		Feature                 *StagedFeatureName    `json:"feature,omitempty"`
		IsAppliedToOrganization *bool                 `json:"isAppliedToOrganization,omitempty"`
		IsEnabled               *bool                 `json:"isEnabled,omitempty"`
		Id                      *string               `json:"id,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppliesTo_ODataBind = decoded.AppliesTo_ODataBind
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Feature = decoded.Feature
	s.IsAppliedToOrganization = decoded.IsAppliedToOrganization
	s.IsEnabled = decoded.IsEnabled
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FeatureRolloutPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appliesTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppliesTo into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppliesTo' for 'FeatureRolloutPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppliesTo = &output
	}

	return nil
}
