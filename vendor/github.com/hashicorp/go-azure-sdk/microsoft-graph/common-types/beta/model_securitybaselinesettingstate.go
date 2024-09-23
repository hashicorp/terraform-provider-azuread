package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityBaselineSettingState{}

type SecurityBaselineSettingState struct {
	// The policies that contribute to this setting instance
	ContributingPolicies *[]SecurityBaselineContributingPolicy `json:"contributingPolicies,omitempty"`

	// The error code if the setting is in error state
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// The setting category id which this setting belongs to
	SettingCategoryId nullable.Type[string] `json:"settingCategoryId,omitempty"`

	// The setting category name which this setting belongs to
	SettingCategoryName nullable.Type[string] `json:"settingCategoryName,omitempty"`

	// The setting id guid
	SettingId nullable.Type[string] `json:"settingId,omitempty"`

	// The setting name that is being reported
	SettingName *string `json:"settingName,omitempty"`

	// The policies that contribute to this setting instance
	SourcePolicies *[]SettingSource `json:"sourcePolicies,omitempty"`

	// Security Baseline Compliance State
	State *SecurityBaselineComplianceState `json:"state,omitempty"`

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

func (s SecurityBaselineSettingState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityBaselineSettingState{}

func (s SecurityBaselineSettingState) MarshalJSON() ([]byte, error) {
	type wrapper SecurityBaselineSettingState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityBaselineSettingState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityBaselineSettingState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityBaselineSettingState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityBaselineSettingState: %+v", err)
	}

	return encoded, nil
}
