package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileLobApp = WindowsUniversalAppX{}

type WindowsUniversalAppX struct {
	// Contains properties for Windows architecture.
	ApplicableArchitectures *WindowsArchitecture `json:"applicableArchitectures,omitempty"`

	// Contains properties for Windows device type. Multiple values can be selected. Default value is `none`.
	ApplicableDeviceTypes *WindowsDeviceType `json:"applicableDeviceTypes,omitempty"`

	// The collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app. This property is
	// read-only.
	CommittedContainedApps *[]MobileContainedApp `json:"committedContainedApps,omitempty"`

	// The Identity Name of the app, parsed from the appx file when it is uploaded through the Intune MEM console. For
	// example: 'Contoso.DemoApp'.
	IdentityName nullable.Type[string] `json:"identityName,omitempty"`

	// The Identity Publisher Hash of the app, parsed from the appx file when it is uploaded through the Intune MEM console.
	// For example: 'AB82CD0XYZ'.
	IdentityPublisherHash *string `json:"identityPublisherHash,omitempty"`

	// The Identity Resource Identifier of the app, parsed from the appx file when it is uploaded through the Intune MEM
	// console. For example: 'TestResourceId'.
	IdentityResourceIdentifier nullable.Type[string] `json:"identityResourceIdentifier,omitempty"`

	// The Identity Version of the app, parsed from the appx file when it is uploaded through the Intune MEM console. For
	// example: '1.0.0.0'.
	IdentityVersion nullable.Type[string] `json:"identityVersion,omitempty"`

	// Whether or not the app is a bundle. If TRUE, app is a bundle; if FALSE, app is not a bundle.
	IsBundle *bool `json:"isBundle,omitempty"`

	// The minimum operating system required for a Windows mobile app.
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`

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

func (s WindowsUniversalAppX) MobileLobApp() BaseMobileLobAppImpl {
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

func (s WindowsUniversalAppX) MobileApp() BaseMobileAppImpl {
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

func (s WindowsUniversalAppX) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUniversalAppX{}

func (s WindowsUniversalAppX) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUniversalAppX
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUniversalAppX: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUniversalAppX: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUniversalAppX"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUniversalAppX: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUniversalAppX{}

func (s *WindowsUniversalAppX) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicableArchitectures         *WindowsArchitecture           `json:"applicableArchitectures,omitempty"`
		ApplicableDeviceTypes           *WindowsDeviceType             `json:"applicableDeviceTypes,omitempty"`
		IdentityName                    nullable.Type[string]          `json:"identityName,omitempty"`
		IdentityPublisherHash           *string                        `json:"identityPublisherHash,omitempty"`
		IdentityResourceIdentifier      nullable.Type[string]          `json:"identityResourceIdentifier,omitempty"`
		IdentityVersion                 nullable.Type[string]          `json:"identityVersion,omitempty"`
		IsBundle                        *bool                          `json:"isBundle,omitempty"`
		MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
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

	s.ApplicableArchitectures = decoded.ApplicableArchitectures
	s.ApplicableDeviceTypes = decoded.ApplicableDeviceTypes
	s.IdentityName = decoded.IdentityName
	s.IdentityPublisherHash = decoded.IdentityPublisherHash
	s.IdentityResourceIdentifier = decoded.IdentityResourceIdentifier
	s.IdentityVersion = decoded.IdentityVersion
	s.IsBundle = decoded.IsBundle
	s.MinimumSupportedOperatingSystem = decoded.MinimumSupportedOperatingSystem
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
		return fmt.Errorf("unmarshaling WindowsUniversalAppX into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["committedContainedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CommittedContainedApps into list []json.RawMessage: %+v", err)
		}

		output := make([]MobileContainedApp, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMobileContainedAppImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CommittedContainedApps' for 'WindowsUniversalAppX': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CommittedContainedApps = &output
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
				return fmt.Errorf("unmarshaling index %d field 'Relationships' for 'WindowsUniversalAppX': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Relationships = &output
	}

	return nil
}
