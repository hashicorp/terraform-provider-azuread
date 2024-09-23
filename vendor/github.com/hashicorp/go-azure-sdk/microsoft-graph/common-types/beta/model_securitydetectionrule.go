package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityProtectionRule = SecurityDetectionRule{}

type SecurityDetectionRule struct {
	// Complex type representing the actions taken when a detection is made by this rule.
	DetectionAction *SecurityDetectionAction `json:"detectionAction,omitempty"`

	// The ID of the detector that triggered the alert. Also see the 'detectorId' field in microsoft.graph.security.alert.
	DetectorId nullable.Type[string] `json:"detectorId,omitempty"`

	// Complex type holding details about the last run of this rule.
	LastRunDetails *SecurityRunDetails `json:"lastRunDetails,omitempty"`

	// Complex type holding data about the advanced hunting query of this rule.
	QueryCondition *SecurityQueryCondition `json:"queryCondition,omitempty"`

	// Complex type holding data about the triggering schedule of this rule.
	Schedule *SecurityRuleSchedule `json:"schedule,omitempty"`

	// Fields inherited from SecurityProtectionRule

	// Name of the user or application that created the rule.
	CreatedBy *string `json:"createdBy,omitempty"`

	// Timestamp of rule creation.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Name of the rule.
	DisplayName *string `json:"displayName,omitempty"`

	// Whether rule is turned on for the tenant.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Name of the user or application who last updated the rule.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// Timestamp of when the rule was last updated.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s SecurityDetectionRule) SecurityProtectionRule() BaseSecurityProtectionRuleImpl {
	return BaseSecurityProtectionRuleImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		IsEnabled:            s.IsEnabled,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityDetectionRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityDetectionRule{}

func (s SecurityDetectionRule) MarshalJSON() ([]byte, error) {
	type wrapper SecurityDetectionRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityDetectionRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDetectionRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.detectionRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityDetectionRule: %+v", err)
	}

	return encoded, nil
}
