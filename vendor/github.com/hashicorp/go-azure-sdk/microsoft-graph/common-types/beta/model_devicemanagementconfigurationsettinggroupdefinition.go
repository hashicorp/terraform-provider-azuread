package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingGroupDefinition interface {
	Entity
	DeviceManagementConfigurationSettingDefinition
	DeviceManagementConfigurationSettingGroupDefinition() BaseDeviceManagementConfigurationSettingGroupDefinitionImpl
}

var _ DeviceManagementConfigurationSettingGroupDefinition = BaseDeviceManagementConfigurationSettingGroupDefinitionImpl{}

type BaseDeviceManagementConfigurationSettingGroupDefinitionImpl struct {
	// Dependent child settings to this group of settings.
	ChildIds *[]string `json:"childIds,omitempty"`

	// List of child settings that depend on this setting
	DependedOnBy *[]DeviceManagementConfigurationSettingDependedOnBy `json:"dependedOnBy,omitempty"`

	// List of Dependencies for the setting group
	DependentOn *[]DeviceManagementConfigurationDependentOn `json:"dependentOn,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingDefinition

	AccessTypes *DeviceManagementConfigurationSettingAccessTypes `json:"accessTypes,omitempty"`

	// Details which device setting is applicable on
	Applicability DeviceManagementConfigurationSettingApplicability `json:"applicability"`

	// Base CSP Path
	BaseUri nullable.Type[string] `json:"baseUri,omitempty"`

	// Specifies the area group under which the setting is configured in a specified configuration service provider (CSP)
	CategoryId nullable.Type[string] `json:"categoryId,omitempty"`

	// Description of the item
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the item
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Help text of the item
	HelpText nullable.Type[string] `json:"helpText,omitempty"`

	// List of links more info for the setting can be found at
	InfoUrls *[]string `json:"infoUrls,omitempty"`

	// Tokens which to search settings on
	Keywords *[]string `json:"keywords,omitempty"`

	// Name of the item
	Name nullable.Type[string] `json:"name,omitempty"`

	// Indicates whether the setting is required or not
	Occurrence *DeviceManagementConfigurationSettingOccurrence `json:"occurrence,omitempty"`

	// Offset CSP Path from Base
	OffsetUri nullable.Type[string] `json:"offsetUri,omitempty"`

	// List of referred setting information.
	ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation `json:"referredSettingInformationList,omitempty"`

	// Root setting definition if the setting is a child setting.
	RootDefinitionId nullable.Type[string] `json:"rootDefinitionId,omitempty"`

	// Supported setting types
	SettingUsage *DeviceManagementConfigurationSettingUsage `json:"settingUsage,omitempty"`

	// Setting control type representation in the UX
	UxBehavior *DeviceManagementConfigurationControlType `json:"uxBehavior,omitempty"`

	// Item Version
	Version nullable.Type[string] `json:"version,omitempty"`

	// Supported setting types
	Visibility *DeviceManagementConfigurationSettingVisibility `json:"visibility,omitempty"`

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

func (s BaseDeviceManagementConfigurationSettingGroupDefinitionImpl) DeviceManagementConfigurationSettingGroupDefinition() BaseDeviceManagementConfigurationSettingGroupDefinitionImpl {
	return s
}

func (s BaseDeviceManagementConfigurationSettingGroupDefinitionImpl) DeviceManagementConfigurationSettingDefinition() BaseDeviceManagementConfigurationSettingDefinitionImpl {
	return BaseDeviceManagementConfigurationSettingDefinitionImpl{
		AccessTypes:                    s.AccessTypes,
		Applicability:                  s.Applicability,
		BaseUri:                        s.BaseUri,
		CategoryId:                     s.CategoryId,
		Description:                    s.Description,
		DisplayName:                    s.DisplayName,
		HelpText:                       s.HelpText,
		InfoUrls:                       s.InfoUrls,
		Keywords:                       s.Keywords,
		Name:                           s.Name,
		Occurrence:                     s.Occurrence,
		OffsetUri:                      s.OffsetUri,
		ReferredSettingInformationList: s.ReferredSettingInformationList,
		RootDefinitionId:               s.RootDefinitionId,
		SettingUsage:                   s.SettingUsage,
		UxBehavior:                     s.UxBehavior,
		Version:                        s.Version,
		Visibility:                     s.Visibility,
		Id:                             s.Id,
		ODataId:                        s.ODataId,
		ODataType:                      s.ODataType,
	}
}

func (s BaseDeviceManagementConfigurationSettingGroupDefinitionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementConfigurationSettingGroupDefinition = RawDeviceManagementConfigurationSettingGroupDefinitionImpl{}

// RawDeviceManagementConfigurationSettingGroupDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSettingGroupDefinitionImpl struct {
	deviceManagementConfigurationSettingGroupDefinition BaseDeviceManagementConfigurationSettingGroupDefinitionImpl
	Type                                                string
	Values                                              map[string]interface{}
}

func (s RawDeviceManagementConfigurationSettingGroupDefinitionImpl) DeviceManagementConfigurationSettingGroupDefinition() BaseDeviceManagementConfigurationSettingGroupDefinitionImpl {
	return s.deviceManagementConfigurationSettingGroupDefinition
}

func (s RawDeviceManagementConfigurationSettingGroupDefinitionImpl) DeviceManagementConfigurationSettingDefinition() BaseDeviceManagementConfigurationSettingDefinitionImpl {
	return s.deviceManagementConfigurationSettingGroupDefinition.DeviceManagementConfigurationSettingDefinition()
}

func (s RawDeviceManagementConfigurationSettingGroupDefinitionImpl) Entity() BaseEntityImpl {
	return s.deviceManagementConfigurationSettingGroupDefinition.Entity()
}

var _ json.Marshaler = BaseDeviceManagementConfigurationSettingGroupDefinitionImpl{}

func (s BaseDeviceManagementConfigurationSettingGroupDefinitionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementConfigurationSettingGroupDefinitionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementConfigurationSettingGroupDefinitionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementConfigurationSettingGroupDefinitionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSettingGroupDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementConfigurationSettingGroupDefinitionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseDeviceManagementConfigurationSettingGroupDefinitionImpl{}

func (s *BaseDeviceManagementConfigurationSettingGroupDefinitionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChildIds                       *[]string                                                  `json:"childIds,omitempty"`
		DependedOnBy                   *[]DeviceManagementConfigurationSettingDependedOnBy        `json:"dependedOnBy,omitempty"`
		DependentOn                    *[]DeviceManagementConfigurationDependentOn                `json:"dependentOn,omitempty"`
		AccessTypes                    *DeviceManagementConfigurationSettingAccessTypes           `json:"accessTypes,omitempty"`
		BaseUri                        nullable.Type[string]                                      `json:"baseUri,omitempty"`
		CategoryId                     nullable.Type[string]                                      `json:"categoryId,omitempty"`
		Description                    nullable.Type[string]                                      `json:"description,omitempty"`
		DisplayName                    nullable.Type[string]                                      `json:"displayName,omitempty"`
		HelpText                       nullable.Type[string]                                      `json:"helpText,omitempty"`
		InfoUrls                       *[]string                                                  `json:"infoUrls,omitempty"`
		Keywords                       *[]string                                                  `json:"keywords,omitempty"`
		Name                           nullable.Type[string]                                      `json:"name,omitempty"`
		Occurrence                     *DeviceManagementConfigurationSettingOccurrence            `json:"occurrence,omitempty"`
		OffsetUri                      nullable.Type[string]                                      `json:"offsetUri,omitempty"`
		ReferredSettingInformationList *[]DeviceManagementConfigurationReferredSettingInformation `json:"referredSettingInformationList,omitempty"`
		RootDefinitionId               nullable.Type[string]                                      `json:"rootDefinitionId,omitempty"`
		SettingUsage                   *DeviceManagementConfigurationSettingUsage                 `json:"settingUsage,omitempty"`
		UxBehavior                     *DeviceManagementConfigurationControlType                  `json:"uxBehavior,omitempty"`
		Version                        nullable.Type[string]                                      `json:"version,omitempty"`
		Visibility                     *DeviceManagementConfigurationSettingVisibility            `json:"visibility,omitempty"`
		Id                             *string                                                    `json:"id,omitempty"`
		ODataId                        *string                                                    `json:"@odata.id,omitempty"`
		ODataType                      *string                                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChildIds = decoded.ChildIds
	s.DependedOnBy = decoded.DependedOnBy
	s.DependentOn = decoded.DependentOn
	s.AccessTypes = decoded.AccessTypes
	s.BaseUri = decoded.BaseUri
	s.CategoryId = decoded.CategoryId
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.HelpText = decoded.HelpText
	s.Id = decoded.Id
	s.InfoUrls = decoded.InfoUrls
	s.Keywords = decoded.Keywords
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Occurrence = decoded.Occurrence
	s.OffsetUri = decoded.OffsetUri
	s.ReferredSettingInformationList = decoded.ReferredSettingInformationList
	s.RootDefinitionId = decoded.RootDefinitionId
	s.SettingUsage = decoded.SettingUsage
	s.UxBehavior = decoded.UxBehavior
	s.Version = decoded.Version
	s.Visibility = decoded.Visibility

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseDeviceManagementConfigurationSettingGroupDefinitionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["applicability"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSettingApplicabilityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Applicability' for 'BaseDeviceManagementConfigurationSettingGroupDefinitionImpl': %+v", err)
		}
		s.Applicability = impl
	}

	return nil
}

func UnmarshalDeviceManagementConfigurationSettingGroupDefinitionImplementation(input []byte) (DeviceManagementConfigurationSettingGroupDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingGroupDefinition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionDefinition") {
		var out DeviceManagementConfigurationSettingGroupCollectionDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingGroupCollectionDefinition: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSettingGroupDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSettingGroupDefinitionImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSettingGroupDefinitionImpl{
		deviceManagementConfigurationSettingGroupDefinition: parent,
		Type:   value,
		Values: temp,
	}, nil

}
