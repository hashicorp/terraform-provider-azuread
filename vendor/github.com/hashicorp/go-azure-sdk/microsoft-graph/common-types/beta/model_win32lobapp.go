package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobApp interface {
	Entity
	MobileApp
	MobileLobApp
	Win32LobApp() BaseWin32LobAppImpl
}

var _ Win32LobApp = BaseWin32LobAppImpl{}

type BaseWin32LobAppImpl struct {
	// Indicates whether the uninstall is supported from the company portal for the Win32 app with an available assignment.
	// When TRUE, indicates that uninstall is supported from the company portal for the Windows app (Win32) with an
	// available assignment. When FALSE, indicates that uninstall is not supported for the Windows app (Win32) with an
	// Available assignment. Default value is FALSE.
	AllowAvailableUninstall *bool `json:"allowAvailableUninstall,omitempty"`

	// Indicates the Windows architecture(s) this app should be installed on. The app will be treated as not applicable for
	// devices with architectures not matching the selected value. When a non-null value is provided for the
	// allowedArchitectures property, the value of the applicableArchitectures property is set to none. Possible values are:
	// null, x86, x64, arm64. Possible values are: none, x86, x64, arm, neutral, arm64.
	AllowedArchitectures *WindowsArchitecture `json:"allowedArchitectures,omitempty"`

	// Contains properties for Windows architecture.
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`

	// Indicates the detection rules to detect Win32 Line of Business (LoB) app. Possible values are
	// Win32LobAppPowerShellScriptDetection, Win32LobAppRegistryDetection, Win32LobAppFileSystemDetection,
	// Win32LobAppProductCodeDetection.
	DetectionRules *[]Win32LobAppDetection `json:"detectionRules,omitempty"`

	// Indicates the version displayed in the UX for this app. Used to set the version of the app. Example: 1.0.3.215.
	DisplayVersion nullable.Type[string] `json:"displayVersion,omitempty"`

	// Indicates the command line to install this app. Used to install the Win32 app. Example: msiexec /i 'Orca.Msi' /qn.
	InstallCommandLine nullable.Type[string] `json:"installCommandLine,omitempty"`

	// Indicates the install experience for this app.
	InstallExperience *Win32LobAppInstallExperience `json:"installExperience,omitempty"`

	// Indicates the value for the minimum CPU speed which is required to install this app. Allowed range from 0 to clock
	// speed from WMI helper.
	MinimumCpuSpeedInMHz nullable.Type[int64] `json:"minimumCpuSpeedInMHz,omitempty"`

	// Indicates the value for the minimum free disk space which is required to install this app. Allowed range from 0 to
	// driver's maximum available free space.
	MinimumFreeDiskSpaceInMB nullable.Type[int64] `json:"minimumFreeDiskSpaceInMB,omitempty"`

	// Indicates the value for the minimum physical memory which is required to install this app. Allowed range from 0 to
	// total physical memory from WMI helper.
	MinimumMemoryInMB nullable.Type[int64] `json:"minimumMemoryInMB,omitempty"`

	// Indicates the value for the minimum number of processors which is required to install this app. Minimum value is 0.
	MinimumNumberOfProcessors nullable.Type[int64] `json:"minimumNumberOfProcessors,omitempty"`

	// Indicates the value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`

	// Indicates the value for the minimum supported windows release. Example: Windows11_23H2.
	MinimumSupportedWindowsRelease nullable.Type[string] `json:"minimumSupportedWindowsRelease,omitempty"`

	// Indicates the MSI details if this Win32 app is an MSI app.
	MsiInformation *Win32LobAppMsiInformation `json:"msiInformation,omitempty"`

	// Indicates the requirement rules to detect Win32 Line of Business (LoB) app. Possible values are:
	// Win32LobAppFileSystemRequirement, Win32LobAppPowerShellScriptRequirement, Win32LobAppRegistryRequirement.
	RequirementRules *[]Win32LobAppRequirement `json:"requirementRules,omitempty"`

	// Indicates the return codes for post installation behavior.
	ReturnCodes *[]Win32LobAppReturnCode `json:"returnCodes,omitempty"`

	// Indicates the detection and requirement rules for this app. Possible values are: Win32LobAppFileSystemRule,
	// Win32LobAppPowerShellScriptRule, Win32LobAppProductCodeRule, Win32LobAppRegistryRule.
	Rules *[]Win32LobAppRule `json:"rules,omitempty"`

	// Indicates the relative path of the setup file in the encrypted Win32LobApp package. Example: Intel-SA-00075 Detection
	// and Mitigation Tool.msi.
	SetupFilePath nullable.Type[string] `json:"setupFilePath,omitempty"`

	// Indicates the command line to uninstall this app. Used to uninstall the app. Example: msiexec /x
	// '{85F4CBCB-9BBC-4B50-A7D8-E1106771498D}' /qn.
	UninstallCommandLine nullable.Type[string] `json:"uninstallCommandLine,omitempty"`

	// Fields inherited from MobileLobApp

	// The internal committed content version.
	CommittedContentVersion nullable.Type[string] `json:"committedContentVersion,omitempty"`

	// The list of content versions for this app. This property is read-only.
	ContentVersions *[]MobileAppContent `json:"contentVersions,omitempty"`

	// The name of the main Lob application file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The total size, including all uploaded files. This property is read-only.
	Size *int64 `json:"size,omitempty"`

	// Fields inherited from MobileApp

	// The list of group assignments for this mobile app.
	Assignments *[]MobileAppAssignment `json:"assignments,omitempty"`

	// The list of categories for this app.
	Categories *[]MobileAppCategory `json:"categories,omitempty"`

	// The date and time the app was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The total number of dependencies the child app has.
	DependentAppCount *int64 `json:"dependentAppCount,omitempty"`

	// The description of the app.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The developer of the app.
	Developer nullable.Type[string] `json:"developer,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// The value indicating whether the app is assigned to at least one group.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Notes for the app.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The owner of the app.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The publisher of the app.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates the publishing state of an app.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`

	// List of relationships for this mobile app.
	Relationships *[]MobileAppRelationship `json:"relationships,omitempty"`

	// List of scope tag ids for this mobile app.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The total number of apps this app is directly or indirectly superseded by. This property is read-only.
	SupersededAppCount *int64 `json:"supersededAppCount,omitempty"`

	// The total number of apps this app directly or indirectly supersedes. This property is read-only.
	SupersedingAppCount *int64 `json:"supersedingAppCount,omitempty"`

	// The upload state.
	UploadState *int64 `json:"uploadState,omitempty"`

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

func (s BaseWin32LobAppImpl) Win32LobApp() BaseWin32LobAppImpl {
	return s
}

func (s BaseWin32LobAppImpl) MobileLobApp() BaseMobileLobAppImpl {
	return BaseMobileLobAppImpl{
		CommittedContentVersion: s.CommittedContentVersion,
		ContentVersions:         s.ContentVersions,
		FileName:                s.FileName,
		Size:                    s.Size,
		Assignments:             s.Assignments,
		Categories:              s.Categories,
		CreatedDateTime:         s.CreatedDateTime,
		DependentAppCount:       s.DependentAppCount,
		Description:             s.Description,
		Developer:               s.Developer,
		DisplayName:             s.DisplayName,
		InformationUrl:          s.InformationUrl,
		IsAssigned:              s.IsAssigned,
		IsFeatured:              s.IsFeatured,
		LargeIcon:               s.LargeIcon,
		LastModifiedDateTime:    s.LastModifiedDateTime,
		Notes:                   s.Notes,
		Owner:                   s.Owner,
		PrivacyInformationUrl:   s.PrivacyInformationUrl,
		Publisher:               s.Publisher,
		PublishingState:         s.PublishingState,
		Relationships:           s.Relationships,
		RoleScopeTagIds:         s.RoleScopeTagIds,
		SupersededAppCount:      s.SupersededAppCount,
		SupersedingAppCount:     s.SupersedingAppCount,
		UploadState:             s.UploadState,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s BaseWin32LobAppImpl) MobileApp() BaseMobileAppImpl {
	return BaseMobileAppImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		DependentAppCount:     s.DependentAppCount,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsAssigned:            s.IsAssigned,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Relationships:         s.Relationships,
		RoleScopeTagIds:       s.RoleScopeTagIds,
		SupersededAppCount:    s.SupersededAppCount,
		SupersedingAppCount:   s.SupersedingAppCount,
		UploadState:           s.UploadState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s BaseWin32LobAppImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Win32LobApp = RawWin32LobAppImpl{}

// RawWin32LobAppImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWin32LobAppImpl struct {
	win32LobApp BaseWin32LobAppImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawWin32LobAppImpl) Win32LobApp() BaseWin32LobAppImpl {
	return s.win32LobApp
}

func (s RawWin32LobAppImpl) MobileLobApp() BaseMobileLobAppImpl {
	return s.win32LobApp.MobileLobApp()
}

func (s RawWin32LobAppImpl) MobileApp() BaseMobileAppImpl {
	return s.win32LobApp.MobileApp()
}

func (s RawWin32LobAppImpl) Entity() BaseEntityImpl {
	return s.win32LobApp.Entity()
}

var _ json.Marshaler = BaseWin32LobAppImpl{}

func (s BaseWin32LobAppImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWin32LobAppImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWin32LobAppImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWin32LobAppImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWin32LobAppImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseWin32LobAppImpl{}

func (s *BaseWin32LobAppImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowAvailableUninstall         *bool                          `json:"allowAvailableUninstall,omitempty"`
		AllowedArchitectures            *WindowsArchitecture           `json:"allowedArchitectures,omitempty"`
		ApplicableArchitectures         *WindowsArchitecture           `json:"applicableArchitectures,omitempty"`
		DisplayVersion                  nullable.Type[string]          `json:"displayVersion,omitempty"`
		InstallCommandLine              nullable.Type[string]          `json:"installCommandLine,omitempty"`
		InstallExperience               *Win32LobAppInstallExperience  `json:"installExperience,omitempty"`
		MinimumCpuSpeedInMHz            nullable.Type[int64]           `json:"minimumCpuSpeedInMHz,omitempty"`
		MinimumFreeDiskSpaceInMB        nullable.Type[int64]           `json:"minimumFreeDiskSpaceInMB,omitempty"`
		MinimumMemoryInMB               nullable.Type[int64]           `json:"minimumMemoryInMB,omitempty"`
		MinimumNumberOfProcessors       nullable.Type[int64]           `json:"minimumNumberOfProcessors,omitempty"`
		MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
		MinimumSupportedWindowsRelease  nullable.Type[string]          `json:"minimumSupportedWindowsRelease,omitempty"`
		MsiInformation                  *Win32LobAppMsiInformation     `json:"msiInformation,omitempty"`
		ReturnCodes                     *[]Win32LobAppReturnCode       `json:"returnCodes,omitempty"`
		SetupFilePath                   nullable.Type[string]          `json:"setupFilePath,omitempty"`
		UninstallCommandLine            nullable.Type[string]          `json:"uninstallCommandLine,omitempty"`
		CommittedContentVersion         nullable.Type[string]          `json:"committedContentVersion,omitempty"`
		ContentVersions                 *[]MobileAppContent            `json:"contentVersions,omitempty"`
		FileName                        nullable.Type[string]          `json:"fileName,omitempty"`
		Size                            *int64                         `json:"size,omitempty"`
		Assignments                     *[]MobileAppAssignment         `json:"assignments,omitempty"`
		Categories                      *[]MobileAppCategory           `json:"categories,omitempty"`
		CreatedDateTime                 *string                        `json:"createdDateTime,omitempty"`
		DependentAppCount               *int64                         `json:"dependentAppCount,omitempty"`
		Description                     nullable.Type[string]          `json:"description,omitempty"`
		Developer                       nullable.Type[string]          `json:"developer,omitempty"`
		DisplayName                     nullable.Type[string]          `json:"displayName,omitempty"`
		InformationUrl                  nullable.Type[string]          `json:"informationUrl,omitempty"`
		IsAssigned                      *bool                          `json:"isAssigned,omitempty"`
		IsFeatured                      *bool                          `json:"isFeatured,omitempty"`
		LargeIcon                       *MimeContent                   `json:"largeIcon,omitempty"`
		LastModifiedDateTime            *string                        `json:"lastModifiedDateTime,omitempty"`
		Notes                           nullable.Type[string]          `json:"notes,omitempty"`
		Owner                           nullable.Type[string]          `json:"owner,omitempty"`
		PrivacyInformationUrl           nullable.Type[string]          `json:"privacyInformationUrl,omitempty"`
		Publisher                       nullable.Type[string]          `json:"publisher,omitempty"`
		PublishingState                 *MobileAppPublishingState      `json:"publishingState,omitempty"`
		RoleScopeTagIds                 *[]string                      `json:"roleScopeTagIds,omitempty"`
		SupersededAppCount              *int64                         `json:"supersededAppCount,omitempty"`
		SupersedingAppCount             *int64                         `json:"supersedingAppCount,omitempty"`
		UploadState                     *int64                         `json:"uploadState,omitempty"`
		Id                              *string                        `json:"id,omitempty"`
		ODataId                         *string                        `json:"@odata.id,omitempty"`
		ODataType                       *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowAvailableUninstall = decoded.AllowAvailableUninstall
	s.AllowedArchitectures = decoded.AllowedArchitectures
	s.ApplicableArchitectures = decoded.ApplicableArchitectures
	s.DisplayVersion = decoded.DisplayVersion
	s.InstallCommandLine = decoded.InstallCommandLine
	s.InstallExperience = decoded.InstallExperience
	s.MinimumCpuSpeedInMHz = decoded.MinimumCpuSpeedInMHz
	s.MinimumFreeDiskSpaceInMB = decoded.MinimumFreeDiskSpaceInMB
	s.MinimumMemoryInMB = decoded.MinimumMemoryInMB
	s.MinimumNumberOfProcessors = decoded.MinimumNumberOfProcessors
	s.MinimumSupportedOperatingSystem = decoded.MinimumSupportedOperatingSystem
	s.MinimumSupportedWindowsRelease = decoded.MinimumSupportedWindowsRelease
	s.MsiInformation = decoded.MsiInformation
	s.ReturnCodes = decoded.ReturnCodes
	s.SetupFilePath = decoded.SetupFilePath
	s.UninstallCommandLine = decoded.UninstallCommandLine
	s.Assignments = decoded.Assignments
	s.Categories = decoded.Categories
	s.CommittedContentVersion = decoded.CommittedContentVersion
	s.ContentVersions = decoded.ContentVersions
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DependentAppCount = decoded.DependentAppCount
	s.Description = decoded.Description
	s.Developer = decoded.Developer
	s.DisplayName = decoded.DisplayName
	s.FileName = decoded.FileName
	s.Id = decoded.Id
	s.InformationUrl = decoded.InformationUrl
	s.IsAssigned = decoded.IsAssigned
	s.IsFeatured = decoded.IsFeatured
	s.LargeIcon = decoded.LargeIcon
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Notes = decoded.Notes
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Owner = decoded.Owner
	s.PrivacyInformationUrl = decoded.PrivacyInformationUrl
	s.Publisher = decoded.Publisher
	s.PublishingState = decoded.PublishingState
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Size = decoded.Size
	s.SupersededAppCount = decoded.SupersededAppCount
	s.SupersedingAppCount = decoded.SupersedingAppCount
	s.UploadState = decoded.UploadState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseWin32LobAppImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["detectionRules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DetectionRules into list []json.RawMessage: %+v", err)
		}

		output := make([]Win32LobAppDetection, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWin32LobAppDetectionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DetectionRules' for 'BaseWin32LobAppImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DetectionRules = &output
	}

	if v, ok := temp["relationships"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Relationships into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileAppRelationship, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileAppRelationshipImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Relationships' for 'BaseWin32LobAppImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Relationships = &output
	}

	if v, ok := temp["requirementRules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RequirementRules into list []json.RawMessage: %+v", err)
		}

		output := make([]Win32LobAppRequirement, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWin32LobAppRequirementImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RequirementRules' for 'BaseWin32LobAppImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RequirementRules = &output
	}

	if v, ok := temp["rules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Rules into list []json.RawMessage: %+v", err)
		}

		output := make([]Win32LobAppRule, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWin32LobAppRuleImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Rules' for 'BaseWin32LobAppImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Rules = &output
	}

	return nil
}

func UnmarshalWin32LobAppImplementation(input []byte) (Win32LobApp, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobApp into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.win32CatalogApp") {
		var out Win32CatalogApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32CatalogApp: %+v", err)
		}
		return out, nil
	}

	var parent BaseWin32LobAppImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWin32LobAppImpl: %+v", err)
	}

	return RawWin32LobAppImpl{
		win32LobApp: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
