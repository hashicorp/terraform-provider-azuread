package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileApp = OfficeSuiteApp{}

type OfficeSuiteApp struct {
	// The value to accept the EULA automatically on the enduser's device.
	AutoAcceptEula *bool `json:"autoAcceptEula,omitempty"`

	// The property to represent the apps which are excluded from the selected Office365 Product Id.
	ExcludedApps *ExcludedApps `json:"excludedApps,omitempty"`

	// The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
	InstallProgressDisplayLevel *OfficeSuiteInstallProgressDisplayLevel `json:"installProgressDisplayLevel,omitempty"`

	// The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard
	// RFC 6033. Ref: https://technet.microsoft.com/library/cc179219(v=office.16).aspx
	LocalesToInstall *[]string `json:"localesToInstall,omitempty"`

	// The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence
	// over all other properties. When present, the XML configuration file will be used to create the app.
	OfficeConfigurationXml nullable.Type[string] `json:"officeConfigurationXml,omitempty"`

	// Contains properties for Windows architecture.
	OfficePlatformArchitecture *WindowsArchitecture `json:"officePlatformArchitecture,omitempty"`

	// Describes the OfficeSuiteApp file format types that can be selected.
	OfficeSuiteAppDefaultFileFormat *OfficeSuiteDefaultFileFormatType `json:"officeSuiteAppDefaultFileFormat,omitempty"`

	// The Product Ids that represent the Office365 Suite SKU.
	ProductIds *[]OfficeProductId `json:"productIds,omitempty"`

	// The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the
	// device or not.
	ShouldUninstallOlderVersionsOfOffice *bool `json:"shouldUninstallOlderVersionsOfOffice,omitempty"`

	// The property to represent the specific target version for the Office365 app suite that should be remained deployed on
	// the devices.
	TargetVersion nullable.Type[string] `json:"targetVersion,omitempty"`

	// The Enum to specify the Office365 Updates Channel.
	UpdateChannel *OfficeUpdateChannel `json:"updateChannel,omitempty"`

	// The property to represent the update version in which the specific target version is available for the Office365 app
	// suite.
	UpdateVersion nullable.Type[string] `json:"updateVersion,omitempty"`

	// The property to represent that whether the shared computer activation is used not for Office365 app suite.
	UseSharedComputerActivation *bool `json:"useSharedComputerActivation,omitempty"`

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

func (s OfficeSuiteApp) MobileApp() BaseMobileAppImpl {
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

func (s OfficeSuiteApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OfficeSuiteApp{}

func (s OfficeSuiteApp) MarshalJSON() ([]byte, error) {
	type wrapper OfficeSuiteApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OfficeSuiteApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OfficeSuiteApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.officeSuiteApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OfficeSuiteApp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OfficeSuiteApp{}

func (s *OfficeSuiteApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AutoAcceptEula                       *bool                                   `json:"autoAcceptEula,omitempty"`
		ExcludedApps                         *ExcludedApps                           `json:"excludedApps,omitempty"`
		InstallProgressDisplayLevel          *OfficeSuiteInstallProgressDisplayLevel `json:"installProgressDisplayLevel,omitempty"`
		LocalesToInstall                     *[]string                               `json:"localesToInstall,omitempty"`
		OfficeConfigurationXml               nullable.Type[string]                   `json:"officeConfigurationXml,omitempty"`
		OfficePlatformArchitecture           *WindowsArchitecture                    `json:"officePlatformArchitecture,omitempty"`
		OfficeSuiteAppDefaultFileFormat      *OfficeSuiteDefaultFileFormatType       `json:"officeSuiteAppDefaultFileFormat,omitempty"`
		ProductIds                           *[]OfficeProductId                      `json:"productIds,omitempty"`
		ShouldUninstallOlderVersionsOfOffice *bool                                   `json:"shouldUninstallOlderVersionsOfOffice,omitempty"`
		TargetVersion                        nullable.Type[string]                   `json:"targetVersion,omitempty"`
		UpdateChannel                        *OfficeUpdateChannel                    `json:"updateChannel,omitempty"`
		UpdateVersion                        nullable.Type[string]                   `json:"updateVersion,omitempty"`
		UseSharedComputerActivation          *bool                                   `json:"useSharedComputerActivation,omitempty"`
		Assignments                          *[]MobileAppAssignment                  `json:"assignments,omitempty"`
		Categories                           *[]MobileAppCategory                    `json:"categories,omitempty"`
		CreatedDateTime                      *string                                 `json:"createdDateTime,omitempty"`
		DependentAppCount                    *int64                                  `json:"dependentAppCount,omitempty"`
		Description                          nullable.Type[string]                   `json:"description,omitempty"`
		Developer                            nullable.Type[string]                   `json:"developer,omitempty"`
		DisplayName                          nullable.Type[string]                   `json:"displayName,omitempty"`
		InformationUrl                       nullable.Type[string]                   `json:"informationUrl,omitempty"`
		IsAssigned                           *bool                                   `json:"isAssigned,omitempty"`
		IsFeatured                           *bool                                   `json:"isFeatured,omitempty"`
		LargeIcon                            *MimeContent                            `json:"largeIcon,omitempty"`
		LastModifiedDateTime                 *string                                 `json:"lastModifiedDateTime,omitempty"`
		Notes                                nullable.Type[string]                   `json:"notes,omitempty"`
		Owner                                nullable.Type[string]                   `json:"owner,omitempty"`
		PrivacyInformationUrl                nullable.Type[string]                   `json:"privacyInformationUrl,omitempty"`
		Publisher                            nullable.Type[string]                   `json:"publisher,omitempty"`
		PublishingState                      *MobileAppPublishingState               `json:"publishingState,omitempty"`
		RoleScopeTagIds                      *[]string                               `json:"roleScopeTagIds,omitempty"`
		SupersededAppCount                   *int64                                  `json:"supersededAppCount,omitempty"`
		SupersedingAppCount                  *int64                                  `json:"supersedingAppCount,omitempty"`
		UploadState                          *int64                                  `json:"uploadState,omitempty"`
		Id                                   *string                                 `json:"id,omitempty"`
		ODataId                              *string                                 `json:"@odata.id,omitempty"`
		ODataType                            *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AutoAcceptEula = decoded.AutoAcceptEula
	s.ExcludedApps = decoded.ExcludedApps
	s.InstallProgressDisplayLevel = decoded.InstallProgressDisplayLevel
	s.LocalesToInstall = decoded.LocalesToInstall
	s.OfficeConfigurationXml = decoded.OfficeConfigurationXml
	s.OfficePlatformArchitecture = decoded.OfficePlatformArchitecture
	s.OfficeSuiteAppDefaultFileFormat = decoded.OfficeSuiteAppDefaultFileFormat
	s.ProductIds = decoded.ProductIds
	s.ShouldUninstallOlderVersionsOfOffice = decoded.ShouldUninstallOlderVersionsOfOffice
	s.TargetVersion = decoded.TargetVersion
	s.UpdateChannel = decoded.UpdateChannel
	s.UpdateVersion = decoded.UpdateVersion
	s.UseSharedComputerActivation = decoded.UseSharedComputerActivation
	s.Assignments = decoded.Assignments
	s.Categories = decoded.Categories
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DependentAppCount = decoded.DependentAppCount
	s.Description = decoded.Description
	s.Developer = decoded.Developer
	s.DisplayName = decoded.DisplayName
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
	s.SupersededAppCount = decoded.SupersededAppCount
	s.SupersedingAppCount = decoded.SupersedingAppCount
	s.UploadState = decoded.UploadState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OfficeSuiteApp into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Relationships' for 'OfficeSuiteApp': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Relationships = &output
	}

	return nil
}
