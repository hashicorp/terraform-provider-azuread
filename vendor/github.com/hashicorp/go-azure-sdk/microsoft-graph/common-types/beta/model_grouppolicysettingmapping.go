package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicySettingMapping{}

type GroupPolicySettingMapping struct {
	// Admx Group Policy Id
	AdmxSettingDefinitionId nullable.Type[string] `json:"admxSettingDefinitionId,omitempty"`

	// List of Child Ids of the group policy setting.
	ChildIdList *[]string `json:"childIdList,omitempty"`

	// The Intune Setting Definition Id
	IntuneSettingDefinitionId nullable.Type[string] `json:"intuneSettingDefinitionId,omitempty"`

	// The list of Intune Setting URIs this group policy setting maps to
	IntuneSettingUriList *[]string `json:"intuneSettingUriList,omitempty"`

	// Indicates if the setting is supported by Intune or not
	IsMdmSupported *bool `json:"isMdmSupported,omitempty"`

	// The CSP name this group policy setting maps to.
	MdmCspName nullable.Type[string] `json:"mdmCspName,omitempty"`

	// The minimum OS version this mdm setting supports.
	MdmMinimumOSVersion *int64 `json:"mdmMinimumOSVersion,omitempty"`

	// The MDM CSP URI this group policy setting maps to.
	MdmSettingUri nullable.Type[string] `json:"mdmSettingUri,omitempty"`

	// Mdm Support Status of the setting.
	MdmSupportedState *MdmSupportedState `json:"mdmSupportedState,omitempty"`

	// Parent Id of the group policy setting.
	ParentId nullable.Type[string] `json:"parentId,omitempty"`

	// The category the group policy setting is in.
	SettingCategory nullable.Type[string] `json:"settingCategory,omitempty"`

	// The display name of this group policy setting.
	SettingDisplayName nullable.Type[string] `json:"settingDisplayName,omitempty"`

	// The display value of this group policy setting.
	SettingDisplayValue nullable.Type[string] `json:"settingDisplayValue,omitempty"`

	// The display value type of this group policy setting.
	SettingDisplayValueType nullable.Type[string] `json:"settingDisplayValueType,omitempty"`

	// The name of this group policy setting.
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// Scope of the group policy setting.
	SettingScope *GroupPolicySettingScope `json:"settingScope,omitempty"`

	// Setting type of the group policy.
	SettingType *GroupPolicySettingType `json:"settingType,omitempty"`

	// The value of this group policy setting.
	SettingValue nullable.Type[string] `json:"settingValue,omitempty"`

	// The display units of this group policy setting value
	SettingValueDisplayUnits nullable.Type[string] `json:"settingValueDisplayUnits,omitempty"`

	// The value type of this group policy setting.
	SettingValueType nullable.Type[string] `json:"settingValueType,omitempty"`

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

func (s GroupPolicySettingMapping) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicySettingMapping{}

func (s GroupPolicySettingMapping) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicySettingMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicySettingMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicySettingMapping: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicySettingMapping"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicySettingMapping: %+v", err)
	}

	return encoded, nil
}
