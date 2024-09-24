package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicySetItem = EnrollmentRestrictionsConfigurationPolicySetItem{}

type EnrollmentRestrictionsConfigurationPolicySetItem struct {
	// Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
	Limit nullable.Type[int64] `json:"limit,omitempty"`

	// Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// Fields inherited from PolicySetItem

	// Creation time of the PolicySetItem.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// DisplayName of the PolicySetItem.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	ErrorCode *ErrorCode `json:"errorCode,omitempty"`

	// Tags of the guided deployment
	GuidedDeploymentTags *[]string `json:"guidedDeploymentTags,omitempty"`

	// policySetType of the PolicySetItem.
	ItemType nullable.Type[string] `json:"itemType,omitempty"`

	// Last modified time of the PolicySetItem.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// PayloadId of the PolicySetItem.
	PayloadId *string `json:"payloadId,omitempty"`

	// The enum to specify the status of PolicySet.
	Status *PolicySetStatus `json:"status,omitempty"`

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

func (s EnrollmentRestrictionsConfigurationPolicySetItem) PolicySetItem() BasePolicySetItemImpl {
	return BasePolicySetItemImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		ErrorCode:            s.ErrorCode,
		GuidedDeploymentTags: s.GuidedDeploymentTags,
		ItemType:             s.ItemType,
		LastModifiedDateTime: s.LastModifiedDateTime,
		PayloadId:            s.PayloadId,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s EnrollmentRestrictionsConfigurationPolicySetItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnrollmentRestrictionsConfigurationPolicySetItem{}

func (s EnrollmentRestrictionsConfigurationPolicySetItem) MarshalJSON() ([]byte, error) {
	type wrapper EnrollmentRestrictionsConfigurationPolicySetItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnrollmentRestrictionsConfigurationPolicySetItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnrollmentRestrictionsConfigurationPolicySetItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enrollmentRestrictionsConfigurationPolicySetItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnrollmentRestrictionsConfigurationPolicySetItem: %+v", err)
	}

	return encoded, nil
}
