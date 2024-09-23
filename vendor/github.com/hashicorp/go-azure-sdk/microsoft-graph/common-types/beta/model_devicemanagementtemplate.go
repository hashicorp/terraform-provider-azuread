package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplate interface {
	Entity
	DeviceManagementTemplate() BaseDeviceManagementTemplateImpl
}

var _ DeviceManagementTemplate = BaseDeviceManagementTemplateImpl{}

type BaseDeviceManagementTemplateImpl struct {
	// Collection of setting categories within the template
	Categories *[]DeviceManagementTemplateSettingCategory `json:"categories,omitempty"`

	// The template's description
	Description nullable.Type[string] `json:"description,omitempty"`

	// The template's display name
	DisplayName *string `json:"displayName,omitempty"`

	// Number of Intents created from this template.
	IntentCount *int64 `json:"intentCount,omitempty"`

	// The template is deprecated or not. Intents cannot be created from a deprecated template.
	IsDeprecated *bool `json:"isDeprecated,omitempty"`

	// Collection of templates this template can migrate to
	MigratableTo *[]DeviceManagementTemplate `json:"migratableTo,omitempty"`

	// Supported platform types for policies.
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`

	// When the template was published
	PublishedDateTime *string `json:"publishedDateTime,omitempty"`

	// Collection of all settings this template has
	Settings *[]DeviceManagementSettingInstance `json:"settings,omitempty"`

	// Template subtype
	TemplateSubtype *DeviceManagementTemplateSubtype `json:"templateSubtype,omitempty"`

	// Template type
	TemplateType *DeviceManagementTemplateType `json:"templateType,omitempty"`

	// The template's version information
	VersionInfo nullable.Type[string] `json:"versionInfo,omitempty"`

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

func (s BaseDeviceManagementTemplateImpl) DeviceManagementTemplate() BaseDeviceManagementTemplateImpl {
	return s
}

func (s BaseDeviceManagementTemplateImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementTemplate = RawDeviceManagementTemplateImpl{}

// RawDeviceManagementTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementTemplateImpl struct {
	deviceManagementTemplate BaseDeviceManagementTemplateImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawDeviceManagementTemplateImpl) DeviceManagementTemplate() BaseDeviceManagementTemplateImpl {
	return s.deviceManagementTemplate
}

func (s RawDeviceManagementTemplateImpl) Entity() BaseEntityImpl {
	return s.deviceManagementTemplate.Entity()
}

var _ json.Marshaler = BaseDeviceManagementTemplateImpl{}

func (s BaseDeviceManagementTemplateImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementTemplateImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementTemplateImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementTemplateImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementTemplateImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseDeviceManagementTemplateImpl{}

func (s *BaseDeviceManagementTemplateImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Categories        *[]DeviceManagementTemplateSettingCategory `json:"categories,omitempty"`
		Description       nullable.Type[string]                      `json:"description,omitempty"`
		DisplayName       *string                                    `json:"displayName,omitempty"`
		IntentCount       *int64                                     `json:"intentCount,omitempty"`
		IsDeprecated      *bool                                      `json:"isDeprecated,omitempty"`
		PlatformType      *PolicyPlatformType                        `json:"platformType,omitempty"`
		PublishedDateTime *string                                    `json:"publishedDateTime,omitempty"`
		TemplateSubtype   *DeviceManagementTemplateSubtype           `json:"templateSubtype,omitempty"`
		TemplateType      *DeviceManagementTemplateType              `json:"templateType,omitempty"`
		VersionInfo       nullable.Type[string]                      `json:"versionInfo,omitempty"`
		Id                *string                                    `json:"id,omitempty"`
		ODataId           *string                                    `json:"@odata.id,omitempty"`
		ODataType         *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Categories = decoded.Categories
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IntentCount = decoded.IntentCount
	s.IsDeprecated = decoded.IsDeprecated
	s.PlatformType = decoded.PlatformType
	s.PublishedDateTime = decoded.PublishedDateTime
	s.TemplateSubtype = decoded.TemplateSubtype
	s.TemplateType = decoded.TemplateType
	s.VersionInfo = decoded.VersionInfo
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseDeviceManagementTemplateImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["migratableTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MigratableTo into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementTemplate, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementTemplateImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MigratableTo' for 'BaseDeviceManagementTemplateImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MigratableTo = &output
	}

	if v, ok := temp["settings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Settings into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingInstance, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingInstanceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Settings' for 'BaseDeviceManagementTemplateImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Settings = &output
	}

	return nil
}

func UnmarshalDeviceManagementTemplateImplementation(input []byte) (DeviceManagementTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.securityBaselineTemplate") {
		var out SecurityBaselineTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBaselineTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementTemplateImpl: %+v", err)
	}

	return RawDeviceManagementTemplateImpl{
		deviceManagementTemplate: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
