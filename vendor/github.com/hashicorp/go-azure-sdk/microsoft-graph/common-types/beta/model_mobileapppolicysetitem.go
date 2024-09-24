package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicySetItem = MobileAppPolicySetItem{}

type MobileAppPolicySetItem struct {
	// Possible values for the install intent chosen by the admin.
	Intent *InstallIntent `json:"intent,omitempty"`

	// Settings of the MobileAppPolicySetItem.
	Settings MobileAppAssignmentSettings `json:"settings"`

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

func (s MobileAppPolicySetItem) PolicySetItem() BasePolicySetItemImpl {
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

func (s MobileAppPolicySetItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppPolicySetItem{}

func (s MobileAppPolicySetItem) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppPolicySetItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppPolicySetItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppPolicySetItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppPolicySetItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppPolicySetItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MobileAppPolicySetItem{}

func (s *MobileAppPolicySetItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Intent               *InstallIntent        `json:"intent,omitempty"`
		CreatedDateTime      *string               `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		ErrorCode            *ErrorCode            `json:"errorCode,omitempty"`
		GuidedDeploymentTags *[]string             `json:"guidedDeploymentTags,omitempty"`
		ItemType             nullable.Type[string] `json:"itemType,omitempty"`
		LastModifiedDateTime *string               `json:"lastModifiedDateTime,omitempty"`
		PayloadId            *string               `json:"payloadId,omitempty"`
		Status               *PolicySetStatus      `json:"status,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Intent = decoded.Intent
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.ErrorCode = decoded.ErrorCode
	s.GuidedDeploymentTags = decoded.GuidedDeploymentTags
	s.Id = decoded.Id
	s.ItemType = decoded.ItemType
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PayloadId = decoded.PayloadId
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MobileAppPolicySetItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settings"]; ok {
		impl, err := UnmarshalMobileAppAssignmentSettingsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Settings' for 'MobileAppPolicySetItem': %+v", err)
		}
		s.Settings = impl
	}

	return nil
}
