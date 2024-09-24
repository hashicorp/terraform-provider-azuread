package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementTemplate = SecurityBaselineTemplate{}

type SecurityBaselineTemplate struct {
	// The security baseline per category device state summary
	CategoryDeviceStateSummaries *[]SecurityBaselineCategoryStateSummary `json:"categoryDeviceStateSummaries,omitempty"`

	// The security baseline device state summary
	DeviceStateSummary *SecurityBaselineStateSummary `json:"deviceStateSummary,omitempty"`

	// The security baseline device states
	DeviceStates *[]SecurityBaselineDeviceState `json:"deviceStates,omitempty"`

	// Fields inherited from DeviceManagementTemplate

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

func (s SecurityBaselineTemplate) DeviceManagementTemplate() BaseDeviceManagementTemplateImpl {
	return BaseDeviceManagementTemplateImpl{
		Categories:        s.Categories,
		Description:       s.Description,
		DisplayName:       s.DisplayName,
		IntentCount:       s.IntentCount,
		IsDeprecated:      s.IsDeprecated,
		MigratableTo:      s.MigratableTo,
		PlatformType:      s.PlatformType,
		PublishedDateTime: s.PublishedDateTime,
		Settings:          s.Settings,
		TemplateSubtype:   s.TemplateSubtype,
		TemplateType:      s.TemplateType,
		VersionInfo:       s.VersionInfo,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s SecurityBaselineTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityBaselineTemplate{}

func (s SecurityBaselineTemplate) MarshalJSON() ([]byte, error) {
	type wrapper SecurityBaselineTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityBaselineTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityBaselineTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityBaselineTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityBaselineTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityBaselineTemplate{}

func (s *SecurityBaselineTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CategoryDeviceStateSummaries *[]SecurityBaselineCategoryStateSummary    `json:"categoryDeviceStateSummaries,omitempty"`
		DeviceStates                 *[]SecurityBaselineDeviceState             `json:"deviceStates,omitempty"`
		Categories                   *[]DeviceManagementTemplateSettingCategory `json:"categories,omitempty"`
		Description                  nullable.Type[string]                      `json:"description,omitempty"`
		DisplayName                  *string                                    `json:"displayName,omitempty"`
		IntentCount                  *int64                                     `json:"intentCount,omitempty"`
		IsDeprecated                 *bool                                      `json:"isDeprecated,omitempty"`
		PlatformType                 *PolicyPlatformType                        `json:"platformType,omitempty"`
		PublishedDateTime            *string                                    `json:"publishedDateTime,omitempty"`
		TemplateSubtype              *DeviceManagementTemplateSubtype           `json:"templateSubtype,omitempty"`
		TemplateType                 *DeviceManagementTemplateType              `json:"templateType,omitempty"`
		VersionInfo                  nullable.Type[string]                      `json:"versionInfo,omitempty"`
		Id                           *string                                    `json:"id,omitempty"`
		ODataId                      *string                                    `json:"@odata.id,omitempty"`
		ODataType                    *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CategoryDeviceStateSummaries = decoded.CategoryDeviceStateSummaries
	s.DeviceStates = decoded.DeviceStates
	s.Categories = decoded.Categories
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.IntentCount = decoded.IntentCount
	s.IsDeprecated = decoded.IsDeprecated
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PlatformType = decoded.PlatformType
	s.PublishedDateTime = decoded.PublishedDateTime
	s.TemplateSubtype = decoded.TemplateSubtype
	s.TemplateType = decoded.TemplateType
	s.VersionInfo = decoded.VersionInfo

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityBaselineTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deviceStateSummary"]; ok {
		impl, err := UnmarshalSecurityBaselineStateSummaryImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DeviceStateSummary' for 'SecurityBaselineTemplate': %+v", err)
		}
		s.DeviceStateSummary = &impl
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
				return fmt.Errorf("unmarshaling index %d field 'MigratableTo' for 'SecurityBaselineTemplate': %+v", i, err)
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
				return fmt.Errorf("unmarshaling index %d field 'Settings' for 'SecurityBaselineTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Settings = &output
	}

	return nil
}
